package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ajmalazeem/flurn/models"
	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req models.LoanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	// log.Println(req)
	// log.Println(req.CustomerName)

	return req, nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeCreateEndpoint(svc Loan) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req := request.(models.LoanRequest)
		svc.Create(req)

		return nil, nil
	}
}

func decodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {

	status := r.URL.Query().Get("status")
	log.Println(status)

	return nil, nil
}

func makeListEndpoint(svc Loan) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return svc.List()
	}
}

func MakeHandler(svc Loan) http.Handler {

	createHandler := httptransport.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
	)

	listEndpoint := httptransport.NewServer(
		makeListEndpoint(svc),
		decodeListRequest,
		encodeResponse,
	)

	router := mux.NewRouter()

	router.Methods(http.MethodPost).Path("/loans").Handler(createHandler)
	router.Methods(http.MethodGet).Path("/loans").Handler(listEndpoint)

	return router
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
