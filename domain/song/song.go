package song

import "github.com/google/uuid"

type Song struct {
	id     uuid.UUID
	name   string
	length int64

	recordID uuid.UUID
}

func NewSong(name string, length int64, recordID uuid.UUID) (Song, error) {
	return Song{
		name:     name,
		length:   length,
		recordID: recordID,
	}, nil
}

func (s Song) GetID() uuid.UUID {
	return s.id
}

func (s Song) GetName() string {
	return s.name
}

func (s Song) GetLength() int64 {
	return s.length
}

func (s Song) GetRecordID() uuid.UUID {
	return s.recordID
}

func (s *Song) SetID(id uuid.UUID) {
	s.id = id
}

func (s *Song) SetName(name string) {
	s.name = name
}

func (s *Song) SetLength(length int64) {
	s.length = length
}

func (s *Song) SetRecordID(recordID uuid.UUID) {
	s.recordID = recordID
}
