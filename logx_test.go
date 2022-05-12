package logx

import (
	"errors"
	"testing"
)

func TestNewLog(t *testing.T) {
	log := NewStdLog("DEBUG")

	err := errors.New("test error")

	log.Debug("test debug")
	log.Warn("test warn")
	log.Info("test info")
	log.Error(err)
}
