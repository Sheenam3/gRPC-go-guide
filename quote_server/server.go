package main

import (

	
	"context"
	"google.golang.org/grpc"	
	pb "github.com/quotes/quotes"
	"net"
	"log"
	"io"

)

type QuoteGuideServer struct{
	pb.UnimplementedQuoteGuideServer
	savedThreeDayQuotes []*pb.ThreeQuoteRequest

}


func (s *QuoteGuideServer) OneDayQuote(ctx context.Context, quote *pb.OneQuoteRequest) (*pb.OneQuoteResponse, error) {
	
	dayquote := quote.GetQuoteone() //this will get the quote from client	
	log.Printf("Recieving quote from client for formatting: %s", dayquote)
	return &pb.OneQuoteResponse{Quote: "Quote of the Day: " + dayquote}, nil

}


func (s *QuoteGuideServer) ThreeDayQuote(quote *pb.ThreeQuoteRequest, stream pb.QuoteGuide_ThreeDayQuoteServer) error {

	for  i := 0; i<= 2; i++ {
		 
		err := stream.Send(&pb.ThreeQuoteResponse {
			Quote1Response: quote.GetQuote1Request(), 
			Quote2Response: quote.GetQuote2Request(),
			Quote3Response: quote.GetQuote3Request()})
		if err != nil { return err
		}
	}
		
	 /*for _, quotes := range stream {
	
		//log.Printf("Index %d: \n", i)	
		if err := stream.Send(quotes); err != nil {
			log.Printf("Returning err")
			return err
		}
	}*/

/*	for i, quotes := range s.savedThreeDayQuotes {
                log.Printf("Index %d: \n", i)
                if err := stream.Send(&pb.ThreeQuoteResponse{quotes}); err != nil {
                        return err
                }
        }*/

	return nil 
}


func (s *QuoteGuideServer) Quotes(stream pb.QuoteGuide_QuotesServer) error {

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ThreeQuoteResponse{
			Quote1Response: "response 1", 
			Quote2Response: "response 2",
			Quote3Response: "response 3"})
		}
		if err != nil {
			return err
		}

		log.Printf("end of the server %s", r.GetQuote3Request())
	}

	return nil
		

}


func (s *QuoteGuideServer) SayHi(stream pb.QuoteGuide_SayHiServer) error {

	n := 0
	for {
		err := stream.Send(&pb.OneQuoteResponse{
			Quote: "Hi, from the Server"})

		if err != nil {
			return err
		}

		hello, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf(" From Client: %s \n", hello )
	}

	return nil


}


func main() {

listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterQuoteGuideServer(srv, &QuoteGuideServer{})
	

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

