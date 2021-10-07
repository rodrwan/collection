package services_test

import (
	"errors"
	"reflect"
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
		services    []services.CollectionConfiguration
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
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
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
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
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
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
		},
		{
			name:        "missing name",
			description: "",
			args:        args{kind: "aiff"},
			want: record.PublicRecord{
				Name: "lala",
				Kind: "aiff",
			},
			expectedErr: record.ErrMissingValues,
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
		},
		{
			name:        "Error in record add method",
			description: "",
			args:        args{name: "lala", kind: "vinyl"},
			want: record.PublicRecord{
				Name: "lala",
				Kind: "vinyl",
			},
			expectedErr: errors.New("something went wrong"),
			services: []services.CollectionConfiguration{
				services.WithFakeRecordService(true, uuid.Nil),
				services.WithSongMemoryRepository(),
			},
		},
		{
			name:        "Error in record add method",
			description: "",
			args:        args{name: "lala", kind: "mp3"},
			want: record.PublicRecord{
				Name: "lala",
				Kind: "mp3",
			},
			expectedErr: errors.New("something went wrong"),
			services: []services.CollectionConfiguration{
				services.WithFakeRecordService(true, uuid.Nil),
				services.WithSongMemoryRepository(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				test.services...,
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
	id := uuid.New()
	tests := []struct {
		name        string
		description string
		args        string
		want        record.PublicRecord
		expectedErr error
		services    []services.CollectionConfiguration
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
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
		},
		{
			name:        "Vinyl",
			description: "",
			args:        id.String(),
			want: record.PublicRecord{
				Name: "lala",
				Kind: "vinyl",
			},
			expectedErr: nil,
			services: []services.CollectionConfiguration{
				services.WithFakeRecordService(false, id),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				test.services...,
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
		services    []services.CollectionConfiguration
	}{
		{
			name:        "Vinyl",
			description: "",
			want:        0,
			expectedErr: errors.New("record not found"),
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
		},
		{
			name:        "Vinyl",
			description: "",
			want:        0,
			expectedErr: errors.New("something went wrong"),
			services: []services.CollectionConfiguration{
				services.WithFakeRecordService(true, uuid.Nil),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				test.services...,
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
		services    []services.CollectionConfiguration
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
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs, _ := services.NewCollectionService(
				test.services...,
			)

			if err := cs.AddSongToRecord(test.args.rec, test.args.name, test.args.length); err != nil {
				assert.Equalf(t, test.expectedErr.Error(), err.Error(), test.description)
				return
			}
		})
	}
}

func TestNewCollectionService(t *testing.T) {
	type args struct {
		cfgs []services.CollectionConfiguration
	}
	tests := []struct {
		name    string
		args    args
		want    *services.CollectionService
		wantErr bool
	}{
		{
			name: "bla",
			args: args{
				cfgs: []services.CollectionConfiguration{
					services.WithError(),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := services.NewCollectionService(tt.args.cfgs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCollectionService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollectionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
