package song_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/song"
)

func TestSong_NewSong(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		test        string
		name        string
		length      int64
		recordId    uuid.UUID
		expectedErr error
	}
	id := uuid.New()
	// Create new test cases
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			length:      0,
			recordId:    uuid.UUID{},
			expectedErr: song.ErrMissingValues,
		}, {
			test:        "Valid Name",
			name:        "Percy Bolmer",
			length:      12314,
			recordId:    id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			_, err := song.NewSong(tc.name, tc.length, tc.recordId)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}

func TestNewSongWithID(t *testing.T) {
	type args struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}

	id := uuid.New()
	name := "s1"
	length := int64(100)
	recordID := uuid.New()

	s, _ := song.NewSongWithID(id, name, length, recordID)
	tests := []struct {
		name    string
		args    args
		want    song.Song
		wantErr error
	}{
		{
			name: "Create new song with id",
			args: args{
				id:       id,
				name:     name,
				length:   length,
				recordID: recordID,
			},
			want:    s,
			wantErr: nil,
		},
		{
			name: "Create new song with id, but no name",
			args: args{
				id:       id,
				name:     "",
				length:   length,
				recordID: recordID,
			},
			want:    s,
			wantErr: song.ErrMissingValues,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := song.NewSongWithID(tt.args.id, tt.args.name, tt.args.length, tt.args.recordID)
			fmt.Println(err != nil)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("NewSongWithID() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSongWithID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_GetID(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}

	id := uuid.New()
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "Get ID",
			fields: fields{
				id:       id,
				name:     "s1",
				length:   10,
				recordID: uuid.New(),
			},
			want: id,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := song.NewSongWithID(tt.fields.id, tt.fields.name, tt.fields.length, tt.fields.recordID)
			if got := s.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_GetName(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get ID",
			fields: fields{
				id:       uuid.New(),
				name:     "s1",
				length:   10,
				recordID: uuid.New(),
			},
			want: "s1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := song.NewSongWithID(tt.fields.id, tt.fields.name, tt.fields.length, tt.fields.recordID)
			if got := s.GetName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_GetLength(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Get ID",
			fields: fields{
				id:       uuid.New(),
				name:     "s1",
				length:   10,
				recordID: uuid.New(),
			},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := song.NewSongWithID(tt.fields.id, tt.fields.name, tt.fields.length, tt.fields.recordID)
			if got := s.GetLength(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_GetRecordID(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}
	id := uuid.New()
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "Get ID",
			fields: fields{
				id:       uuid.New(),
				name:     "s1",
				length:   10,
				recordID: id,
			},
			want: id,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := song.NewSongWithID(tt.fields.id, tt.fields.name, tt.fields.length, tt.fields.recordID)
			if got := s.GetRecordID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_SetID(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set id",
			args: args{
				id: uuid.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &song.Song{}
			s.SetID(tt.args.id)
		})
	}
}

func TestSong_SetName(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set id",
			args: args{
				name: "s1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &song.Song{}
			s.SetName(tt.args.name)
		})
	}
}
func TestSong_SetLength(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}
	type args struct {
		length int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set id",
			args: args{
				length: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &song.Song{}
			s.SetLength(tt.args.length)
		})
	}
}
func TestSong_SetRecordID(t *testing.T) {
	type fields struct {
		id       uuid.UUID
		name     string
		length   int64
		recordID uuid.UUID
	}
	type args struct {
		recordID uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set id",
			args: args{
				recordID: uuid.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &song.Song{}
			s.SetRecordID(tt.args.recordID)
		})
	}
}
