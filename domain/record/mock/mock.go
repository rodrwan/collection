package mock

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/record"
)

// MOCKS
type MockRecordRepository struct {
	WithError bool
	RecordId  uuid.UUID
}

func (mrr MockRecordRepository) Get(id uuid.UUID) (record.Record, error) {
	if mrr.WithError {
		return record.Record{}, errors.New("something went wrong")
	}

	return record.NewRecordWithID(id, "lala", "vinyl")
}

func (mrr MockRecordRepository) Add(rec record.Record) error {
	if mrr.WithError {
		return errors.New("something went wrong")
	}

	return nil
}

func (mrr MockRecordRepository) Update(rec *record.Record) error {
	if mrr.WithError {
		return errors.New("something went wrong")
	}

	return nil
}

func (mrr MockRecordRepository) FindRecords() ([]record.Record, error) {
	if mrr.WithError {
		return []record.Record{}, errors.New("something went wrong")
	}

	return []record.Record{}, nil
}
