package logre

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type LogResponse struct {
	ID struct {
		Data struct {
			Box struct {
				Ref struct {
					ID string `json:"id"`
				} `json:"@ref"`
			} `json:"box"`
			Body interface{} `json:"body"`
		} `json:"data"`
	} `json:"id"`
}

func GetLogBody(r *http.Response, t *testing.T) *LogResponse {
	var m LogResponse

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		t.Error(err)
	}

	return &m
}

func log(boxId string, payload []byte) (*http.Response, error) {
	var response *http.Response
	var err error

	url := fmt.Sprintf("https://logre.io/api/boxes/%s/logs", boxId)
	response, err = http.Post(url, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	return response, nil
}

func logMessage(boxId string, message string) (*http.Response, error) {
	var response *http.Response
	var err error

	url := fmt.Sprintf("https://logre.io/api/boxes/%s/logs", boxId)

	response, err = http.Post(url, "text/plain", bytes.NewBuffer([]byte(message)))

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Util
func addSeverity(i *interface{}, severity string) ([]byte, error) {
	var m map[string]string

	body, err := json.Marshal(i)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	m["severity"] = severity

	body, err = json.Marshal(m)

	if err != nil {
		return nil, err
	}

	return body, nil
}
