package server_test

import (
	"log"
	"testing"

	"github.com/rodrwan/collection/pkg/server"
	"github.com/rodrwan/collection/services"
)

func TestServer_NewServer(t *testing.T) {
	type testCase struct {
		test              string
		collectionService *services.CollectionService
		expectedErr       error
	}

	collectionService, err := services.NewCollectionService(
		services.WithRecordMemoryRepository(),
		services.WithSongMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}
	testCases := []testCase{
		{
			test:              "Wrong initialization",
			collectionService: nil,
			expectedErr:       server.ErrServiceCannotBeNil,
		}, {
			test:              "Correct initialization",
			collectionService: collectionService,
			expectedErr:       nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			_, err := server.NewServer(tc.collectionService)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}
