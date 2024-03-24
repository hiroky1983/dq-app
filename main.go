package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	prams := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		TableName: aws.String("test_table"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType: aws.String("HASH"),
			},
		},
		BillingMode: aws.String("PAY_PER_REQUEST"),
	}
	resp, err := ddb.CreateTable(prams)
	if err != nil {
		panic(err)
	}
	println(resp)
}