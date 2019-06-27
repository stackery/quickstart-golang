package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"os"
	"strconv"
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
  
  item := Item{
    Id:   "1", // modify with each invoke so the id does not repeat
    Content:  "This is my content" // modify content here
  }

  input := &dynamodb.PutItemInput{
    Item:      item,
    TableName: aws.String(os.Getenv("TABLE_NAME")) // get the table name from the automatically populated environment variables
  }

  _, err = svc.PutItem(input)
  if err != nil {
      fmt.Println("Error adding " + itemId + " to table " + tableName )
      fmt.Println(err.Error())
      os.Exit(1)
  }

  fmt.Println("Adding item '" + itemId + " to table " + tableName)
}

func main() {
	lambda.Start(Handler)
}
