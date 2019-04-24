package api

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lmorais/supernova-message-gateway/db"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func CreateConstellation(ctx *routing.Context) error {
	go createConstellation()

	var cst models.Constellation

	e := json.Unmarshal(ctx.PostBody(), &cst)

	if e != nil {
		fmt.Println(e)
	}

	res, err := json.Marshal(cst)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(res)
	ctx.SetStatusCode(fasthttp.StatusAccepted)

	return nil
}
