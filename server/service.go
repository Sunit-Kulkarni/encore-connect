package server

import (
	"connectrpc.com/grpcreflect"
	"net/http"

	"github.com/sunit-kulkarni/encore-connect/gen/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//encore:service
type Service struct {
	routes http.Handler
}

//encore:api public raw path=/*endpoint
func (s *Service) GreetService(w http.ResponseWriter, req *http.Request) {
	s.routes.ServeHTTP(w, req)
}

func initService() (*Service, error) {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		greetv1connect.GreetServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	routes := h2c.NewHandler(mux, &http2.Server{})
	return &Service{routes: routes}, nil
}
