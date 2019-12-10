package api

import (
	"github.com/lmorais/supernova-message-gateway/aws"
	"log"
	//"github.com/lmorais/supernova-message-gateway/gcloud"
	"github.com/lmorais/supernova-message-gateway/resource"
)

const (
	AWS_CLOUD_PROVIDER    = "aws"
	GOOGLE_CLOUD_PROVIDER = "gcloud"
)

func dispatchMessageToProvider(data resource.Data) (interface{}, error) {

	if data.Provider == AWS_CLOUD_PROVIDER {
		response, err := aws.PublishMessage(data)
		if err != nil {
			log.Fatal(err)
		}

		return response, nil
	}

	return nil, nil

}
