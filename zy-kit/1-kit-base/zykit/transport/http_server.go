package transport

import (
	"context"
	"log"
	"net/http"

	"zykit/endpoint"
)

type Server struct {
	ep  endpoint.Endpoint
	dec DecodeRequestFunc
	enc EncodeResponseFunc
}

func NewServer(
	ep endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
) *Server {

	return &Server{
		ep,
		dec,
		enc,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	req, err := s.dec(ctx, r)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	response, err := s.ep(ctx, req)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	if err := s.enc(ctx, w, response); err != nil {
		log.Printf("error: %v", err)
		return
	}
}

// DecodeRequestFunc 将http.Request请求，转换为服务业务的请求
type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

// EncodeResponseFunc 将服务的业务响应写入到http.Response中
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error
