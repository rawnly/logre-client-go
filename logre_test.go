package logre

import (
	"encoding/json"
	"net/http"
	"testing"
)

var testBox = "263048710008603138"

type StandardLog struct {
	Message  string `json:"message,omitempty"`
	Severity string `json:"severity,omitempty"`
}

func TestInfo(t *testing.T) {
	Init(testBox)

	payload := StandardLog{
		Message: "Test: Hello World",
	}

	response, err := Info(payload)

	testResponse(payload.Message, Severity.Info, response, err, t)
}

func TestDebug(t *testing.T) {
	Init(testBox)

	payload := StandardLog{
		Message: "Test: Hello World",
	}

	response, err := Debug(payload)

	testResponse(payload.Message, Severity.Debug, response, err, t)
}

func TestError(t *testing.T) {
	Init(testBox)

	payload := StandardLog{
		Message: "Test: Hello World",
	}

	response, err := Error(payload)

	testResponse(payload.Message, Severity.Error, response, err, t)
}

func TestWarning(t *testing.T) {
	Init(testBox)

	payload := StandardLog{
		Message: "Test: Hello World",
	}

	response, err := Warning(payload)

	testResponse(payload.Message, Severity.Warning, response, err, t)
}

func TestFatal(t *testing.T) {
	Init(testBox)

	payload := StandardLog{
		Message: "Test: Hello World",
	}

	response, err := Fatal(payload)

	testResponse(payload.Message, Severity.Fatal, response, err, t)
}

func TestLog(t *testing.T) {
	Init(testBox)

	payload := StandardLog{
		Message: "Test: Hello World",
	}

	response, err := Log(payload)

	testResponse(payload.Message, "", response, err, t)
}

// Messages
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

	body := GetLogBody(response, t)

	if body.ID.Data.Box.Ref.ID != testBox {
		t.Fatalf("Expected box id %s received %s", testBox, body.ID.Data.Box.Ref.ID)
	}

	if s := body.ID.Data.Body.(string); s != logText {
		t.Fatalf("Expected log: '%s' got '%s'", logText, s)
	}
}

func TestInfoMessage(t *testing.T) {
	Init(testBox)

	message := "Test: Simple Log"

	response, err := InfoMessage(message)

	testResponse(message, Severity.Info, response, err, t)
}

func TestDebugMessage(t *testing.T) {
	Init(testBox)

	message := "Test: Simple Log"

	response, err := DebugMessage(message)

	testResponse(message, Severity.Debug, response, err, t)
}

func TestWarningMessage(t *testing.T) {
	Init(testBox)

	message := "Test: Simple Log"

	response, err := WarningMessage(message)

	testResponse(message, Severity.Warning, response, err, t)
}

func TestErrorMessage(t *testing.T) {
	Init(testBox)

	message := "Test: Simple Log"

	response, err := ErrorMessage(message)

	testResponse(message, Severity.Error, response, err, t)
}

func TestFatalMessage(t *testing.T) {
	Init(testBox)

	message := "Test: Simple Log"

	response, err := FatalMessage(message)

	testResponse(message, Severity.Fatal, response, err, t)
}

func testResponse(expected string, severity string, response *http.Response, err error, t *testing.T) {
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected status code %d received %d", 200, response.StatusCode)
	}

	body := GetLogBody(response, t)

	if body.ID.Data.Box.Ref.ID != testBox {
		t.Fatalf("Expected box id %s received %s", testBox, body.ID.Data.Box.Ref.ID)
	}

	var m map[string]string

	data, e := json.Marshal(body.ID.Data.Body)

	if e != nil {
		t.Error(e)
		t.Fail()
	}

	if err := json.Unmarshal(data, &m); err != nil {
		t.Error(err)
		t.Fail()
	}

	if logMessage := m["message"]; logMessage != expected {
		t.Fatalf("Expected log: '%s' got '%s'", expected, logMessage)
	}

	if s := m["severity"]; s != severity && severity != "" {
		t.Fatalf("Expected log severity to be '%s' but got '%s'", severity, s)
	}
}
