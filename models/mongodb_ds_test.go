package models

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type fakeDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	payload      []string
	feedbackData FeedbackData
}

func (fakeDynamo *fakeDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	for _, v := range fakeDynamo.payload {
		if v == *input.Item["Timestamp"].S {
			return nil, awserr.New(
				dynamodb.ErrCodeConditionalCheckFailedException,
				"The conditional request failed",
				errors.New("status code: 400, request id: 8VD67UJTFVS4FMOH8UMEV65M6BVV4KQNSO5AEMVJF66Q9ASUAAJG"),
			)
		}
	}
	feedback, ok := input.Item["Feedback"]
	if ok && feedback.S != nil && len(*feedback.S) == 0 {
		return nil, awserr.New(
			"ValidationException",
			"One or more parameter values were invalid: An AttributeValue may not contain an empty string",
			nil)
	}
	fakeDynamo.payload = append(fakeDynamo.payload, *input.Item["Timestamp"].S)
	return nil, nil
}

func TestDynamoDBRegistry_Put(t *testing.T) {
	f := FeedbackData{Timestamp: "12-12-2012", Feedback: "Very helpful test"}
	db := fakeDynamoDB{payload: []string{}, feedbackData: f}
	var feedbackDS FeedbackDS = dynamoDBDS{&db}

	err := feedbackDS.Put(f)
	if err != nil {
		t.Fatalf("FAIL, Error received: %v", err)
	}
	if db.payload[0] != "12-12-2012" {
		t.Fatalf("FAIL, Feeback not saved in DynamoDB")
	}
}
