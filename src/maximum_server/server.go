package main

import (
	"fmt"
	"io"
	"log"
	"maximum/src/maximum_proto"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	maximum_proto.UnimplementedCalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v\n", err)
	}
	log.Printf("Server running...\n")
	s := grpc.NewServer()
	maximum_proto.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao setar server: %v\n", err)
	}
}

func (*server) FindMaximum(stream maximum_proto.CalculatorService_FindMaximumServer) error {
	fmt.Println("Verifica maximo na stream")
	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Falha ao ler request %v\n", err)
			return err
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&maximum_proto.FindMaximumResponse{
				Maximum: maximum,
			})
			if sendErr != nil {
				log.Fatalf("Erro enquanto enviada resposta: %v\n", err)
				return err
			}
		}
	}
}
