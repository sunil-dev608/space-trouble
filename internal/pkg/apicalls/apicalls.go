package apicalls

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func APICall(ctx context.Context, method string, client *http.Client, query []byte, apiURL string, response interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, apiURL, bytes.NewBuffer(query))
	if err != nil {
		return err
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}

	return nil
}
