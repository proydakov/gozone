package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Record struct {
    Email           string
    CheckEmailToken string
}

func main() {
    svc := dynamodb.New(session.New(), &aws.Config{Region: aws.String("eu-central-1")})

    r := Record{
        Email:           "test@gmail.com",
        CheckEmailToken: "29jd9j2dm3932jdhewk4",
    }

    item, err := dynamodbattribute.ConvertToMap(r)
    if err != nil {
        panic(err)
    }
    result, err := svc.PutItem(&dynamodb.PutItemInput{
        TableName: aws.String("Users"),
        Item:      item,
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
