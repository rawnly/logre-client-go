package logre_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Log(boxId string, message string, severity string) (*http.Response, error) {
	var response *http.Response
	var err error
	var body []byte

	url := fmt.Sprintf("https://logre.io/api/boxes/%s/logs", boxId)
	payload := map[string]string {
		"message": message,
		"severity": severity,
	}

	body, err = json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	response, err = http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	return response, nil
}

