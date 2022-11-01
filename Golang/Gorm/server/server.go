package main

import (
	pb "Exercise/Gorm/proto"
	db "Exercise/Gorm/db"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

var (
	port= flag.Int("port", 50051, "The server port")
)


type EmployeeDetailsServer struct{
	pb.UnimplementedEmployeeManagementServer
	db *gorm.DB

} 

func(emp *EmployeeDetailsServer) CreateNewEmployee(ctx context.Context,in *pb.NewEmployee) (*pb.Employee,error){
	id:=uint(in.ManagerID)
	employee:=db.Employee{Name: in.Name,Designation: in.Designation,DepartmentID: uint(in.DepartmentID),ManagerID:&(id)}

	emp.db.Save(&employee)

	return &pb.Employee{Id:int32(employee.ID),Name: employee.Name},nil

}


func main(){
	username := flag.String("user_name", "postgres", "UserName String")
	password := flag.String("password", "Reddy@1234", "Password string")
	database := flag.String("db_name", "employees", "databaseName String")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:=grpc.NewServer()
	pb.RegisterEmployeeManagementServer(s,&EmployeeDetailsServer{db:db.NewClient(*username, *password, *database)})
	log.Printf("server listening")
	if err:= s.Serve(lis); err!=nil{
		log.Fatal("error")
	}
}

