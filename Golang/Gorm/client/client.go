package main

import (
	pb "Exercise/Gorm/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const ( address="localhost:50051")

func main(){
	conn,err:=grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())
	if err!=nil{
		log.Fatalf("did not connect : %v",err)
	}
	defer conn.Close()
	
	NewEmployee(conn)
}

func NewEmployee(conn *grpc.ClientConn){

	c:=pb.NewEmployeeManagementClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	response,err:=c.CreateNewEmployee(ctx,&pb.NewEmployee{Name: "Vamshi",Designation: "manager",DepartmentID:2 })
	if err!=nil{
		log.Fatal("Error creating employee")
	}
	log.Print(response.Name,response.Id)
}