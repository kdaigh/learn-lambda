package main

import (
	"fmt"
	"math/rand"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID  uint64 `json: "id"`
	Max int    `json: "max"`
}

type Response struct {
	Message   string `json: "message"`
	Generated bool   `json: "generated"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message:   fmt.Sprintf("Random Number: %d", rand.Intn(request.Max)),
		Generated: true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
