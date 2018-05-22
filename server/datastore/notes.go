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

func (ds *NotesDatastore) GetNotePreviews(pageNum int, perPage int) ([]NoteDTO, error) {
	var offset int
	if pageNum == 1 {
		offset = 0
	} else {
		offset = ((perPage * pageNum) - perPage)
	}

	limit := perPage + 1

	var noteDTOs []NoteDTO

	if err := ds.db.Select(&noteDTOs, sqlGetNotePreviews, offset, limit); err != nil {
		return nil, err
	}
	return noteDTOs, nil
}

var sqlCreateNote = `INSERT INTO notes (name, note, filename)
VALUES (:name, :note, :filename)
RETURNING id, name, note, filename, created_at, updated_at`

var sqlGetNotePreviews = `SELECT id, name, note, filename, created_at, updated_at
FROM notes
ORDER BY created_at ASC
OFFSET $1
LIMIT $2`
