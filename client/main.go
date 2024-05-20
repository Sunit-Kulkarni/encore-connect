package main

import (
	"context"
	"log"
	"net/http"

	"buf.build/gen/go/sunny-buf/connect-starter/connectrpc/go/greet/v1/greetv1connect"
	greetv1 "buf.build/gen/go/sunny-buf/connect-starter/protocolbuffers/go/greet/v1"

	"connectrpc.com/connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:4000",
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
