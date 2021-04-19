package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeCreateEndpoint(svc Loan) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}

func MakeHandler(svc Loan) http.Handler {

	createHandler := httptransport.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
	)

	router := mux.NewRouter()

	router.Methods(http.MethodPost).Path("/loans").Handler(createHandler)

	return router
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
