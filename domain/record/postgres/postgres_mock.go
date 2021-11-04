package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type MockDB struct {
	callParams []interface{}
}

// Create a new mongodb repository
func NewMockDB(mock *MockDB) *PostgresRepository {
	return &PostgresRepository{
		db: mock,
	}
}

func (mdb *MockDB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	fmt.Println(">>>> GetContext <<<<")
	fmt.Println(query)
	fmt.Println(args...)
	mdb.callParams = []interface{}{query}
	mdb.callParams = append(mdb.callParams, args...)
	return nil
}

func (mdb *MockDB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	fmt.Println(">>>> NamedExecContext <<<<")
	fmt.Println(query)
	fmt.Println(arg)
	mdb.callParams = []interface{}{query}
	mdb.callParams = append(mdb.callParams, arg)
	return nil, nil
}

func (mdb *MockDB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	fmt.Println(">>>> SelectContext <<<<")
	fmt.Println(query)
	fmt.Println(args...)
	mdb.callParams = []interface{}{query}
	mdb.callParams = append(mdb.callParams, args...)
	return nil
}

// Add a helper method to inspect the `callParams` field
func (mdb *MockDB) CalledWith() []interface{} {
	return mdb.callParams
}
