package mock

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/song"
)

// MOCKS
type MockSongRepository struct {
	WithError bool
	RecordId  uuid.UUID
}

func (mrr MockSongRepository) Get(id uuid.UUID) (song.Song, error) {
	if mrr.WithError {
		return song.Song{}, errors.New("something went wrong")
	}

	return song.NewSongWithID(id, "lala", 100, mrr.RecordId)
}

func (mrr MockSongRepository) Add(rec song.Song) error {
	if mrr.WithError {
		return errors.New("something went wrong")
	}

	return nil
}

func (mrr MockSongRepository) Update(rec *song.Song) error {
	if mrr.WithError {
		return errors.New("something went wrong")
	}

	return nil
}

func (mrr MockSongRepository) FindRecords() ([]song.Song, error) {
	if mrr.WithError {
		return []song.Song{}, errors.New("something went wrong")
	}

	return []song.Song{}, nil
}
