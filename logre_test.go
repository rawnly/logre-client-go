package logre_client

import (
	"encoding/json"
	"net/http"
	"testing"
)

var testBox = "263048710008603138"

func TestMessage(t *testing.T) {
	Init(testBox)

	logText := "Test: Simple Log"

	response, err := Message(logText)

	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected status code %d received %d", 200, response.StatusCode)
	}

	body := getBodyString(response, t)

	if body.ID.Data.Box.Ref.ID != testBox {
		t.Fatalf("Expected box id %s received %s", testBox, body.ID.Data.Box.Ref.ID)
	}

	if body.ID.Data.Body != logText {
		t.Fatalf("Expected log: '%s' got '%s'", logText, body.ID.Data.Body)
	}
}


func getBodyString(r *http.Response, t *testing.T) *ResponseStringBody {
	 var m ResponseStringBody

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		t.Error(err)
	}

	return &m
}

func getBody(r *http.Response, t *testing.T) *ResponseBody {
	var m ResponseBody

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		t.Error(err)
	}

	return &m
}


type ResponseStringBody struct {
	ID struct {
		Data struct {
			Box struct {
				Ref struct {
					ID string `json:"id"`
				} `json:"@ref"`
			} `json:"box"`
			Body string `json:"body"`
		} `json:"data"`
	} `json:"id"`
}

type ResponseBody struct {
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