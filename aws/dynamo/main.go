package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"pos-graduacao-go-lang/aws/session"
)

var (
	dynamoDBClient *dynamodb.DynamoDB
)

const tableName = "products"

func init() {
	dynamoDBClient = dynamodb.New(session.AwsSession)
}

type Product struct {
	Id    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Price float64  `json:"price,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

func main() {
	listTables()
	scanItems()
	product := putItem()
	scanItems()
	getItem(product)
}

func getItem(product *Product) {
	key, err := dynamodbattribute.MarshalMap(map[string]string{
		"id": product.Id,
	})
	handleError(err)
	item, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		AttributesToGet:          nil,
		ConsistentRead:           nil,
		ExpressionAttributeNames: nil,
		Key:                      key,
		ProjectionExpression:     nil,
		ReturnConsumedCapacity:   nil,
		TableName:                aws.String(tableName),
	})
	handleError(err)
	fmt.Println(item)
}

func listTables() {
	tables, err := dynamoDBClient.ListTables(nil)
	handleError(err)
	fmt.Println("Tables:", tables)
}

func putItem() *Product {
	product := &Product{
		Id:    uuid.New().String(),
		Name:  "Produto " + uuid.New().String(),
		Price: 19.90,
		Tags:  []string{"Tag 1", "Tag 2", "Tag 3"},
	}
	item, err := dynamodbattribute.MarshalMap(product)
	handleError(err)
	_, err = dynamoDBClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	handleError(err)
	return product
}

func scanItems() {
	scan, err := dynamoDBClient.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	handleError(err)
	fmt.Println("Scan:", scan)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
