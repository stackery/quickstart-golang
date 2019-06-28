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
	Id      string `json:"id"`
	Content string `json:"content"`
}

func Handler(ctx context.Context, event interface{}) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	item := Item{
		Id:      "1",                  // modify with each invoke so the id does not repeat
		Content: "This is my content"	 // modify content here
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshaling item: ", err.Error())
		os.Exit(1)
	}

	tableName := os.Getenv("TABLE_NAME") // get the table name from the automatically populated environment variables

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Error adding " + item.Id + " to table " + tableName)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Adding item '" + item.Id + " to table " + tableName)
	return "", nil
}

func main() {
	lambda.Start(Handler)
}
