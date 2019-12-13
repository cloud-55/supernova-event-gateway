package resource

import (
	"os"
)

type Data struct {
	Provider string  `json:"provider"`
	Message  Message `json:"message,omitempty"`
}

type Message struct {
	TopicName string `json:"topic_name"`
	Payload   string `json:"payload"`
}

func (d *Data) GetMessageTopic() string {

	//TODO: This needs to get Account ID from iam:GetUser. This method returns an ARN and the AccountID should be extracted from that.
	if d.Provider == "AwsProvider" {
		return "arn:aws:sns:" + os.Getenv("AWS_REGION") + ":" + os.Getenv("AWS_ACCOUNT_ID") + ":" + d.Message.TopicName
	}

	return ""
}
