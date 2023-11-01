generate:
	protoc -Iproto --go_out=. --go_opt=module=github.com/KhetwalDevesh/project-totality --go-grpc_out=. --go-grpc_opt=module=github.com/KhetwalDevesh/project-totality proto/users.proto
	go build -o bin/server ./server
	go build -o bin/client ./client

network:
	docker network create mynetwork

build-server:
	docker build -t server-totality-img -f server/Dockerfile .

run-server:
	docker rm -f server-totality && docker run -it --name server-totality -p 50051:50051 --network mynetwork server-totality-img

build-client:
	docker build -t client-totality-img -f client/Dockerfile .

run-client:
	docker rm -f client-totality && docker run -it --name client-totality --network mynetwork client-totality-img

