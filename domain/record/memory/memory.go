package memory

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/record"
)

type MemoryRepository struct {
	records []memoryRecord

	sync.Mutex
}

type memoryRecord struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
	Kind string    `db:"kind"`
}

// NewFromCustomer takes in a aggregate and converts into internal structure
func NewFromRecord(r record.Record) memoryRecord {
	return memoryRecord{
		ID:   r.GetID(),
		Name: r.GetName(),
		Kind: r.GetKind(),
	}
}

func (pr memoryRecord) ToRecord() record.Record {
	r := record.Record{}

	r.SetID(pr.ID)
	r.SetName(pr.Name)
	r.SetKind(pr.Kind)

	return r
}

// Create a new mongodb repository
func New(ctx context.Context) (*MemoryRepository, error) {
	return &MemoryRepository{
		records: make([]memoryRecord, 0),
	}, nil
}

func (mr *MemoryRepository) Get(id uuid.UUID) (record.Record, error) {
	mr.Lock()
	defer mr.Unlock()

	for _, rec := range mr.records {
		if rec.ID == id {
			return rec.ToRecord(), nil
		}
	}

	return record.Record{}, nil
}

func (mr *MemoryRepository) Add(r record.Record) error {
	mr.Lock()
	defer mr.Unlock()
	log.Printf("Add record: %s\n", r.GetID().String())
	internal := NewFromRecord(r)
	mr.records = append(mr.records, internal)

	return nil
}

func (mr *MemoryRepository) FindRecords() ([]record.Record, error) {
	mr.Lock()
	defer mr.Unlock()

	// Convert to aggregate
	var rr []record.Record
	for _, r := range mr.records {
		rr = append(rr, r.ToRecord())
	}

	return rr, nil
}

func (mr *MemoryRepository) Update(r *record.Record) error {
	panic("to implement")
}