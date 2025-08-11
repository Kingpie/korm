package log

import "testing"

func TestSetLevel(t *testing.T) {
	SetLevel(DebugLevel)
	Debug("debug")
	Error("debug")
	Info("debug")
	Fatal("debug")

	SetLevel(InfoLevel)
	Error("info")
	Info("info")

	SetLevel(ErrorLevel)
	Error("error")
	Info("error")

	SetLevel(FatalLevel)
	Error("fatal")
	Info("fatal")
	Fatal("fatal")

	SetLevel(Disabled)
	Error("disabled")
	Info("disabled")
	Fatal("disabled")
}
