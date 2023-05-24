package main

import s "github.com/krishnakantha1/expenseTrackerAuth/server"

func main() {
	server := s.Server{}
	server.Init()
}

//protoc --go_out=../pb --go-grpc_out=../pb Service.proto
