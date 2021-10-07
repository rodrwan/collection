package record_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rodrwan/collection/domain/record"
	"github.com/rodrwan/collection/domain/song"
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

func TestNewSongWithID(t *testing.T) {
	type args struct {
		id   uuid.UUID
		name string
		kind string
	}

	id := uuid.New()
	name := "s1"
	kind := "vinyl"

	r, _ := record.NewRecordWithID(id, name, kind)
	tests := []struct {
		name    string
		args    args
		want    record.Record
		wantErr error
	}{
		{
			name: "Create new record with id",
			args: args{
				id:   id,
				name: name,
				kind: kind,
			},
			want:    r,
			wantErr: nil,
		},
		{
			name: "Create new record with id, but no name",
			args: args{
				id:   id,
				name: "",
				kind: kind,
			},
			want:    r,
			wantErr: record.ErrMissingValues,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := record.NewRecordWithID(tt.args.id, tt.args.name, tt.args.kind)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("NewRecordWithID() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecordWithID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecord_GetID(t *testing.T) {
	type fields struct {
		id   uuid.UUID
		name string
		kind string
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
				id:   id,
				name: "s1",
				kind: "vinyl",
			},
			want: id,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := record.NewRecordWithID(tt.fields.id, tt.fields.name, tt.fields.kind)
			if got := s.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecord_GetName(t *testing.T) {
	type fields struct {
		id   uuid.UUID
		name string
		kind string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get ID",
			fields: fields{
				id:   uuid.New(),
				name: "s1",
				kind: "vinyl",
			},
			want: "s1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := record.NewRecordWithID(tt.fields.id, tt.fields.name, tt.fields.kind)
			if got := s.GetName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecord_GetKind(t *testing.T) {
	type fields struct {
		id   uuid.UUID
		name string
		kind string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get ID",
			fields: fields{
				id:   uuid.New(),
				name: "s1",
				kind: "vinyl",
			},
			want: "vinyl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := record.NewRecordWithID(tt.fields.id, tt.fields.name, tt.fields.kind)
			if got := s.GetKind(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecord_GetSongs(t *testing.T) {
	type fields struct {
		id    uuid.UUID
		name  string
		kind  string
		songs []*song.Song
	}

	r1 := uuid.New()
	s1, _ := song.NewSongWithID(uuid.New(), "la1", 10, r1)
	s2, _ := song.NewSongWithID(uuid.New(), "la2", 20, r1)

	tests := []struct {
		name   string
		fields fields
		want   []*song.Song
	}{
		{
			name: "Get ID",
			fields: fields{
				id:   uuid.New(),
				name: "s1",
				kind: "vinyl",
				songs: []*song.Song{
					&s1, &s2,
				},
			},
			want: []*song.Song{
				&s1, &s2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := record.NewRecordWithID(tt.fields.id, tt.fields.name, tt.fields.kind)
			r.AddSong(&s1)
			r.AddSong(&s2)
			if got := r.GetSongs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecord_SetID(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		args args
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
			s := &record.Record{}
			s.SetID(tt.args.id)
		})
	}
}

func TestRecord_SetName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
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
			s := &record.Record{}
			s.SetName(tt.args.name)
		})
	}
}
func TestRecord_SetKind(t *testing.T) {
	type args struct {
		kind string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Set id",
			args: args{
				kind: "vinyl",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &record.Record{}
			s.SetKind(tt.args.kind)
		})
	}
}

func TestRecord_ToPublic(t *testing.T) {
	type fields struct {
		id    uuid.UUID
		name  string
		kind  string
		songs []*song.Song
	}
	id := uuid.New()
	tests := []struct {
		name   string
		fields fields
		want   record.PublicRecord
	}{
		{
			name: "Transform record to public record",
			fields: fields{
				id:   id,
				name: "r1",
				kind: "vinyl",
			},
			want: record.PublicRecord{
				ID:   id,
				Name: "r1",
				Kind: "vinyl",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &record.Record{}
			r.SetID(tt.fields.id)
			r.SetName(tt.fields.name)
			r.SetKind(tt.fields.kind)

			if got := r.ToPublic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Record.ToPublic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecord_ToRecord(t *testing.T) {
	type fields struct {
		id    uuid.UUID
		name  string
		kind  string
		songs []*song.Song
	}
	id := uuid.New()
	rr := &record.Record{}
	rr.SetID(id)
	rr.SetName("r1")
	rr.SetKind("vinyl")

	tests := []struct {
		name   string
		fields fields
		want   *record.Record
	}{
		{
			name: "Transform public record to record",
			fields: fields{
				id:   id,
				name: "r1",
				kind: "vinyl",
			},
			want: rr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &record.PublicRecord{
				ID:   tt.fields.id,
				Name: tt.fields.name,
				Kind: tt.fields.kind,
			}

			if got := r.ToRecord(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Record.ToPublic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPublicArray(t *testing.T) {
	type args struct {
		records []record.Record
	}

	r1, _ := record.NewRecord("r1", "vinyl")
	r2, _ := record.NewRecord("r2", "vinyl")
	rr := []record.Record{
		r1, r2,
	}
	rrp := []record.PublicRecord{
		r1.ToPublic(), r2.ToPublic(),
	}

	tests := []struct {
		name string
		args args
		want []record.PublicRecord
	}{
		{
			name: "transform record slice to public record slice",
			args: args{
				records: rr,
			},
			want: rrp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := record.ToPublicArray(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPublicArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
