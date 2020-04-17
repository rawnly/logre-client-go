package logre

import (
	"encoding/json"
	"github.com/rawnly/go-logre/configuration"
	"net/http"
)

type severity struct {
	Warning string `json:"warning"`
	Info    string `json:"info"`
	Debug   string `json:"debug"`
	Error   string `json:"error"`
	Fatal   string `json:"fatal"`
}

var Severity = severity{
	Warning: "warning",
	Info:    "info",
	Debug:   "debug",
	Error:   "error",
	Fatal:   "fatal",
}

// Initialize the box
func Init(box string) {
	configuration.Config.BoxId = box
}

//func setDebug(debugMode bool) {
//	configuration.Config.Debug = debugMode
//}

// Log with DEBUG severity level
func Debug(i interface{}) (*http.Response, error) {
	data, err := addSeverity(&i, Severity.Debug)

	if err != nil {
		return nil, err
	}

	return log(configuration.Config.BoxId, data)
}

// Log with INFO severity level
func Info(i interface{}) (*http.Response, error) {
	data, err := addSeverity(&i, Severity.Info)

	if err != nil {
		return nil, err
	}

	return log(configuration.Config.BoxId, data)
}

// Log with WARNING severity level
func Warning(i interface{}) (*http.Response, error) {
	data, err := addSeverity(&i, Severity.Warning)

	if err != nil {
		return nil, err
	}

	return log(configuration.Config.BoxId, data)
}

// Log with ERROR severity level
func Error(i interface{}) (*http.Response, error) {
	data, err := addSeverity(&i, Severity.Error)

	if err != nil {
		return nil, err
	}

	return log(configuration.Config.BoxId, data)
}

// Log with FATAL severity level
func Fatal(i interface{}) (*http.Response, error) {
	data, err := addSeverity(&i, Severity.Fatal)

	if err != nil {
		return nil, err
	}

	return log(configuration.Config.BoxId, data)
}

// Log with no severity level
func Log(i interface{}) (*http.Response, error) {
	data, err := json.Marshal(&i)

	if err != nil {
		return nil, err
	}

	return log(configuration.Config.BoxId, data)
}

//
type logPayload struct {
	Message  string `json:"message,omitempty"`
	Severity string `json:"severity,omitempty"`
}

func Message(message string) (*http.Response, error) {
	return logMessage(configuration.Config.BoxId, message)
}

func InfoMessage(message string) (*http.Response, error) {
	var payload = logPayload{
		Message: message,
	}

	return Info(payload)
}

func DebugMessage(message string) (*http.Response, error) {
	var payload = logPayload{
		Message: message,
	}

	return Debug(payload)
}

func WarningMessage(message string) (*http.Response, error) {
	var payload = logPayload{
		Message: message,
	}

	return Warning(payload)
}

func ErrorMessage(message string) (*http.Response, error) {
	var payload = logPayload{
		Message: message,
	}

	return Error(payload)
}

func FatalMessage(message string) (*http.Response, error) {
	var payload = logPayload{
		Message: message,
	}

	return Fatal(payload)
}
