package api

import (
	"fmt"
	"github.com/lmorais/supernova-message-gateway/aws"
	"log"
	//"github.com/lmorais/supernova-message-gateway/gcloud"
	"github.com/lmorais/supernova-message-gateway/resource"
	"gopkg.in/oleiade/reflections.v1"
)

type providerFunction func(data resource.Data) (interface{}, error)

type providerDispatch struct {
	Data                resource.Data
	AwsProvider         providerFunction
	GoogleCloudProvider providerFunction
}

func NewProviderDispatch(data resource.Data) (pd *providerDispatch) {

	return &providerDispatch{
		Data: data,
		AwsProvider: func(data resource.Data) (interface{}, error) {

			response, err := aws.PublishMessage(data)
			if err != nil {
				log.Fatal(err)
			}

			return response, nil
		},
		GoogleCloudProvider: func(data resource.Data) (interface{}, error) {
			fmt.Println("GoogleCloudProvider called, but needs to be implemented !!")
			return nil, nil
		},
	}
}

func (pd *providerDispatch) DispatchMessageToProvider() (interface{}, error) {

	getFunction, err := reflections.GetField(pd, pd.Data.Provider)

	if err != nil {
		log.Fatal(err)
	}

	fn := getFunction.(providerFunction)

	response, err := fn(data)

	if err != nil {
		log.Fatal(err)
	}

	return response, nil

}
