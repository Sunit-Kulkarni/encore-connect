module github.com/sunit-kulkarni/encore-connect

go 1.22.2

require (
	buf.build/gen/go/sunny-buf/connect-starter/connectrpc/go v1.16.2-20240520224700-9643f8537c64.1
	buf.build/gen/go/sunny-buf/connect-starter/protocolbuffers/go v1.34.1-20240520224700-9643f8537c64.1
	connectrpc.com/connect v1.16.2
	connectrpc.com/grpcreflect v1.2.0
	golang.org/x/net v0.27.0
	google.golang.org/protobuf v1.34.2
)

require golang.org/x/text v0.16.0 // indirect
