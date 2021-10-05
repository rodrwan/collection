package memory

import (
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
