package main

import (
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connetc: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUsers = make(map[string]int32)
	newUsers["Alice"] = 43
	newUsers["Bob"] = 30

	for name, age := range newUsers {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{
			Name: name,
			Age:  age,
		})

		if err != nil {
			log.Fatalf("Could not create user: %v", err)
		}
		log.Printf(`User Details:
Name: %s
Age: %d
Id: %d
`, r.GetName(), r.GetAge(), r.GetId())
	}

}
