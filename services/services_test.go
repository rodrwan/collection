package services_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/record"
	"github.com/rodrwan/collection/services"
	"github.com/stretchr/testify/assert"
)

func TestCollectionService_AddRecord(t *testing.T) {
	type args struct {
		name string
		kind string
	}

	tests := []struct {
		name        string
		description string
		args        args
		want        record.PublicRecord
		expectedErr error
	}{
		{
			name:        "Vinyl",
			description: "",
			args:        args{name: "lala", kind: "vinyl"},
			want: record.PublicRecord{
				Name: "lala",
				Kind: "vinyl",
			},
			expectedErr: nil,
		},
		{
			name:        "mp3",
			description: "",
			args:        args{name: "lala", kind: "mp3"},
			want: record.PublicRecord{
				Name: "lala",
				Kind: "mp3",
			},
			expectedErr: nil,
		},
		{
			name:        "aiff",
			description: "",
			args:        args{name: "lala", kind: "aiff"},
			want: record.PublicRecord{
				Name: "lala",
				Kind: "aiff",
			},
			expectedErr: services.ErrInvalidType,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			)

			got, err := cs.AddRecord(test.args.name, test.args.kind)
			if err != nil {
				assert.Equalf(t, test.expectedErr.Error(), err.Error(), test.description)
				return
			}

			assert.Equalf(t, test.want.Kind, got.Kind, test.description)
		})
	}
}

func TestCollectionService_FindRecord(t *testing.T) {
	tests := []struct {
		name        string
		description string
		args        string
		want        record.PublicRecord
		expectedErr error
	}{
		{
			name:        "Vinyl",
			description: "",
			args:        uuid.NewString(),
			want: record.PublicRecord{
				Name: "lala",
				Kind: "vinyl",
			},
			expectedErr: errors.New("record not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			)

			got, err := cs.FindRecord(test.args)
			if err != nil {
				assert.Equalf(t, test.expectedErr.Error(), err.Error(), test.description)
				return
			}

			assert.Equalf(t, test.want.Kind, got.Kind, test.description)
		})
	}
}

func TestCollectionService_FindAllRecord(t *testing.T) {
	tests := []struct {
		name        string
		description string
		want        int
		expectedErr error
	}{
		{
			name:        "Vinyl",
			description: "",
			want:        0,
			expectedErr: errors.New("record not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			)

			got, err := cs.FindAllRecord()
			if err != nil {
				assert.Equalf(t, test.expectedErr.Error(), err.Error(), test.description)
				return
			}

			assert.Equalf(t, test.want, len(got), test.description)
		})
	}

}
func TestCollectionService_AddSongToRecord(t *testing.T) {
	r, _ := record.NewRecord("lala", "vinyl")

	type args struct {
		rec    *record.Record
		name   string
		length int64
	}

	tests := []struct {
		name        string
		description string
		args        args
		want        int
		expectedErr error
	}{
		{
			name:        "Vinyl",
			description: "",
			args: args{
				rec:    &r,
				name:   "lalo",
				length: 100,
			},
			want:        0,
			expectedErr: errors.New("record not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			)

			if err := cs.AddSongToRecord(test.args.rec, test.args.name, test.args.length); err != nil {
				assert.Equalf(t, test.expectedErr.Error(), err.Error(), test.description)
				return
			}

			// assert.Equalf(t, test.want, got, test.description)
		})
	}
}
