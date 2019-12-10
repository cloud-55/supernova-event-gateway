package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/lmorais/supernova-message-gateway/resource"
	"log"
)

func PublishMessage(data resource.Data) (interface{}, error) {

	session := NewSession()

	snsService := sns.New(session.Session)

	publishInput := &sns.PublishInput{
		TopicArn: aws.String(data.GetMessageTopic()),
		Message:  aws.String(data.Message.Payload),
	}

	response, err := snsService.Publish(publishInput)

	if err != nil {
		log.Fatal(err)
	}

	return response, nil

}
