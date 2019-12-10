package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/lmorais/supernova-message-gateway/config"
)

type Provider struct {
	AccountID         string `key:"account_id" json: "account_id" bson:"account_id"`
	AccessKeyID       string `key:"access_key_id" json: "access_key_id" bson:"access_key_id"`
	SecretAccessKeyID string `key:"secret_access_key_id" json: "secret_access_key_id" bson:"secret_access_key_id"`
	Session           *session.Session
}

func NewSession() *Provider {
	session := session.Must(session.NewSession(&aws.Config{Region: aws.String(config.Fetch().ProviderRegion)}))

	return &Provider{
		Session: session,
	}
}
