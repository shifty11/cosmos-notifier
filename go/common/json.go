package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

func GetJson(httpClient *http.Client, url string, target interface{}) (int, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return resp.StatusCode, errors.New(resp.Status)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	return resp.StatusCode, json.NewDecoder(resp.Body).Decode(target)
}
