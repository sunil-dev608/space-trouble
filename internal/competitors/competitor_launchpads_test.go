package competitors

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_competitorLaunchpadsProvier_FetchLaunchpads(t *testing.T) {
	type fields struct {
		APIURL       string
		Client       *http.Client
		lauchpadsMap map[string]string
	}
	tests := []struct {
		name     string
		fields   fields
		response []map[string]string
		wantErr  bool
	}{
		{
			name: "success",
			fields: fields{
				APIURL:       "",
				Client:       &http.Client{},
				lauchpadsMap: map[string]string{"ksc_lc_39a": "active"},
			},
			response: []map[string]string{{"id": "ksc_lc_39a", "status": "active"}},
			wantErr:  false,
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
			p := &competitorLaunchpadsProvier{
				APIURL:       ts.URL,
				Client:       tt.fields.Client,
				lauchpadsMap: tt.fields.lauchpadsMap,
			}
			got, err := p.FetchLaunchpads()
			if (err != nil) != tt.wantErr {
				t.Errorf("competitorLaunchpadsProvier.FetchLaunchpads() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.fields.lauchpadsMap) {
				t.Errorf("competitorLaunchpadsProvier.FetchLaunchpads() = %v, response %v", got, tt.response)
			}
		})
	}
}
