package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rodrwan/collection/domain/record"
	"github.com/rodrwan/collection/domain/song"
)

type PostgresRepository struct {
	db *sqlx.DB
}

type (
	SqlOpener func(string, string) (*sqlx.DB, error)
)

// mongoCustomer is an internal type that is used to store a CustomerAggregate
// we make an internal struct for this to avoid coupling this mongo implementation to the customeraggregate.
// Mongo uses bson so we add tags for that
type postgresRecord struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
	Kind string    `db:"kind"`
}

// NewFromCustomer takes in a aggregate and converts into internal structure
func NewFromRecord(r record.Record) postgresRecord {
	return postgresRecord{
		ID:   r.GetID(),
		Name: r.GetName(),
		Kind: r.GetKind(),
	}
}

func (pr postgresRecord) ToRecord() record.Record {
	r := record.Record{}

	r.SetID(pr.ID)
	r.SetName(pr.Name)
	r.SetKind(pr.Kind)

	return r
}

// Create a new mongodb repository
func New(ctx context.Context, connectionString string, open SqlOpener) (*PostgresRepository, error) {
	client, err := open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Find Metabot DB
	if err := client.PingContext(ctx); err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db: client,
	}, nil
}

func (mr *PostgresRepository) Get(id uuid.UUID) (record.Record, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var r postgresRecord
	if err := mr.db.GetContext(ctx, r, "SELECT * FROM records WHERE id = ?", id); err != nil {
		return record.Record{}, err
	}

	// Convert to aggregate
	return r.ToRecord(), nil
}

func (mr *PostgresRepository) Add(r record.Record) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromRecord(r)
	_, err := mr.db.NamedExecContext(ctx, `INSERT INTO records (id, name, kind) VALUES (:id, :name, :kind)`, internal)
	if err != nil {
		return err
	}

	return nil
}

func (mr *PostgresRepository) FindRecords() ([]record.Record, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var r []*postgresRecord
	if err := mr.db.SelectContext(ctx, r, "SELECT * FROM records"); err != nil {
		return []record.Record{}, err
	}

	// Convert to aggregate
	var rr []record.Record
	for _, r := range r {
		rr = append(rr, r.ToRecord())
	}

	return rr, nil
}

func (mr *PostgresRepository) Update(r *record.Record) error {
	panic("to implement")
}

func (mr *PostgresRepository) AddSong(id uuid.UUID, s *song.Song) error {
	panic("to implement")
}
