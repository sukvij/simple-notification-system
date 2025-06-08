#!/bin/bash
awslocal sns create-topic --name user-notifications
awslocal sqs create-queue --queue-name user-queue
awslocal sns subscribe \
  --topic-arn arn:aws:sns:us-east-1:000000000000:user-notifications \
  --protocol sqs \
  --notification-endpoint arn:aws:sqs:us-east-1:000000000000:user-queue

# multiple topic and multiple queues has subscribed it.
  #!/bin/bash

# Create SNS Topics
# awslocal sns create-topic --name order-events
# awslocal sns create-topic --name payment-events
# awslocal sns create-topic --name offer-notifications

# # Create SQS Queues (services)
# awslocal sqs create-queue --queue-name inventory-queue
# awslocal sqs create-queue --queue-name analytics-queue
# awslocal sqs create-queue --queue-name billing-queue
# awslocal sqs create-queue --queue-name fraud-queue
# awslocal sqs create-queue --queue-name notification-queue
# awslocal sqs create-queue --queue-name marketing-queue

# # Subscriptions
# # Subscribe inventory and analytics to order-events
# awslocal sns subscribe \
#   --topic-arn arn:aws:sns:us-east-1:000000000000:order-events \
#   --protocol sqs \
#   --notification-endpoint arn:aws:sqs:us-east-1:000000000000:inventory-queue

# awslocal sns subscribe \
#   --topic-arn arn:aws:sns:us-east-1:000000000000:order-events \
#   --protocol sqs \
#   --notification-endpoint arn:aws:sqs:us-east-1:000000000000:analytics-queue

# # Subscribe billing and fraud to payment-events
# awslocal sns subscribe \
#   --topic-arn arn:aws:sns:us-east-1:000000000000:payment-events \
#   --protocol sqs \
#   --notification-endpoint arn:aws:sqs:us-east-1:000000000000:billing-queue

# awslocal sns subscribe \
#   --topic-arn arn:aws:sns:us-east-1:000000000000:payment-events \
#   --protocol sqs \
#   --notification-endpoint arn:aws:sqs:us-east-1:000000000000:fraud-queue

# # Subscribe notification and marketing to offer-notifications
# awslocal sns subscribe \
#   --topic-arn arn:aws:sns:us-east-1:000000000000:offer-notifications \
#   --protocol sqs \
#   --notification-endpoint arn:aws:sqs:us-east-1:000000000000:notification-queue

# awslocal sns subscribe \
#   --topic-arn arn:aws:sns:us-east-1:000000000000:offer-notifications \
#   --protocol sqs \
#   --notification-endpoint arn:aws:sqs:us-east-1:000000000000:marketing-queue
