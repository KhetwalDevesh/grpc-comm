# Use an official Golang runtime as a parent image
FROM golang:1.21.3-bookworm

# Set the working directory to /go/src/app
WORKDIR /go/src/app

RUN go mod init github.com/KhetwalDevesh/project-totality

# Install the Protocol Buffers compiler (protoc)
RUN apt-get update &&  apt install -y protobuf-compiler

# Install the protoc-gen-go plugin using 'go get'
RUN apt-get update && apt-get install -y protoc-gen-go
RUN apt-get update && apt-get install -y protoc-gen-go-grpc

RUN go get google.golang.org/grpc
RUN go get google.golang.org/grpc/codes
RUN go get google.golang.org/grpc/status

ENV GOBIN /go/bin

ENV PATH="$PATH:/go/bin"

# Copy the proto file from the root directory into the container
COPY proto/users.proto ./proto/

# Copy your server code
COPY server/ .

# Copy the data folder into the container
COPY data/ ./data/

# Generate Go code from the proto file
RUN protoc -Iproto --go_out=. --go_opt=module=github.com/KhetwalDevesh/project-totality --go-grpc_out=. --go-grpc_opt=module=github.com/KhetwalDevesh/project-totality proto/users.proto

# Build your Go server application
RUN go build -o bin/server . 2>&1

# Expose port 50051 for gRPC
EXPOSE 50051

# Command to run the server
CMD ["./bin/server"]
