package memory

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/record"
	"github.com/rodrwan/collection/domain/song"
)

func TestMemory_GetSong(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Create a fake record to add to repository
	rec, err := record.NewRecord("Percy", "Vinyl")
	if err != nil {
		t.Fatal(err)
	}
	id := rec.GetID()

	s, err := song.NewSong("lala", 100, id)
	if err != nil {
		t.Fatal(err)
	}

	repo := MemoryRepository{
		songs: []memorySong{
			NewFromSong(s),
		},
	}

	testCases := []testCase{
		{
			name:        "No Song By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: song.ErrSongNotFound,
		}, {
			name:        "Song By ID",
			id:          s.GetID(),
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
		length      int64
		recordID    uuid.UUID
		expectedErr error
	}
	id := uuid.New()
	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Percy",
			length:      42421,
			recordID:    id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				songs: []memorySong{
					NewFromSong(song.Song{}),
				},
			}
			s, err := song.NewSong(tc.cust, tc.length, tc.recordID)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(s)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(s.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != s.GetID() {
				t.Errorf("Expected %v, got %v", s.GetID(), found.GetID())
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	ctx := context.Background()
	tests := []struct {
		name    string
		args    args
		want    *MemoryRepository
		wantErr bool
	}{
		{
			name: "new memory service",
			args: args{
				ctx: ctx,
			},
			want: &MemoryRepository{
				songs: make([]memorySong, 0),
			},
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
		songs []memorySong
	}
	id1 := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()

	song1, _ := song.NewSongWithID(id1, "10", 10, id3)
	song2, _ := song.NewSongWithID(id2, "20", 20, id3)
	expectedSongs := []song.Song{
		song1,
		song2,
	}
	tests := []struct {
		name    string
		fields  fields
		want    []song.Song
		wantErr bool
	}{
		{
			name: "Get 2 records with FindRecords",
			fields: fields{
				songs: []memorySong{
					{
						ID:     id1,
						Name:   "10",
						Length: 10,
					},
					{
						ID:     id2,
						Name:   "20",
						Length: 20,
					},
				},
			},
			want:    expectedSongs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MemoryRepository{
				songs: tt.fields.songs,
			}
			got, err := mr.FindRecords()
			if err != nil {
				t.Errorf("Song MemoryRepository.FindRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[0].GetID(), tt.want[0].GetID()) {
				t.Errorf("Song MemoryRepository.FindRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryRepository_Update(t *testing.T) {
	type fields struct {
		songs []memorySong
	}
	type args struct {
		s *song.Song
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MemoryRepository{
				songs: tt.fields.songs,
			}
			if err := mr.Update(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("MemoryRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_FindSongsByRecord(t *testing.T) {
	type fields struct {
		songs []memorySong
	}
	type args struct {
		id uuid.UUID
	}
	id1 := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()

	song1, _ := song.NewSongWithID(id1, "10", 10, id3)
	song2, _ := song.NewSongWithID(id2, "20", 20, id3)
	expectedSongs := []song.Song{
		song1,
		song2,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []song.Song
		wantErr bool
	}{
		{
			name: "Get songs by record id",
			fields: fields{
				songs: []memorySong{
					{
						ID:       id1,
						Name:     "10",
						Length:   10,
						RecordID: id3,
					},
					{
						ID:       id2,
						Name:     "20",
						Length:   20,
						RecordID: id3,
					},
				},
			},
			args: args{
				id: id3,
			},
			want:    expectedSongs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MemoryRepository{
				songs: tt.fields.songs,
			}
			got, err := mr.FindSongsByRecord(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoryRepository.FindSongsByRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoryRepository.FindSongsByRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
