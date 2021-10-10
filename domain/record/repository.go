package record

import (
	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/song"
)

type RecordRepository interface {
	Get(uuid.UUID) (Record, error)
	Add(Record) error
	Update(*Record) error
	FindRecords() ([]Record, error)
	AddSong(uuid.UUID, *song.Song) error
}
