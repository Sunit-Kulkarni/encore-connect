package greet

import (
	"net/http"

	"encore.app/gen/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//encore:service
type Service struct {
	routes http.Handler
}

//encore:api public raw path=/greet.v1.GreetService/*endpoint
func (s *Service) GreetService(w http.ResponseWriter, req *http.Request) {
	s.routes.ServeHTTP(w, req)
}

func initService() (*Service, error) {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	routes := h2c.NewHandler(mux, &http2.Server{})
	return &Service{routes: routes}, nil
}
