package service

import (
	"context"
	"strconv"

	"github.com/apex/log"
	pb "github.com/team-til/til/server/_proto"
	"github.com/team-til/til/server/datastore"
	"github.com/team-til/til/server/mapper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	notesDefaultPerPage = 20
	notesDefaultPageNum = 1
)

type NotesDatastore interface {
	Create(*datastore.NoteDTO) (*datastore.NoteDTO, error)
	GetNotePreviews(pageNum int, perPage int) ([]datastore.NoteDTO, error)
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

func (ts *TILServer) GetNotePreviews(ctx context.Context, request *pb.GetNotePreviewsRequest) (*pb.GetNotePreviewsResponse, error) {
	perPage := int(request.Pagination.GetPerPage())
	if perPage == 0 {
		perPage = notesDefaultPerPage
	}

	pageNum := int(request.Pagination.GetPageNumber())
	if pageNum == 0 {
		pageNum = notesDefaultPageNum
	}

	noteDTOs, err := ts.nds.GetNotePreviews(pageNum, perPage)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to GetNotePreviews. Err: %+v", err)
	}

	notePreviews := make([]*pb.NotePreview, len(noteDTOs))

	for i, noteDTO := range noteDTOs {
		notePreviews[i] = mapper.ToNotePreview(&noteDTO)
	}

	fullPage := len(notePreviews) == perPage
	responsePagination := pb.PaginationResponse{PerPage: int64(perPage), PageNumber: int64(pageNum), PagesRemaining: strconv.FormatBool(fullPage)}
	response := pb.GetNotePreviewsResponse{}
	response.NotePreviews = notePreviews
	response.Pagination = &responsePagination
	log.Infof("%+v\n", response)

	return &response, nil
}
