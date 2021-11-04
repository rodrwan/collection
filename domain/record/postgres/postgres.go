package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rodrwan/collection/domain/record"
	"github.com/rodrwan/collection/domain/song"
)

// ConnectionConfig ...
type ConnectionConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func dropConnections(db *sqlx.DB, name string) {
	query := `
		select pg_terminate_backend(pg_stat_activity.pid)
		from pg_stat_activity
		where pg_stat_activity.datname = $1 and pid <> pg_backend_pid()`
	_, err := db.Exec(query, name)
	if err != nil {
		panic(err)
	}
}

type IPostgresGetContext interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type IPostgresNamedExecContext interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type IPostgresSelectContext interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type IPostgresSQl interface {
	IPostgresGetContext
	IPostgresNamedExecContext
	IPostgresSelectContext
}

type PostgresRepository struct {
	db IPostgresSQl
}

type (
	SqlOpener func(string, string) (*sqlx.DB, error)
)

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
func New(ctx context.Context, connectionString, database string, open SqlOpener) (*PostgresRepository, error) {
	client, err := open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Find Metabot DB
	if err := client.PingContext(ctx); err != nil {
		return nil, err
	}

	dropConnections(client, database)
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
