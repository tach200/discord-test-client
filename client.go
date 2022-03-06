package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/achmang/go-discord-chat/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1122", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server %v", err)
	}
	defer conn.Close()

	client := pb.NewDiscordMessageClient(conn)

	for {
		time.Sleep(10 * time.Second)
		resp, err := client.SendChanMessage(context.Background(), &pb.MessageChannel{
			Subject: "Test Subject",
			Content: "Some String Content goes here",
		})
		if err != nil {
			log.Fatalf("send channel message error : %v", err)
		}
		fmt.Println(resp)
	}
}
