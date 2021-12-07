package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"zykit"
)

type StringService interface {
	Upper(s string) string

	Count(s string) int
}

func main() {
	var port string
	{
		flag.StringVar(&port, "port", ":8080", "port of http server")
		flag.Parse()
	}

	var svr StringService
	{
		svr = &simpleStringServer{}
	}

	// add middleware, and enable request counter for metrics
	{
		counter := zykit.NewCounter(zykit.CounterOpts{
			Name: "request_count",
			Help: "Counter of request",
		})
		svr = &instrumentMiddleware{*counter, svr}
	}

	var upperEndpoint zykit.Endpoint
	{
		upperEndpoint = makeUpperEndpoint(svr)
		upperEndpoint = zykit.WithTimeUsed(upperEndpoint)
	}

	var countEndpoint zykit.Endpoint
	{
		countEndpoint = makeCountEndpoint(svr)
		countEndpoint = zykit.WithTimeUsed(countEndpoint)
	}

	stringUpperServer := zykit.NewServer(
		upperEndpoint,
		decodeUpperRequest,
		encodeUpperResponse,
	)

	stringCountServer := zykit.NewServer(
		countEndpoint,
		decodeCountRequest,
		encodeCountResponse,
	)

	http.Handle("/upper", stringUpperServer)
	http.Handle("/count", stringCountServer)

	http.Handle("/metrics", zykit.NewPromHTTPHandler())

	log.Printf("listen on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func makeUpperEndpoint(svc StringService) zykit.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(upperRequest)
		result := svc.Upper(req.S)
		return result, nil
	}
}

func makeCountEndpoint(svc StringService) zykit.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(countRequest)
		result := svc.Count(req.S)
		return result, nil
	}
}

type upperRequest struct {
	S string `json:"s"`
}

func decodeUpperRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var request upperRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeUpperResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	return json.NewEncoder(w).Encode(response)
}

type countRequest struct {
	S string `json:"s"`
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeCountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	return json.NewEncoder(w).Encode(response)
}
