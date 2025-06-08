package sns_notification

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"user-notification/user/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gin-gonic/gin"
)

type SNSService struct {
	SnsClient *sns.Client
	SqsClient *sqs.Client
	TopicARN  string
	QueueURL  string
}

func NewSNSService(topicARN, queueURL string) *SNSService {
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localstack:4566"}, nil
			},
		)),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     "test",
				SecretAccessKey: "test",
			}, nil
		})),
	)
	if err != nil {
		// log.Fatalf("Failed to load AWS config: %v", err)
		fmt.Println("Failed to load AWS config: err - ", err)
	}
	snsClient := sns.NewFromConfig(awsCfg)
	sqsClient := sqs.NewFromConfig(awsCfg)
	return &SNSService{
		SnsClient: snsClient,
		SqsClient: sqsClient,
		TopicARN:  topicARN,
		QueueURL:  queueURL,
	}
}

func (s *SNSService) PublishUserCreated(ctx context.Context, user *model.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %v", err)
	}
	_, err = s.SnsClient.Publish(ctx, &sns.PublishInput{
		TopicArn: aws.String(s.TopicARN),
		Message:  aws.String(string(userJSON)),
	})
	if err != nil {
		return fmt.Errorf("failed to publish SNS message: %v", err)
	}
	return nil
}

func (service *SNSService) GetMessages(ctx context.Context) ([]string, error) {
	output, err := service.SqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(service.QueueURL),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     5,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to receive SQS messages: %v", err)
	}
	var messages []string
	for _, msg := range output.Messages {
		var snsMsg struct {
			Message string `json:"Message"`
		}
		if err := json.Unmarshal([]byte(*msg.Body), &snsMsg); err != nil {
			continue
		}
		messages = append(messages, snsMsg.Message)
		_, err = service.SqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(service.QueueURL),
			ReceiptHandle: msg.ReceiptHandle,
		})
		if err != nil {
			continue
		}
	}
	return messages, nil
}

func (s *SNSService) RegisterSNSRoutes(r *gin.Engine) {
	r.GET("/notifications", s.GetNotifications)
}

func (s *SNSService) GetNotifications(ctx *gin.Context) {
	messages, err := s.GetMessages(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}
