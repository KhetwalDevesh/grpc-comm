generate:
	protoc -Iproto --go_out=. --go_opt=module=github.com/KhetwalDevesh/project-totality --go-grpc_out=. --go-grpc_opt=module=github.com/KhetwalDevesh/project-totality proto/users.proto
