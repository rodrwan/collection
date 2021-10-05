package record_test

import (
	"testing"

	"github.com/rodrwan/collection/domain/record"
)

func TestRecord_NewRecord(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		test        string
		name        string
		kind        string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			kind:        "",
			expectedErr: record.ErrMissingValues,
		}, {
			test:        "Valid Name",
			name:        "Percy Bolmer",
			kind:        "Vinyl",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			_, err := record.NewRecord(tc.name, tc.kind)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
