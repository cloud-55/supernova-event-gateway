package main

import (
	"github.com/joho/godotenv"
	"github.com/lmorais/supernova-message-gateway/api"
	"github.com/lmorais/supernova-message-gateway/config"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routing.New()

	router.Post("/api/v1/messages", func(ctx *routing.Context) error {
		api.PublishMessage(ctx)
		return nil
	})

	router.Post("/api/v1/subscriptions", func(ctx *routing.Context) error {
		api.SubscribeMessage(ctx)
		return nil
	})

	if err := fasthttp.ListenAndServe(config.Fetch().APIPort, router.HandleRequest); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}

}
