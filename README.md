# simple-notification-system

# about 
    The SNS implementation sends a notification when a user is created, publishing a message to the user-notifications topic in LocalStack.
    
# sns topics
    docker exec localstack awslocal sns list-topics --endpoint-url http://localhost:4566

# list messages in sqs queue
    docker exec localstack awslocal sqs receive-message --queue-url http://localhost:4566/000000000000/user-queue --endpoint-url http://localhost:4566

  json
    {
        "Messages": [
            {
            "MessageId": "1234-5678-...",
            "ReceiptHandle": "...",
            "Body": "{\"Message\":\"{\\\"id\\\":1,\\\"email\\\":\\\"alice@example.com\\\",\\\"name\\\":\\\"Alice\\\",\\\"created_at\\\":\\\"2025-06-08T14:00:00Z\\\"}\",\"Subject\":null,...}"
            }
        ]
        }

# The Body contains the SNS message, which is a JSON object wrapping the user data. Use jq to parse:
    docker exec localstack awslocal sqs receive-message --queue-url http://localhost:4566/000000000000/user-queue --endpoint-url http://localhost:4566 | jq '.Messages[].Body | fromjson | .Message | fromjson'

    {"id":1,"email":"alice@example.com","name":"Alice","created_at":"2025-06-08T14:00:00Z"}


# list sns topics to confirm
    docker exec localstack awslocal sns list-topics --endpoint-url http://localhost:4566

# List subscriptions to verify SQS is subscribed
    docker exec localstack awslocal sns list-subscriptions-by-topic --topic-arn arn:aws:sns:us-east-1:000000000000:user-notifications --endpoint-url http://localhost:4566

