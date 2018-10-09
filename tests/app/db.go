package main

import (
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var config = &aws.Config{
    Region:   aws.String("eu-west-1"),
    MaxRetries: aws.Int(0),
}

var db  = dynamodb.New(
    session.Must(session.NewSession(config)))

//var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-west-1"))

func getTableName(stage string) string {
    return os.Getenv("DynamoDbTableName"+ "_" + stage)
}

func getItem(isbn string, stage string) (*book, error) {
    input := &dynamodb.GetItemInput{
        TableName: aws.String(getTableName(stage)),
        Key: map[string]*dynamodb.AttributeValue{
            "ISBN": {
                S: aws.String(isbn),
            },
        },
    }

    result, err := db.GetItem(input)
    if err != nil {
        return nil, err
    }
    if result.Item == nil {
        return nil, nil
    }

    bk := new(book)
    err = dynamodbattribute.UnmarshalMap(result.Item, bk)
    if err != nil {
        return nil, err
    }

    return bk, nil
}

// Add a book record to DynamoDB.
func putItem(bk *book, stage string) error {
    fmt.Println("creating input")
    input := &dynamodb.PutItemInput{
        TableName: aws.String(getTableName(stage)),
        Item: map[string]*dynamodb.AttributeValue{
            "ISBN": {
                S: aws.String(bk.ISBN),
            },
            "Title": {
                S: aws.String(bk.Title),
            },
            "Author": {
                S: aws.String(bk.Author),
            },
        },
    }
    fmt.Println("starting to create item:", input)
    _, err := db.PutItem(input)
    fmt.Println("done to create item:", err)
    return err
}