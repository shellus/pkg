package logger_test

import (
	"testing"
	"github.com/shellus/pkg/logs/logger"
)

func TestLogger(t *testing.T) {
	log := logger.NewLogger()
	log.Debug("abc")
}