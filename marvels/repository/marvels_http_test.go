package repository_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	repo "github.com/ivantedja/xmarvel/marvels/repository"
)

type DummyHTTP struct {
	Resp *http.Response
	Err  error
}

func (dh *DummyHTTP) RoundTrip(req *http.Request) (*http.Response, error) {
	return dh.Resp, dh.Err
}

func newDummyHTTP(statusCode int, err error, respBody string) *DummyHTTP {
	return &DummyHTTP{
		&http.Response{
			StatusCode: statusCode,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(respBody))),
		},
		err,
	}
}

func TestMarvelHttp_GetCollection(t *testing.T) {
	dummyResponse := `
		{
			"code": 200,
			"status": "Ok",
			"data": {
				"offset": 0,
				"limit": 20,
				"total": 1,
				"count": 1,
				"results": [
					{
						"id": 1011334,
						"name": "3-D Man"
					}
				]
			}
		}
	`
	type args struct {
		baseUrl string
		filter  map[string]string
	}
	tests := map[string]struct {
		args        args
		dummyServer *DummyHTTP
		wantErr     bool
	}{
		"response success":   {args{"http://example.com", map[string]string{}}, newDummyHTTP(200, nil, dummyResponse), false},
		"with filter":        {args{"http://example.com", map[string]string{"limit": "100"}}, newDummyHTTP(200, nil, dummyResponse), false},
		"not found error":    {args{"http://example.com", map[string]string{}}, newDummyHTTP(404, errors.New("not found"), dummyResponse), true},
		"not found response": {args{"http://example.com", map[string]string{}}, newDummyHTTP(404, nil, dummyResponse), true},
		"fail marshal":       {args{"http://example.com", map[string]string{}}, newDummyHTTP(200, nil, "zzz"), true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// fake the request, then return to normal later on
			tmp := repo.HttpClient
			defer func() {
				repo.HttpClient = tmp
			}()
			repo.HttpClient = &http.Client{
				Transport: tt.dummyServer,
			}
			a := repo.NewAPI(
				tt.args.baseUrl,
				"",
				"",
				time.Second,
			)

			_, err := a.Search(context.TODO(), tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("api.GetCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMarvelHttp_Show(t *testing.T) {
	dummyResponse := `
		{
			"code": 200,
			"status": "Ok",
			"data": {
				"offset": 0,
				"limit": 20,
				"total": 1,
				"count": 1,
				"results": [
					{
						"id": 1011334,
						"name": "3-D Man"
					}
				]
			}
		}
	`
	type args struct {
		baseUrl string
		ID      int
	}
	tests := map[string]struct {
		args        args
		dummyServer *DummyHTTP
		wantErr     bool
	}{
		"response success":   {args{"http://example.com", 1016823}, newDummyHTTP(200, nil, dummyResponse), false},
		"not found response": {args{"http://example.com", 99999999999}, newDummyHTTP(404, nil, dummyResponse), true},
		"other error":        {args{"http://example.com", 99999999999}, newDummyHTTP(403, errors.New("forbidden"), dummyResponse), true},
		"fail marshal":       {args{"http://example.com", 1016823}, newDummyHTTP(200, nil, "zzz"), true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// fake the request, then return to normal later on
			tmp := repo.HttpClient
			defer func() {
				repo.HttpClient = tmp
			}()
			repo.HttpClient = &http.Client{
				Transport: tt.dummyServer,
			}
			a := repo.NewAPI(
				tt.args.baseUrl,
				"",
				"",
				time.Second,
			)

			_, err := a.Show(context.TODO(), tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("api.Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
