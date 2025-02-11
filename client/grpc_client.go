package main

import (
	pb "billsHotelService/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHotelServiceClient(conn)

	//ホテル一覧の取得
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetHotel(ctx, &pb.HotelRequest{})
	if err != nil {
		//log.Fatalf("could not get hotels: %v", err)
		panic(err)
	}
	fmt.Println("Hotels:", resp)
}
