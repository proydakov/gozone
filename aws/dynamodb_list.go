package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
    svc := dynamodb.New(session.New(), &aws.Config{Region: aws.String("eu-central-1")})

    params := &dynamodb.ListTablesInput{
        ExclusiveStartTableName: aws.String("TableName"),
        Limit: aws.Int64(1),
    }
    resp, err := svc.ListTables(params)

    if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
        return
    }

    // Pretty-print the response data.
    fmt.Println(resp)
}
