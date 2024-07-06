# Encore + Buf Schema Registry + Connect RPC

This project forms the basis to combine best in class technologies to make a rapidly iteratable and scalable backend.
The heart of the application uses the Encore runtime. 
Within the runtime, a raw http endpoint is used in order to allow for publicly accessible RPC apis built from the Connect RPC protocol. This repo expands upon this guide here: https://encore.dev/docs/how-to/grpc-connect.

## Repo structure

1) `proto` directory is the source of truth that defines the proto definitions and hence the APIs that will be publicly accessible. Through the use of `buf.yaml`, this source of truth is replicated to the Buf Schema Registry at https://buf.build/sunny-buf/connect-starter
2) `gen` directory is where the generated go code is placed which is defined by the `proto` definitions and the `buf.gen.yaml` configuration file.
3) `server` directory takes the generated go code from `gen` to make a backend implementation with business logic.
   - The `service.go` file wraps the server with gRPC server reflection. This will allow for clients like postman to recognize the methods that can be called on the endpoint rather than having to import a proto definition manually.
   - The `greet.go` file contains server business logic.
4) `client` directory has an example of calling the API when deployed locally. This utilizes a generated sdk from the Buf Schema Registry that is generated from first bullet point. Thanks to the Buf Schema Registry, almost any kind of client can call the API such as gRPC-Web, gRPC, and other language generated Connect SDKs. Integrating with a frontend React client could utilize any of the provided TypeScript/JavaScript generated SDKs.

## Running Locally

With the encore runtime installed on machine, in a terminal window run the command:

```shell
encore run
```

To call the API locally with client, in a separate terminal window run the command:
```shell
go run client/main.go
```