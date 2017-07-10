package logs_test

import (
	"testing"
	"github.com/shellus/pkg/logs"
)

func TestLogs(t *testing.T) {
	logs.Debug("it's okay!")
	logger := logs.SetPrefix("feel: ")
	logger.Debug("it's okay!")
	logs.Debug("it's okay!")
}
