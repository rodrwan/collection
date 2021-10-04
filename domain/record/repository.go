package record

import (
	"github.com/google/uuid"
)

type RecordRepository interface {
	Get(uuid.UUID) (Record, error)
	Add(Record) error
	Update(*Record) error
	FindRecords() ([]Record, error)
}
