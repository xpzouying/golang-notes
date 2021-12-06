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
		svr = WithTimeUsedMiddleware(svr)
	}

	stringHTTPSvr := zykit.NewServer(
		makeUpperEndpoint(svr),
		decodeUpperRequest,
		encodeUpperResponse,
	)

	http.Handle("/upper", stringHTTPSvr)

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
