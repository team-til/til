package mapper

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/team-til/til/server/_proto"
	"github.com/team-til/til/server/datastore"
)

func ToNoteDTO(note *pb.Note) *datastore.NoteDTO {
	return &datastore.NoteDTO{
		Name:     note.Name,
		FileName: note.Filename,
		Note:     note.Note,
	}
}

func FromNoteDTO(ndto *datastore.NoteDTO) *pb.Note {
	return &pb.Note{
		Id:        ndto.ID,
		Name:      ndto.Name,
		Filename:  ndto.FileName,
		Note:      ndto.Note,
		CreatedAt: genPbTimestamp(ndto.CreatedAt),
		UpdatedAt: genPbTimestamp(ndto.UpdatedAt),
	}
}

func ToNotePreview(ndto *datastore.NoteDTO) *pb.NotePreview {
	var notePreview string
	if len(ndto.Note) > 400 {
		notePreview = ndto.Note[0:399]
	} else {
		notePreview = ndto.Note
	}

	return &pb.NotePreview{
		Id:          ndto.ID,
		Name:        ndto.Name,
		Filename:    ndto.FileName,
		NotePreview: notePreview,
		CreatedAt:   genPbTimestamp(ndto.CreatedAt),
		UpdatedAt:   genPbTimestamp(ndto.UpdatedAt),
	}
}

func genPbTimestamp(time time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{Seconds: time.Unix(), Nanos: 0}
}
