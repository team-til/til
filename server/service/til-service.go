package service

import (
	"context"

	pb "github.com/team-til/til/server/_proto"
)

type TILServer struct {
}

func NewTILServer() *TILServer {
	return &TILServer{}
}

func (ts *TILServer) Ping(ctx context.Context, request *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Response: "pong"}, nil
}
