package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
    svc := dynamodb.New(session.New(), &aws.Config{Region: aws.String("eu-central-1")})

    params := &dynamodb.CreateTableInput{
        TableName: aws.String("ProductCatalog"),
        AttributeDefinitions: []*dynamodb.AttributeDefinition{
            {
                AttributeName: aws.String("Id"),
                AttributeType: aws.String("N"),
            },
        },
        KeySchema: []*dynamodb.KeySchemaElement{
            {
                AttributeName: aws.String("Id"),
                KeyType: aws.String("HASH"),
            },
        },
        ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
            ReadCapacityUnits: aws.Int64(5),
            WriteCapacityUnits: aws.Int64(5),
        },
    }
    resp, err := svc.CreateTable(params)

    if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
        return
    }

    // Pretty-print the response data.
    fmt.Println(resp)
}
