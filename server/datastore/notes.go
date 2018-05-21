package datastore

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // database/sql Postgres driver
)

type NotesDatastore struct {
	db *sqlx.DB
}

type NoteDTO struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Note      string    `db:"note"`
	FileName  string    `db:"filename"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewNotesDatastore(db *sqlx.DB) *NotesDatastore {
	return &NotesDatastore{
		db: db,
	}
}

func (ds *NotesDatastore) Create(ndto *NoteDTO) (*NoteDTO, error) {
	stmt, err := ds.db.PrepareNamed(sqlCreateNote)
	if err != nil {
		return nil, err
	}

	var dto NoteDTO
	if err := stmt.QueryRowx(ndto).StructScan(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

var sqlCreateNote = `INSERT INTO notes (name, note, filename)
VALUES (:name, :note, :filename)
RETURNING id, name, note, filename, created_at, updated_at`
