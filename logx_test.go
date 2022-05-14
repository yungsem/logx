package logx

import (
	"errors"
	"testing"
)

func TestNewLog(t *testing.T) {
	log  := NewStdoutLog("ERROR")

	log.Warn("test %s error", "warn")
	log.Debug("test %s error", "debug")
	log.Info("test %s error", "info")

	err := errors.New("test error")
	log.Error(err)
}
