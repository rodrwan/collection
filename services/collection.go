package services

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/record"
	"github.com/rodrwan/collection/domain/record/memory"
	"github.com/rodrwan/collection/domain/record/postgres"
	"github.com/rodrwan/collection/domain/song"
	smemory "github.com/rodrwan/collection/domain/song/memory"
)

var (
	ErrInvalidType = errors.New("Invalid record type")
)

// ICollectionService ...
type ICollectionService interface {
	// AddRecord ...
	AddRecord(name string, kind string) (record.Record, error)
	// FindRecord ...
	FindRecord(id string) (record.Record, error)
	// AddSongToRecord ...
	AddSongToRecord(record record.Record, name string, length int64) (record.Record, error)
	// FindAllRecord...
	FindAllRecord() ([]record.Record, error)
}

// CollectionConfiguration ...
type CollectionConfiguration func(cs *CollectionService) error

// CollectionService ...
type CollectionService struct {
	records record.RecordRepository
	songs   song.SongRepository
}

// WithRecordMemoryRepository ...
func WithRecordMemoryRepository() CollectionConfiguration {
	return func(os *CollectionService) error {
		ctx := context.Background()
		mem, err := memory.New(ctx)
		if err != nil {
			log.Fatal(err)
		}

		os.records = mem
		return nil
	}
}

// WithRecordPostgresRepository ...
func WithRecordPostgresRepository(connectionString string) CollectionConfiguration {
	return func(os *CollectionService) error {
		pg, err := postgres.New(context.Background(), connectionString)
		if err != nil {
			return err
		}

		os.records = pg
		return nil
	}
}

// WithSongMemoryRepository ...
func WithSongMemoryRepository() CollectionConfiguration {
	return func(os *CollectionService) error {
		ctx := context.Background()
		mem, err := smemory.New(ctx)
		if err != nil {
			log.Fatal(err)
		}

		os.songs = mem
		return nil
	}
}

// // WithSongPostgresRepository ...
// func WithSongPostgresRepository(connectionString string) CollectionConfiguration {
// 	return func(os *CollectionService) error {
// 		pg, err := postgres.New(context.Background(), connectionString)
// 		if err != nil {
// 			return err
// 		}

// 		os.records = pg
// 		return nil
// 	}
// }

// NewCollectionService ...
func NewCollectionService(cfgs ...CollectionConfiguration) (*CollectionService, error) {
	cs := &CollectionService{}

	for _, cfg := range cfgs {
		err := cfg(cs)
		if err != nil {
			return nil, err
		}
	}

	return cs, nil
}

// AddRecord ...
func (cs *CollectionService) AddRecord(name string, kind string) (record.PublicRecord, error) {
	rec, err := record.NewRecord(name, kind)
	if err != nil {
		return (&record.Record{}).ToPublic(), err
	}

	switch kind {
	case "vinyl":
		if err := cs.records.Add(rec); err != nil {
			return (&record.Record{}).ToPublic(), err
		}

		return rec.ToPublic(), nil
	case "mp3":
		if err := cs.records.Add(rec); err != nil {
			return (&record.Record{}).ToPublic(), err
		}

		return rec.ToPublic(), nil
	}

	return (&record.Record{}).ToPublic(), ErrInvalidType
}

// FindRecord ...
func (cs *CollectionService) FindRecord(id string) (record.PublicRecord, error) {
	uuidId := uuid.MustParse(id)
	rec, err := cs.records.Get(uuidId)
	if err != nil {
		return (&record.Record{}).ToPublic(), err
	}

	return rec.ToPublic(), nil
}

// AddSongToRecord ...
func (cs *CollectionService) AddSongToRecord(record *record.Record, name string, length int64) error {
	s, err := song.NewSong(name, length, record.GetID())
	if err != nil {
		return err
	}

	return record.AddSong(&s)
}

// FindAllRecord ...
func (cs *CollectionService) FindAllRecord() ([]record.PublicRecord, error) {
	records, err := cs.records.FindRecords()
	if err != nil {
		return []record.PublicRecord{}, err
	}
	return record.ToPublicArray(records), nil
}
