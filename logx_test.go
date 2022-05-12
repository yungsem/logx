package logx

import (
	"errors"
	"testing"
)

func TestNewLog(t *testing.T) {
	log := NewFileJsonLog("DEBUG", "log")

	err := errors.New("test error")

	log.Debug("test debug")
	log.Warn("test warn")
	log.Info("test info")
	log.Error(err)
}
