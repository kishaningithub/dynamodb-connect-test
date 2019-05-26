package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func main()  {
	log.Println("Using AWS instance")
	runTest(awsInstance())
	log.Println()
	log.Println("Using local instance")
	runTest(localInstance())
}

func awsInstance() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession())
	dynamodb.New(sess)
	return dynamodb.New(sess)
}

func localInstance() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession())
	return dynamodb.New(sess, aws.NewConfig().WithEndpoint("http://localhost:4569"))
}

func runTest(db *dynamodb.DynamoDB) {
	log.Println("Request 1: Start")
	output, e := db.ListTables(&dynamodb.ListTablesInput{})
	handleError(e)
	log.Printf("Request 1: Complete: Total no. of tables: %d ", len(output.TableNames))
	log.Println("Request 2: Start")
	output, e = db.ListTables(&dynamodb.ListTablesInput{})
	handleError(e)
	log.Printf("Request 2: Complete: Total no. of tables: %d ", len(output.TableNames))
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}



