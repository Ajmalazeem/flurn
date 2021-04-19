package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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
	var req models.ListRequest

	statuses := r.URL.Query().Get("status")
	if len(statuses) > 0 {
		req.Status = strings.Split(statuses, ",")
	}

	loanAmt := r.URL.Query().Get("loanAmountGreater")
	val, err := strconv.ParseFloat(loanAmt, 64)
	if err == nil {
		req.LoanAmountGreater = val
	}

	return req, nil
}

func makeListEndpoint(svc Loan) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.ListRequest)
		return svc.List(req)
	}
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req models.GetRequest
	vars := mux.Vars(r)
	id := vars["id"]

	var err error
	req.Id, err = strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeGetEndpoint(svc Loan) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req := request.(models.GetRequest)

		return svc.Get(req)
	}
}

func decodeApproveRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req models.ApproveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id := vars["id"]

	var err error
	req.Id, err = strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeApproveEndpoint(svc Loan) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req := request.(models.ApproveRequest)

		return svc.Approve(req)
	}
}

func decodeCancelRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req models.CancelRequest
	req.Status = "Cancelled"

	vars := mux.Vars(r)
	id := vars["id"]

	var err error
	req.Id, err = strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeCancelEndpoint(svc Loan) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req := request.(models.CancelRequest)

		return svc.Cancel(req)
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

	getEndpoint := httptransport.NewServer(
		makeGetEndpoint(svc),
		decodeGetRequest,
		encodeResponse,
	)

	ApproveEndpoint := httptransport.NewServer(
		makeApproveEndpoint(svc),
		decodeApproveRequest,
		encodeResponse,
	)

	CancelEndpoint := httptransport.NewServer(
		makeCancelEndpoint(svc),
		decodeCancelRequest,
		encodeResponse,
	)

	router := mux.NewRouter()

	router.Methods(http.MethodPost).Path("/loans").Handler(createHandler)
	router.Methods(http.MethodGet).Path("/loans").Handler(listEndpoint)
	router.Methods(http.MethodGet).Path("/loans/{id}").Handler(getEndpoint)
	router.Methods(http.MethodPatch).Path("/loans/{id}").Handler(ApproveEndpoint)
	router.Methods(http.MethodDelete).Path("/loans/{id}").Handler(CancelEndpoint)

	return router
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
