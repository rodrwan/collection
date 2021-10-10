package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rodrwan/collection/config"
	"github.com/rodrwan/collection/pkg/server"
	"github.com/rodrwan/collection/services"
	"github.com/stretchr/testify/assert"
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
			assert.Equalf(t, tc.expectedErr, err, fmt.Sprintf("Expected error %v, got %v", tc.expectedErr, err))
		})
	}
}

func TestServer_CreateRecord(t *testing.T) {
	tests := []struct {
		description   string // description of the test case
		route         string // route path to test
		data          []byte // request data
		method        string // request method
		expectedCode  int    // expected HTTP status code
		expectedOk    bool   // expected ok messaje
		expectedError string // expected error message messaje
	}{
		{
			description:   "get HTTP status 201 on record creation",
			route:         "/CreateRecord",
			data:          []byte(`{ "name": "lala", "kind": "vinyl"}`),
			method:        fiber.MethodPost,
			expectedCode:  201,
			expectedOk:    true,
			expectedError: "",
		},
		{
			description:   "get HTTP status 400 when create a record with empty body",
			route:         "/CreateRecord",
			data:          []byte(`{}`),
			method:        fiber.MethodPost,
			expectedCode:  400,
			expectedOk:    false,
			expectedError: "missing value",
		},
		{
			description:   "get HTTP status 400 when create a record with missing record type",
			route:         "/CreateRecord",
			data:          []byte(`{ "name": "lala" }`),
			method:        fiber.MethodPost,
			expectedCode:  400,
			expectedOk:    false,
			expectedError: "Invalid record type",
		},
		{
			description:   "get HTTP status 422 when create a record with missing record type",
			route:         "/CreateRecord",
			data:          []byte(``),
			method:        fiber.MethodPost,
			expectedCode:  422,
			expectedOk:    false,
			expectedError: "json: unexpected end of JSON input: ",
		},
	}

	app := fiber.New()
	collectionService, err := services.NewCollectionService(
		services.WithRecordMemoryRepository(),
		services.WithSongMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}
	srv, err := server.NewServer(collectionService)
	if err != nil {
		log.Fatal(err)
	}

	// Create route with GET method for test
	app.Post("/CreateRecord", srv.CreateRecord)

	type response struct {
		Ok     bool        `json:"ok,omitempty"`
		Record interface{} `json:"record,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	// Iterate through test single test cases
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.route, bytes.NewBuffer(test.data))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, 1000)
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var r response
			if err := json.Unmarshal(body, &r); err != nil {
				log.Fatal(err)
			}
			// Verify, if the status code is as expected
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			assert.Equalf(t, test.expectedOk, r.Ok, test.description)
			assert.Equalf(t, test.expectedError, r.Error, test.description)
		})
	}
}

func TestServer_GetRecords(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		method       string
		expectedCode int // expected HTTP status code
		services     []services.CollectionConfiguration
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/GetRecords",
			method:       fiber.MethodGet,
			expectedCode: 200,
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
		},
		// First test case
		{
			description:  "get HTTP status 500",
			route:        "/GetRecords",
			method:       fiber.MethodGet,
			expectedCode: 500,
			services: []services.CollectionConfiguration{
				services.WithFakeRecordService(true, uuid.Nil),
				services.WithSongMemoryRepository(),
			},
		},
	}

	for _, test := range tests {
		app := fiber.New()
		collectionService, err := services.NewCollectionService(
			test.services...,
		)
		if err != nil {
			log.Fatal(err)
		}
		srv, err := server.NewServer(collectionService)
		if err != nil {
			log.Fatal(err)
		}

		app.Get("/GetRecords", srv.GetRecords)

		req := httptest.NewRequest(test.method, test.route, nil)
		resp, err := app.Test(req, 1)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestServer_GetRecordsWithRecords(t *testing.T) {
	tests := []struct {
		description     string // description of the test case
		route           string // route path to test
		method          string // request method
		expectedCode    int    // expected HTTP status code
		expectedOk      bool   // expected ok message
		expectedError   string // expected error message message
		expectedRecords int
	}{
		// First test case
		{
			description:     "get HTTP status 200",
			route:           "/GetRecords",
			method:          fiber.MethodGet,
			expectedCode:    200,
			expectedOk:      true,
			expectedRecords: 2,
		},
	}

	app := fiber.New()
	collectionService, err := services.NewCollectionService(
		services.WithRecordMemoryRepository(),
		services.WithSongMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}

	id1 := uuid.New()
	collectionService.AddRecord(id1, "test", "vinyl")
	id2 := uuid.New()
	collectionService.AddRecord(id2, "test1", "vinyl")

	srv, err := server.NewServer(collectionService)
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/GetRecords", srv.GetRecords)

	type response struct {
		Ok      bool `json:"ok,omitempty"`
		Records []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Kind string `json:"kind"`
		} `json:"records,omitempty"`
		Error string `json:"error,omitempty"`
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, nil)
		resp, err := app.Test(req, 1)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)

		}

		var r response
		if err := json.Unmarshal(body, &r); err != nil {
			log.Fatal(err)
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		assert.Equalf(t, test.expectedRecords, len(r.Records), test.description)
	}
}

func TestServer_GetRecordWithEmptyStore(t *testing.T) {
	id := uuid.New()
	tests := []struct {
		description   string // description of the test case
		route         string // route path to test
		method        string // request method
		expectedCode  int    // expected HTTP status code
		expectedOk    bool   // expected ok message
		expectedError string // expected error message message
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        fmt.Sprintf("/GetRecordById/%s", id),
			method:       fiber.MethodGet,
			expectedCode: 404,
			expectedOk:   true,
		},
	}

	app := fiber.New()
	collectionService, err := services.NewCollectionService(
		services.WithRecordMemoryRepository(),
		services.WithSongMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}
	srv, err := server.NewServer(collectionService)
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/GetRecordById/:id", srv.GetRecordById)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, nil)
		resp, err := app.Test(req, 1)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestServer_GetRecordWithRecords(t *testing.T) {
	tests := []struct {
		description   string // description of the test case
		route         string // route path to test
		method        string // request method
		expectedCode  int    // expected HTTP status code
		expectedOk    bool   // expected ok messaje
		expectedError string // expected error message messaje
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/GetRecordById",
			method:       fiber.MethodGet,
			expectedCode: 200,
			expectedOk:   true,
		},
	}

	app := fiber.New()
	collectionService, err := services.NewCollectionService(
		services.WithRecordMemoryRepository(),
		services.WithSongMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}
	srv, err := server.NewServer(collectionService)
	if err != nil {
		log.Fatal(err)
	}
	id1 := uuid.New()
	rec, _ := collectionService.AddRecord(id1, "test", "vinyl")

	app.Get("/GetRecordById/:id", srv.GetRecordById)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, fmt.Sprintf("%s/%s", test.route, rec.ID.String()), nil)
		resp, err := app.Test(req, 1)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestServer_AddSongToExistingRecord(t *testing.T) {
	type args struct {
		id   uuid.UUID
		name string
		kind string
	}
	id := uuid.New()
	tests := []struct {
		description   string // description of the test case
		route         string // route path to test
		method        string // request method
		data          []byte
		expectedCode  int    // expected HTTP status code
		expectedOk    bool   // expected ok messaje
		expectedError string // expected error message messaje
		services      []services.CollectionConfiguration
		args          []args
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        fmt.Sprintf("/AddSongToRecordById/%s", id.String()),
			method:       fiber.MethodPost,
			expectedCode: 200,
			expectedOk:   true,
			data:         []byte(`{ "name": "lala", "length": 100 }`),
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
			args: []args{
				{
					id:   id,
					name: "lala",
					kind: "vinyl",
				},
			},
		},
		{
			description:  "get HTTP status 400",
			route:        fmt.Sprintf("/AddSongToRecordById/%s", uuid.New().String()),
			method:       fiber.MethodPost,
			expectedCode: 400,
			expectedOk:   false,
			data:         []byte(`{ "name": "lala", "length": 100 }`),
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
			args: []args{
				{
					id:   id,
					name: "lala",
					kind: "vinyl",
				},
			},
		},
		{
			description:   "get HTTP status 422",
			route:         fmt.Sprintf("/AddSongToRecordById/%s", id.String()),
			method:        fiber.MethodPost,
			expectedCode:  422,
			expectedOk:    false,
			data:          []byte(``),
			expectedError: "json: unexpected end of JSON input: ",
			services: []services.CollectionConfiguration{
				services.WithRecordMemoryRepository(),
				services.WithSongMemoryRepository(),
			},
			args: []args{
				{
					id:   id,
					name: "lala",
					kind: "vinyl",
				},
			},
		},
		{
			description:   "get HTTP status 400",
			route:         fmt.Sprintf("/AddSongToRecordById/%s", id.String()),
			method:        fiber.MethodPost,
			expectedCode:  400,
			expectedOk:    false,
			data:          []byte(``),
			expectedError: "json: unexpected end of JSON input: ",
			services: []services.CollectionConfiguration{
				services.WithFakeRecordService(true, id),
				services.WithSongMemoryRepository(),
			},
			args: []args{
				{
					id:   id,
					name: "lala",
					kind: "vinyl",
				},
			},
		},
	}

	type response struct {
		Ok     bool `json:"ok,omitempty"`
		Record struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Kind string `json:"kind"`
		} `json:"record,omitempty"`
		Error string `json:"error,omitempty"`
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			app := fiber.New(config.NewFiberConfig)
			collectionService, err := services.NewCollectionService(
				test.services...,
			)
			if err != nil {
				log.Fatal(err)
			}
			srv, err := server.NewServer(collectionService)
			if err != nil {
				log.Fatal(err)
			}

			for _, arg := range test.args {
				collectionService.AddRecord(arg.id, arg.name, arg.kind)
			}

			app.Post("/AddSongToRecordById/:id", srv.AddSongToRecordById)

			req := httptest.NewRequest(test.method, test.route, bytes.NewBuffer(test.data))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, 1000)
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)

			}

			var r response
			if err := json.Unmarshal(body, &r); err != nil {
				log.Fatal(err)
			}

			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			assert.Equalf(t, test.expectedOk, r.Ok, test.description)
			app.Shutdown()
		})
	}
}
