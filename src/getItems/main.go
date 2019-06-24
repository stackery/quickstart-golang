package main

import (
    "fmt"
    "context"
    "github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event interface{}) (string, error) {
  fmt.Println(event)

  return "OK", nil
}

func main() {
  lambda.Start(Handler)
}
