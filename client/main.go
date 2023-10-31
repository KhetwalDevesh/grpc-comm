package main

import (
	"bufio"
	"context"
	"fmt"
	pb "github.com/KhetwalDevesh/project-totality/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"strings"
)

var address = "localhost:50051"

func getUserDetailsByID(client pb.UserServiceClient) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the user ID : \n(Enter User ID from 1 to 20 for better result)\n")
	userIdString, _ := reader.ReadString('\n')
	userIdString = strings.TrimSpace(userIdString)

	userID, err := strconv.Atoi(userIdString)
	if err != nil {
		log.Fatalf("Invalid user ID : %v", err)
	}

	user, err := client.GetUserById(context.Background(), &pb.UserRequest{Id: int32(userID)})
	if err != nil {
		log.Fatalf("Error calling GetUserByID : %v", err)
	}
	fmt.Printf("\nUser Details by ID:\n\n%+v\n\n", user)
}

func getUsersDetailsByIDs(client pb.UserServiceClient) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a list of User IDs separated by commas (e.g., 1,2,3): \n(Enter User IDs from 1 to 20 for better results)\n")
	idsString, _ := reader.ReadString('\n')
	idsString = strings.TrimSpace(idsString)
	ids := strings.Split(idsString, ",")

	var userIDs []int32

	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("Invalid user ID : %v, Only enter numbers from 1 to 20 as User IDs", err)
		}
		userIDs = append(userIDs, int32(id))
	}

	users, err := client.GetUsersByIds(context.Background(), &pb.UsersRequest{Ids: userIDs})
	if err != nil {
		log.Fatalf("Error calling GetUsersByIDs : %v", err)
	}
	fmt.Println("\nUser Details by IDs : ")
	for _, user := range users.Users {
		fmt.Printf("\n%+v\n", user)
	}
	if len(userIDs) != len(users.Users) {
		fmt.Println("\nAll Users with entered User Id's not found. Only enter User Id's from 1 to 20 for better results.\n")
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose an option :\n 1 for getting a User by ID\n 2 for getting Users by ID's\n")
	optionStr, _ := reader.ReadString('\n')
	optionStr = strings.TrimSpace(optionStr)

	option, err := strconv.Atoi(optionStr)
	if err != nil {
		log.Fatalf("Invalid option: %v", err)
	}

	if option == 1 {
		getUserDetailsByID(client)
	} else if option == 2 {
		getUsersDetailsByIDs(client)
	} else {
		log.Fatal("Invalid option. Please choose 1 or 2.")
	}
}
