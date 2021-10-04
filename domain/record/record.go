package record

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/song"
)

var (
	ErrMissingValues = errors.New("missing value")
)

type Record struct {
	id    uuid.UUID
	name  string
	kind  string
	songs []*song.Song
}

type PublicRecord struct {
	ID    uuid.UUID    `json:"id,omitempty"`
	Name  string       `json:"name,omitempty"`
	Kind  string       `json:"kind,omitempty"`
	Songs []*song.Song `json:"songs,omitempty"`
}

func (r *Record) ToPublic() PublicRecord {
	return PublicRecord{
		ID:   r.GetID(),
		Name: r.GetName(),
		Kind: r.GetKind(),
	}
}

func (pr *PublicRecord) ToRecord() *Record {
	r := &Record{}

	r.SetID(pr.ID)
	r.SetName(pr.Name)
	r.SetKind(pr.Kind)

	return r
}

func ToPublicArray(records []Record) []PublicRecord {
	var rr []PublicRecord

	for _, r := range records {
		rr = append(rr, r.ToPublic())
	}
	return rr
}

func NewRecord(name, kind string) (Record, error) {
	if name == "" {
		return Record{}, ErrMissingValues
	}

	return Record{
		id:    uuid.New(),
		name:  name,
		kind:  kind,
		songs: make([]*song.Song, 0),
	}, nil
}

func (r *Record) AddSong(song *song.Song) error {
	r.songs = append(r.songs, song)

	return nil
}

func (r *Record) SetID(id uuid.UUID) {
	r.id = id
}

func (r *Record) SetName(name string) {
	r.name = name
}

func (r *Record) SetKind(kind string) {
	r.kind = kind
}

func (r Record) GetID() uuid.UUID {
	return r.id
}

func (r Record) GetName() string {
	return r.name
}

func (r Record) GetKind() string {
	return r.kind
}

func (r Record) GetSongs() []*song.Song {
	return r.songs
}
