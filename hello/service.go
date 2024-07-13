package hello

import (
	"connectrpc.com/grpcreflect"
	"context"
	encore "encore.dev"
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"net/http"

	"github.com/sunit-kulkarni/encore-connect/gen/hello/v1/hellov1connect"
	"github.com/sunit-kulkarni/encore-connect/hello/workflow"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Use an environment-specific task queue so we can use the same
// Temporal Cluster for all cloud environments.
var (
	envName           = encore.Meta().Environment.Name
	greetingTaskQueue = envName + "-greeting"
)

//encore:service
type Service struct {
	routes http.Handler
	client client.Client
	worker worker.Worker
}

//encore:api public raw path=/*endpoint
func (s *Service) HelloService(w http.ResponseWriter, req *http.Request) {
	s.routes.ServeHTTP(w, req)
}

func initService() (*Service, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return nil, fmt.Errorf("create temporal client: %v", err)
	}

	w := worker.New(c, greetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(workflow.Greeting)
	w.RegisterActivity(workflow.ComposeGreeting)

	err = w.Start()
	if err != nil {
		c.Close()
		return nil, fmt.Errorf("start temporal worker: %v", err)
	}

	greeter := &Server{client: c}
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		hellov1connect.HelloServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	path, handler := hellov1connect.NewHelloServiceHandler(greeter)
	mux.Handle(path, handler)

	routes := h2c.NewHandler(mux, &http2.Server{})
	return &Service{routes: routes, client: c, worker: w}, nil
}

func (s *Service) Shutdown(force context.Context) {
	s.client.Close()
	s.worker.Stop()
}
