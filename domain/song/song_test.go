package song_test

import (
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
