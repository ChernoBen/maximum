package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"maximum/src/maximum_proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Falha ao conectar com servidor: %v\n", err)
	}
	defer conn.Close()
	c := maximum_proto.NewCalculatorServiceClient(conn)
	if c != nil {
		fmt.Println("Conexão com servidor criada com sucesso!")
	}
	doBiDiStreaming(c)
}

func doBiDiStreaming(c maximum_proto.CalculatorServiceClient) {
	fmt.Println("Req BiDi")
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Erro ao abrir stream: %v\n", err)
	}
	waitch := make(chan struct{})
	//envia goroutine
	go func() {
		numbers := []int32{1, 2, 3, 4, 5, 6, 8, 2}
		for _, num := range numbers {
			err := stream.Send(&maximum_proto.FindMaximumRequest{
				Number: num,
			})
			if err != nil {
				log.Fatalf("Falha ao enviar request: %v\n", err)
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	//recebe goroutine
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Falha ao obter response: %v\n", err)
				break
			}
			maximum := req.GetMaximum()
			fmt.Printf("Maximo atual: %v\n", maximum)
		}

		close(waitch)
	}()
	<-waitch
}
