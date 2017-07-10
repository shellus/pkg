package logger

import (
	"log"
	"os"
	"io"
	"sync"
)

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

// levelLogLogger is defined to implement log.Logger
// the real log level will be LevelEmergency
const levelLoggerImpl = -1

// Legacy log level constants to ensure backwards compatibility.
const (
	LevelInfo = LevelInformational
	LevelTrace = LevelDebug
	LevelWarn = LevelWarning
)

type Logger struct {
	log    *log.Logger
	prefix string
	out    io.Writer
	level  int
	mu     sync.Mutex
}

func NewLogger() *Logger {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	return &Logger{log:logger, level: LevelDebug, }
}

func (bl Logger) SetPrefix(prefix string) (*Logger) {
	bl.mu.Lock()
	defer bl.mu.Unlock()
	bl.prefix = prefix
	return &bl
}
func (bl Logger) SetOutput(w io.Writer) (*Logger)  {
	bl.mu.Lock()
	defer bl.mu.Unlock()
	bl.out = w
	bl.log.SetOutput(w)
	return &bl
}

func (bl *Logger) writeMsg(logLevel int, format string, v ...interface{}) {
	if len(bl.prefix) != 0 {
		format = bl.prefix + format
	}
	bl.log.Printf(format, v...)
}

// Alert Log ALERT level message.
func (bl *Logger) Alert(format string, v ...interface{}) {
	if LevelAlert > bl.level {
		return
	}
	bl.writeMsg(LevelAlert, format, v...)
}

// Critical Log CRITICAL level message.
func (bl *Logger) Critical(format string, v ...interface{}) {
	if LevelCritical > bl.level {
		return
	}
	bl.writeMsg(LevelCritical, format, v...)
}

// Error Log ERROR level message.
func (bl *Logger) Error(format string, v ...interface{}) {
	if LevelError > bl.level {
		return
	}
	bl.writeMsg(LevelError, format, v...)
}

// Warning Log WARNING level message.
func (bl *Logger) Warning(format string, v ...interface{}) {
	if LevelWarn > bl.level {
		return
	}
	bl.writeMsg(LevelWarn, format, v...)
}

// Notice Log NOTICE level message.
func (bl *Logger) Notice(format string, v ...interface{}) {
	if LevelNotice > bl.level {
		return
	}
	bl.writeMsg(LevelNotice, format, v...)
}

// Informational Log INFORMATIONAL level message.
func (bl *Logger) Informational(format string, v ...interface{}) {
	if LevelInfo > bl.level {
		return
	}
	bl.writeMsg(LevelInfo, format, v...)
}

// Debug Log DEBUG level message.
func (bl *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > bl.level {
		return
	}
	bl.writeMsg(LevelDebug, format, v...)
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func (bl *Logger) Warn(format string, v ...interface{}) {
	if LevelWarn > bl.level {
		return
	}
	bl.writeMsg(LevelWarn, format, v...)
}

// Info Log INFO level message.
// compatibility alias for Informational()
func (bl *Logger) Info(format string, v ...interface{}) {
	if LevelInfo > bl.level {
		return
	}
	bl.writeMsg(LevelInfo, format, v...)
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func (bl *Logger) Trace(format string, v ...interface{}) {
	if LevelDebug > bl.level {
		return
	}
	bl.writeMsg(LevelDebug, format, v...)
}
