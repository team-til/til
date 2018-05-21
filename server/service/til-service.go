package service

import (
	"context"

	pb "github.com/team-til/til/server/_proto"
	"github.com/team-til/til/server/datastore"
	"github.com/team-til/til/server/mapper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotesDatastore interface {
	Create(*datastore.NoteDTO) (*datastore.NoteDTO, error)
}

type TILServer struct {
	nds NotesDatastore
}

func NewTILServer(nds NotesDatastore) *TILServer {
	return &TILServer{
		nds: nds,
	}
}

func (ts *TILServer) Ping(ctx context.Context, request *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Response: "pong"}, nil
}

func (ts *TILServer) CreateNote(ctx context.Context, request *pb.CreateNoteRequest) (*pb.Note, error) {
	reqDTO := mapper.ToNoteDTO(request.Note)

	noteDTO, err := ts.nds.Create(reqDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to Create. Err: %+v", err)
	}
	return mapper.FromNoteDTO(noteDTO), nil
}
