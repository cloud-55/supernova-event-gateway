package api

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/lmorais/supernova-message-gateway/db"

	"cloud.google.com/go/pubsub"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	//"gopkg.in/mgo.v2/bson"
	"log"
)

type Message struct {
	Provider string      `key:"provider" json: "provider" bson:"provider"`
	M        interface{} `key:"m" json: "m" bson:"m"`
}

func PublishMessage(ctx *routing.Context) error {

	var msg Message

	e := json.Unmarshal(ctx.PostBody(), &msg)

	if e != nil {
		fmt.Println(e)
	}

	res, err := json.Marshal(msg.M)
	if err != nil {
		log.Fatal(err)
	}

	if msg.Provider == "aws" {

		sess := session.Must(session.NewSession())
		svc := sns.New(sess)

		params := &sns.PublishInput{
			TopicArn: aws.String("arn:aws:sns:us-east-1:485164690107:spn-test"),
			Message:  aws.String(string(res)),
		}
		resp, err := svc.Publish(params)

		if err != nil {
			log.Fatal(err)
		}

	}

	if msg.Provider == "gcp" {

		client, err := pubsub.NewClient(ctx, "celestial-geode-129616")

		// Publish "hello world" on topic1.
		topic := client.Topic("projects/celestial-geode-129616/topics/repository-changes.default")
		res := topic.Publish(ctx, &pubsub.Message{
			Data: msg.M,
		})
		// The publish happens asynchronously.
		// Later, you can get the result from res:

		msgID, err := res.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println(resp)

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(res)
	ctx.SetStatusCode(fasthttp.StatusAccepted)

	return nil
}

func SubscribeMessage(ctx *routing.Context) error {

	var msg Message
	conn, err := db.Connect("", "")
	defer conn.Close()

	e := json.Unmarshal(ctx.PostBody(), &msg)

	if e != nil {
		fmt.Println(e)
	}

	res, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Collection("messages").Insert(msg)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(res)
	ctx.SetStatusCode(fasthttp.StatusAccepted)

	return nil
}
