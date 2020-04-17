package logger

import (
	logre "logre-client"
	"net/http"
)

var boxId string

func Init(box string) {
	boxId = box
}

func Debug(message string) (*http.Response, error) {
	return logre.Log(boxId, message, "debug")
}

func Info(message string) (*http.Response, error) {
	return logre.Log(boxId, message, "info")
}

func Warning(message string) (*http.Response, error) {
	return logre.Log(boxId, message, "warning")
}

func Error(message string) (*http.Response, error) {
	return logre.Log(boxId, message, "error")
}

func Fatal(message string) (*http.Response, error) {
	return logre.Log(boxId, message, "fatal")
}