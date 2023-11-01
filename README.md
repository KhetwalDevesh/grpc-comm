# Steps to run the project :

1. Open a terminal and go to the root of the project in order to run the GRPC server
    * docker network create mynetwork
    * docker build -t server-totality-img -f server/Dockerfile .
    * docker rm -f server-totality && docker run -it --name server-totality -p 50051:50051 --network mynetwork server-totality-img
   
2. Open another terminal and again go to the root in order to run the GRPC client
    * docker build -t client-totality-img -f client/Dockerfile .
    * docker rm -f client-totality && docker run -it --name client-totality --network mynetwork client-totality-img
   
#### In order to re-run the client to test various cases keep re-running below command :
* docker rm -f client-totality && docker run -it --name client-totality --network mynetwork client-totality-img


# Steps to run for Make users ( optional )

1. Open a terminal and go to the root of the project in order to run the GRPC server
   * make network
   * make build-server
   * make run-server

2. Open another terminal and again go to the root in order to run the GRPC client
   * make build-client
   * make run-client

#### In order to re-run the client to test various cases keep re-running below command :
* make run-client


# Video link in order to see the demo :
https://www.loom.com/share/9e8887a1981742049506b73cafe17948?sid=10cd9873-f4e3-47f3-84f7-230b1d6f77c9