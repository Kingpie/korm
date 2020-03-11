package log

import "testing"

func TestSetLevel(t *testing.T) {
	SetLevel(InfoLevel)
	Error("info")
	Info("info")

	SetLevel(ErrorLevel)
	Error("error")
	Info("error")

	SetLevel(Disabled)
	Error("disabled")
	Info("disabled")
}
