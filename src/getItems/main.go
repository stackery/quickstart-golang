package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Create struct to hold info about new item
type Item struct {
	Id      string
	Content string
}

func Handler(ctx context.Context, event interface{}) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := os.Getenv("TABLE_NAME") // get the table name from the automatically populated environment variables

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Use dynamodb to get items from the ItemTable
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Printf("Error getting items from table %s: %s\n", tableName, err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result: %#v\n", result)

	for _, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Item " + item.Id + ": " + item.Content)
		fmt.Println()
	}

	return fmt.Sprintf("%d items found", *result.Count), nil
}

func main() {
	lambda.Start(Handler)
}
