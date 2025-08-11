package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	debugLog = log.New(os.Stdout, "\033[36m[debug]\033[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	fatalLog = log.New(os.Stdout, "\033[35m[fatal]\033[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mtx      sync.Mutex
)

var (
	Debug  = debugLog.Println
	Debugf = debugLog.Printf
	Fatal  = fatalLog.Println
	Fatalf = fatalLog.Printf
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	DebugLevel = iota
	InfoLevel
	ErrorLevel
	FatalLevel
	Disabled
)

func SetLevel(level int) {
	mtx.Lock()
	defer mtx.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if DebugLevel < level {
		debugLog.SetOutput(io.Discard)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(io.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(io.Discard)
	}

	if FatalLevel < level {
		fatalLog.SetOutput(io.Discard)
	}
}
