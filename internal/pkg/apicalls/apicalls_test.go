package apicalls

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type response struct {
	Status int64 `json:"status"`
}

func TestAPICall(t *testing.T) {
	type args struct {
		ctx      context.Context
		method   string
		client   *http.Client
		query    []byte
		apiURL   string
		response interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		response interface{}
	}{
		{
			name: "success:POST",
			args: args{
				ctx:      context.Background(),
				method:   "POST",
				client:   &http.Client{},
				query:    nil,
				apiURL:   "https://api.spacexdata.com/v5/launches/query",
				response: &response{Status: 100},
			},
			wantErr:  false,
			response: &response{Status: 200},
		},
		{
			name: "success:GET",
			args: args{
				ctx:      context.Background(),
				method:   "GET",
				client:   &http.Client{},
				query:    nil,
				apiURL:   "https://api.spacexdata.com/v5/launches/query",
				response: &response{Status: 300},
			},
			wantErr:  false,
			response: &response{Status: 400},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Simulate the response you expect from the real server
				w.WriteHeader(http.StatusOK)
				bytes, _ := json.Marshal(tt.response)
				w.Write(bytes)
			}))
			defer ts.Close()

			// Replace the real http.Get with the mock server
			http.DefaultClient.Transport = http.DefaultTransport
			tt.args.apiURL = ts.URL
			if err := APICall(tt.args.ctx, tt.args.method, tt.args.client, tt.args.query, tt.args.apiURL, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("APICall() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.response.(*response).Status != tt.response.(*response).Status {
				t.Errorf("APICall() got = %v, want %v", tt.args.response, tt.response)
			}
		})
	}
}
