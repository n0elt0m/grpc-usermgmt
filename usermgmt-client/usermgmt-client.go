package main

import (
	"context"
	pb "github.com/noel/grpc-usermgmt/usermgmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Toni"] = 25
	new_users["Jhon"] = 29
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user : %v", err)
		}
		log.Printf(`User Details:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
