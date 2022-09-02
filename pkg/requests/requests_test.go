package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"chuck-jokes/models"
)

// MockClient is the mock client
//type MockClient struct {
//	DoFunc func(req *http.Request) (*http.Response, error)
//}

// GetDoFunc fetches the mock client's `Do` func
// var GetDoFunc func(req *http.Request) (*http.Response, error)

type MockClient struct {
	response *http.Response
	err      error
}

// Do is the mock client's `Do` func
func (m *MockClient) Do(_ *http.Request) (*http.Response, error) {
	return m.response, m.err
}

var joke = struct {
	ID    string
	Value string
}{
	ID:    "testID",
	Value: "Test value",
}

var modelJoke = models.Joke{
	ExternalID: "testID",
	Value:      "Test value",
}

func TestExternalRequest_CallRandom(t *testing.T) {
	j, err := json.Marshal(joke)
	if err != nil {
		t.Skip("Something went wrong with this test")
	}

	tests := []struct {
		name   string
		random string
		client MockClient
		want   *models.Joke
		panic  bool
	}{
		{
			name:   "Base test",
			random: "test",
			client: MockClient{
				response: &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(j)),
				},
				err: nil,
			},
			want:  &modelJoke,
			panic: false,
		},
		{
			name:   "Panic test",
			random: "test",
			client: MockClient{
				response: &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				},
				err: nil,
			},
			want:  nil,
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewExternalRequest(tt.random, &tt.client)
			defer func() {
				if r := recover(); tt.panic && r == nil {
					t.Errorf("The code did not panic")
				}
			}()
			got := e.CallRandom()

			if got.ExternalID != tt.want.ExternalID {
				t.Errorf("CallRandom() = %v, want %v", got.ExternalID, tt.want.ExternalID)
			}
		})
	}
}

func TestJoke_ChangeToBaseModel(t *testing.T) {
	type fields struct {
		Value string
		ID    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.Joke
	}{
		{
			name: "base test",
			fields: fields{
				Value: joke.Value,
				ID:    joke.ID,
			},
			want: &modelJoke,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Joke{
				Value: tt.fields.Value,
				ID:    tt.fields.ID,
			}
			if got := j.ChangeToBaseModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeToBaseModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
