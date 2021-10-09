package memory

import (
	"context"
	"reflect"
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

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	mr := &MemoryRepository{
		records: make([]memoryRecord, 0),
	}
	tests := []struct {
		name    string
		args    args
		want    *MemoryRepository
		wantErr bool
	}{
		{
			name: "create new memory service",
			args: args{
				ctx: context.Background(),
			},
			want:    mr,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryRepository_FindRecords(t *testing.T) {
	type fields struct {
		records []memoryRecord
	}

	r1, _ := record.NewRecord("r1", "vinyl")
	r2, _ := record.NewRecord("r2", "vinyl")
	rr := []record.Record{
		r1, r2,
	}
	tests := []struct {
		name    string
		fields  fields
		want    []record.Record
		wantErr bool
	}{
		{
			name: "Find records",
			fields: fields{
				records: []memoryRecord{
					{
						ID:   r1.GetID(),
						Name: r1.GetName(),
						Kind: r1.GetKind(),
					},
					{
						ID:   r2.GetID(),
						Name: r2.GetName(),
						Kind: r2.GetKind(),
					},
				},
			},
			want:    rr,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MemoryRepository{
				records: tt.fields.records,
			}
			got, err := mr.FindRecords()
			if err != nil {
				t.Errorf("Record err MemoryRepository.FindRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[0].GetID(), tt.want[0].GetID()) {
				t.Errorf("Record no equal MemoryRepository.FindRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
