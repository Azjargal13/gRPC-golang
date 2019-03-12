package main

import (pb "github.com/GolangTutorials/1wayrpc/proto)

const (
	port = ":50051"
)
type server struct {}

func (s server) Max(srv pb.Math_MaxServer) error{
	log.Println("server started")
	var max int32
	ctx := srv.Context()
	
}