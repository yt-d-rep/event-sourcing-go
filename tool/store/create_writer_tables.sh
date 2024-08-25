#!/usr/bin/env bash

# Create the tables for the writer service

export AWS_REGION=ap-northeast-1
export AWS_ACCESS_KEY_ID=dummy
export AWS_SECRET_ACCESS_KEY=dummy
readonly ENDPOINT=http://dynamodb:8000

aws dynamodb create-table --table-name user \
  --attribute-definitions \
    AttributeName=ID,AttributeType=N \
    AttributeName=CreatedAt,AttributeType=S \
  --key-schema \
    AttributeName=ID,KeyType=HASH \
    AttributeName=CreatedAt,KeyType=RANGE \
  --provisioned-throughput \
    ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --endpoint-url ${ENDPOINT}