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

type Key struct {
    Email string
}

func main() {
    svc := dynamodb.New(session.New(), &aws.Config{Region: aws.String("eu-central-1")})

    k := Key {
        Email: "test@gmail.com",
    }

    item, err := dynamodbattribute.ConvertToMap(k)
    if err != nil {
        panic(err)
    }

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String("Users"),
        Key:       item,
    })
    if err != nil {
        panic(err)
    }

    r := Record{}
    err = dynamodbattribute.ConvertFromMap(result.Item, &r)

    fmt.Println(r)
}
