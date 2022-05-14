package logx

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	Warn = iota
	Debug
	Info
	Error
)

const (
	WarnStr  = "WARN"
	DebugStr = "DEBUG"
	InfoStr  = "INFO"
	ErrorStr = "ERROR"
)

// Logx 包装 zerolog ，适度封装
type Logx struct {
	level int
	out   io.Writer
}

// NewStdoutLog 创建输出日志到标准输出的 Logx
func NewStdoutLog(level string) *Logx {
	return &Logx{
		level: convertLevelStr(level),
		out:   os.Stdout,
	}
}

// NewFileLog 创建输出日志到文件的 Logx
func NewFileLog(level string, path string) *Logx {
	return &Logx{
		level: convertLevelStr(level),
		out:   NewFileWriter(path),
	}
}

// output 按一定格式输出日志
func (l *Logx) output(level int, s string) {
	if l.level > level {
		return
	}
	var sb strings.Builder
	timeStr := time.Now().Format("2006-01-02 15:04:05.99999")

	levelStr := convertLevel(level)
	levelStr = fmt.Sprintf("[%5s]", levelStr)

	sb.WriteString(timeStr)
	sb.WriteString(" ")
	sb.WriteString(levelStr)
	sb.WriteString(" ")
	sb.WriteString(s)
	sb.WriteString("\n")

	_, err := l.out.Write([]byte(sb.String()))
	if err != nil {
		log.Println(err)
	}
}

// Warn 输出级别为 Warn 的日志
func (l *Logx) Warn(format string, v ...interface{}) {
	l.output(Warn, fmt.Sprintf(format, v...))
}

// Debug 输出级别为 Debug 的日志
func (l *Logx) Debug(format string, v ...interface{}) {
	l.output(Debug, fmt.Sprintf(format, v...))
}

// Info 输出级别为 Info 的日志
func (l *Logx) Info(format string, v ...interface{}) {
	l.output(Info, fmt.Sprintf(format, v...))
}

// Error 输出级别为 Error 的日志
func (l *Logx) Error(err error) {
	cls := callers()
	l.output(Error, err.Error() + "\n" + cls)
}

// callers 获取 Logx.Error 的调用栈，并格式化成字符串
func callers() string {
	pc := make([]uintptr, 32)
	n := runtime.Callers(3, pc)
	if n == 0 {
		return ""
	}

	pc = pc[:n]
	frames := runtime.CallersFrames(pc)

	var sb strings.Builder
	for {
		frame, more := frames.Next()

		sb.WriteString("=====>>> ")
		sb.WriteString(frame.File)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(frame.Line))

		if !more || frame.Function == "main.main" {
			break
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// convertLevelStr 将字符串类型的 level 转换为 int 类型的 level
func convertLevelStr(level string) int {
	switch level {
	case ErrorStr:
		return Error
	case InfoStr:
		return Info
	case DebugStr:
		return Debug
	case WarnStr:
		return Warn
	}
	return -1
}

// convertLevel 将 int 类型的 level 转换为字符串类型的 level
func convertLevel(level int) string {
	switch level {
	case Error:
		return ErrorStr
	case Info:
		return InfoStr
	case Debug:
		return DebugStr
	case Warn:
		return WarnStr
	}
	return ""
}
