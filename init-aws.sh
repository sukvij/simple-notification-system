#!/bin/bash
awslocal sns create-topic --name user-notifications
awslocal sqs create-queue --queue-name user-queue
awslocal sns subscribe \
  --topic-arn arn:aws:sns:us-east-1:000000000000:user-notifications \
  --protocol sqs \
  --notification-endpoint arn:aws:sqs:us-east-1:000000000000:user-queue