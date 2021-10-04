package song

import (
	"github.com/google/uuid"
)

type SongRepository interface {
	Get(uuid.UUID) (Song, error)
	Add(Song) error
	Update(*Song) error
	FindSongsByRecord(uuid.UUID) ([]Song, error)
}
