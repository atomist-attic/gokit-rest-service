package mymicroservice

import (
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

type mymicroservicecontract interface {
	HelloWorld() (string, error)
}

type mymicroservice struct{}

// AddServices sets up and starts the service.
func AddServices() {
	ctx := context.Background()
	svc := mymicroservice{}

	mymicroserviceHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/", mymicroserviceHandler)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
