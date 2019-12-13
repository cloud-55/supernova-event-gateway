package api

import (
	"encoding/json"
	"fmt"
	"github.com/lmorais/supernova-message-gateway/resource"
	//"github.com/lmorais/supernova-message-gateway/db"

	"cloud.google.com/go/pubsub"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	//"gopkg.in/mgo.v2/bson"
	"context"
	//"errors"
	"log"
	"sync"
	//"sync/atomic"
)

var data resource.Data

func PublishMessage(ctx *routing.Context) error {

	e := json.Unmarshal(ctx.PostBody(), &data)

	if e != nil {
		fmt.Println(e)
	}

	response, err := NewProviderDispatch(data).DispatchMessageToProvider()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(fasthttp.StatusAccepted)

	return nil
}

func SubscribeMessage(rctx *routing.Context) error {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "celestial-geode-129616")

	if err != nil {
		log.Fatal(err)
	}

	// Consume 10 messages.
	var mu sync.Mutex
	received := 0
	sub := client.Subscription("test")
	cctx, cancel := context.WithCancel(ctx)
	errs := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 1 {
			cancel()
		}
	})
	if errs != nil {
		return errs
	}

	// var msg Message
	// conn, err := db.Connect("", "")
	// defer conn.Close()

	// e := json.Unmarshal(ctx.PostBody(), &msg)

	// if e != nil {
	// 	fmt.Println(e)
	// }

	// res, err := json.Marshal(msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = conn.Collection("messages").Insert(msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx.Response.Header.Set("Content-Type", "application/json")
	// ctx.Write(res)
	// ctx.SetStatusCode(fasthttp.StatusAccepted)

	return nil
}
