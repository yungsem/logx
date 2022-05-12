package tool

import (
	"github.com/rs/zerolog"
	"strings"
)

const (
	LevelError = "ERROR"
	LevelInfo  = "INFO"
	LevelDebug = "DEBUG"
	LevelWarn  = "WARN"
)

// ConvertLevel 将字符串类型的 level 转换为 zerolog.Level
func ConvertLevel(level string) zerolog.Level {
	switch level {
	case LevelError:
		return zerolog.ErrorLevel
	case LevelInfo:
		return zerolog.InfoLevel
	case LevelDebug:
		return zerolog.DebugLevel
	case LevelWarn:
		return zerolog.WarnLevel
	}
	return -1
}

// ConvertZeroLevel 将 zerolog.Level 转换为字符串类型的 level
func ConvertZeroLevel(zeroLevel zerolog.Level) string {
	switch zeroLevel {
	case zerolog.ErrorLevel:
		return LevelError
	case zerolog.InfoLevel:
		return LevelInfo
	case zerolog.DebugLevel:
		return LevelDebug
	case zerolog.WarnLevel:
		return LevelWarn
	}
	return ""
}

// ResolveLevel 从日志内容中解析出日志级别
func ResolveLevel(msg string) string {
	if strings.Contains(msg, "[ERROR]") {
		return LevelError
	} else if strings.Contains(msg, "[ INFO]") {
		return LevelInfo
	} else if strings.Contains(msg, "[DEBUG]") {
		return LevelDebug
	} else if strings.Contains(msg, "[ WARN]") {
		return LevelWarn
	}
	return ""
}
