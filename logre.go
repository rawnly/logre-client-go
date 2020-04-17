package logre_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var boxId string

func Init(box string) {
	boxId = box
}

func Debug(i *interface{}) (*http.Response, error) {
	data, err := addSeverity(i, "debug")

	if err != nil {
		return nil, err
	}

	return log(boxId, data)
}

func Info(i *interface{}) (*http.Response, error) {
	data, err := addSeverity(i, "debug")

	if err != nil {
		return nil, err
	}

	return log(boxId, data)
}

func Warning(i *interface{}) (*http.Response, error) {
	data, err := addSeverity(i, "debug")

	if err != nil {
		return nil, err
	}

	return log(boxId, data)
}

func Error(i *interface{}) (*http.Response, error) {
	data, err := addSeverity(i, "debug")

	if err != nil {
		return nil, err
	}

	return log(boxId, data)
}

func Fatal(i *interface{}) (*http.Response, error) {
	data, err := addSeverity(i, "debug")

	if err != nil {
		return nil, err
	}

	return log(boxId, data)
}

func Log(i *interface{}) (*http.Response, error) {
	data, err := json.Marshal(i)

	if err != nil {
		return nil, err
	}

	return log(boxId, data)
}

func Message(message string) (*http.Response, error) {
	return logMessage(boxId, message)
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

	m["severity"] = severity;

	body, err = json.Marshal(m)

	if err != nil {
		return nil, err
	}

	return body, nil
}

