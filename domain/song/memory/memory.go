package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/song"
)

type MemoryRepository struct {
	songs []memorySong

	sync.Mutex
}

type memorySong struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Length   int64     `db:"length"`
	RecordID uuid.UUID `db:"record_id"`
}

func NewFromSong(s song.Song) memorySong {
	return memorySong{
		ID:       s.GetID(),
		Name:     s.GetName(),
		Length:   s.GetLength(),
		RecordID: s.GetRecordID(),
	}
}

func (pr memorySong) ToSong() song.Song {
	s := song.Song{}

	s.SetID(pr.ID)
	s.SetName(pr.Name)
	s.SetLength(pr.Length)
	s.SetRecordID(pr.RecordID)

	return s
}

// Create a new mongodb repository
func New(ctx context.Context) (*MemoryRepository, error) {
	return &MemoryRepository{
		songs: make([]memorySong, 0),
	}, nil
}

func (mr *MemoryRepository) Get(id uuid.UUID) (song.Song, error) {
	mr.Lock()
	defer mr.Unlock()

	for _, s := range mr.songs {
		fmt.Println(s.ID, id)
		if s.ID == id {
			return s.ToSong(), nil
		}
	}

	// Convert to aggregate
	return song.Song{}, song.ErrSongNotFound
}

func (mr *MemoryRepository) Add(s song.Song) error {
	mr.Lock()
	defer mr.Unlock()

	internal := NewFromSong(s)
	mr.songs = append(mr.songs, internal)

	return nil
}

func (mr *MemoryRepository) FindRecords() ([]song.Song, error) {
	mr.Lock()
	defer mr.Unlock()

	var ss []song.Song
	for _, s := range mr.songs {
		ss = append(ss, s.ToSong())
	}

	return ss, nil
}

func (mr *MemoryRepository) Update(s *song.Song) error {
	mr.Lock()
	defer mr.Unlock()
	panic("to implement")
}

func (mr *MemoryRepository) FindSongsByRecord(id uuid.UUID) ([]song.Song, error) {
	mr.Lock()
	defer mr.Unlock()

	var ss []song.Song
	for _, s := range mr.songs {
		if s.RecordID == id {
			ss = append(ss, s.ToSong())
		}
	}

	return ss, nil
}
