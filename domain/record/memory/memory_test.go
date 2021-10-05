package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/record"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Create a fake customer to add to repository
	rec, err := record.NewRecord("Percy", "Vinyl")
	if err != nil {
		t.Fatal(err)
	}
	id := rec.GetID()
	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := MemoryRepository{
		records: []memoryRecord{
			NewFromRecord(rec),
		},
	}

	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: record.ErrRecordNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		kind        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Percy",
			kind:        "Vinyl",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				records: []memoryRecord{
					NewFromRecord(record.Record{}),
				},
			}
			rec, err := record.NewRecord(tc.cust, tc.kind)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(rec)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(rec.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != rec.GetID() {
				t.Errorf("Expected %v, got %v", rec.GetID(), found.GetID())
			}
		})
	}
}
