package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"time"	
	"google.golang.org/grpc"
	"log"
	 pb "github.com/quotes/quotes"
	
)


func printThree(client pb.QuoteGuideClient, quotes *pb.ThreeQuoteRequest) {

	//log.Printf("Looking for Quotes within %v", quotes)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ThreeDayQuote(ctx, quotes)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		quot, err := stream.Recv()
		log.Printf("Inside PrintThree only")
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(quot)
	}

}


func printThreeAgain(client pb.QuoteGuideClient, quotes *pb.ThreeQuoteRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Quotes(ctx)
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(quotes)
		if err != nil {
			return err
		}
		
	}

	resp, err := stream.CloseAndRecv()
		if err != nil {
                        return err
                }

	log.Printf("resp: quote1: %s", resp)

	return nil

}


func sayHelloBi(client pb.QuoteGuideClient, hi *pb.OneQuoteRequest) error {

stream, err := client.SayHi(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 2; n++ {
		err = stream.Send(hi)
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: Client send : %s /n Server Responded: %s", hi, resp)
	}

	stream.CloseSend()

	return nil



}


func main() {

	flag.Parse()
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

	client := pb.NewQuoteGuideClient(conn)


	quoteone := "quote1 is here"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        quot, err := client.OneDayQuote(ctx, &pb.OneQuoteRequest{Quoteone: quoteone})
        if err != nil {
                fmt.Println("Error : ", err)
        }
        	log.Printf(" %s",quot.GetQuote())

	printThree(client,&pb.ThreeQuoteRequest{Quote1Request: "quote 1", Quote2Request: "quote 2", Quote3Request: "quote 3" })
	printThreeAgain(client,&pb.ThreeQuoteRequest{Quote1Request: "quote 1 streaming client", Quote2Request: "quote 2 streaming client", Quote3Request: "quote 3 streaming client" })
	sayHelloBi(client,&pb.OneQuoteRequest{Quoteone: "Hello, I am Client"})
}
