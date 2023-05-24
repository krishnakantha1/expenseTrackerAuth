package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/krishnakantha1/expenseTrackerAuth/pb"
)

type Server struct {
	authServer *AuthServer
}

func (s *Server) Init() {
	log.Println("Server starting...")

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if len(port) == 1 {
		port = ":8080"
	}

	log.Println("Server will listen at port", port)
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Could not listen to port ", port, "\nerror is :", err)
	}

	log.Println("Creating GRPC server")
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, s.authServer)

	log.Println("GRPC server listening at port", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Could not listen to port", port, "\nerror is :", err)
	}

}
