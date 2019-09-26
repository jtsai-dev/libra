package httpUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var defaultHttpClient *http.Client

func instance() {
	client := *http.DefaultClient
	client.Timeout = time.Second * 5
	defaultHttpClient = &client
}

func FetchingJson(url string, method string, v interface{}) error {
	if defaultHttpClient == nil {
		instance()
	}

	request, _ := http.NewRequest(method, url, nil)

	resp, err := defaultHttpClient.Do(request)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("fail get data with httpCode: %s", resp.StatusCode)
		return errors.New(msg)
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
