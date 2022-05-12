package logx

import (
	"github.com/rs/zerolog"
	"io"
	"logx/tool"
	"logx/writer"
)

// Logx 包装 zerolog ，适度封装
type Logx struct {
	zeroLog zerolog.Logger
}

// Debug 输出级别为 Debug 的日志
func (l *Logx) Debug(format string, v ...interface{}) {
	l.zeroLog.Debug().Msgf(format, v...)
}

// Info 输出级别为 Info 的日志
func (l *Logx) Info(format string, v ...interface{}) {
	l.zeroLog.Info().Msgf(format, v...)
}

// Warn 输出级别为 Warn 的日志
func (l *Logx) Warn(format string, v ...interface{}) {
	l.zeroLog.Warn().Msgf(format, v...)
}

// Error 输出级别为 Error 的日志
func (l *Logx) Error(err error) {
	l.zeroLog.Error().Msg(err.Error())
}

// NewStdLog 创建输出 text 格式的日志到标准输出的 Logx
func NewStdLog(level string) *Logx {
	return newLog(level, writer.NewStdWriter())
}

// NewFileLog 创建输出 text 格式的日志到文件的 Logx
func NewFileLog(level string, path string) *Logx {
	return newLog(level, writer.NewFileWriter(path))
}

// NewStdJsonLog 创建输出 json 格式的日志到标准输出的 Logx
func NewStdJsonLog(level string) *Logx {
	return newLog(level, writer.NewStdJsonWriter())
}

// NewFileJsonLog 创建输出 json 格式的日志到文件的 Logx
func NewFileJsonLog(level string, path string) *Logx {
	return newLog(level, writer.NewFileJsonWriter(path))
}

// newLog 创建 Lox
func newLog(level string, writer io.Writer) *Logx {
	// 创建 zeroLog
	//zeroLog := zerolog.New(writer).With().CallerWithSkipFrameCount(3).Timestamp().Logger()
	zeroLog := zerolog.New(writer).With().Timestamp().Logger()

	// 设置日志的时间格式
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.99999"

	// 转换日志级别
	zeroLevel := tool.ConvertLevel(level)

	// 设置日志级别
	zerolog.SetGlobalLevel(zeroLevel)

	// 创建 Logx
	return &Logx{
		zeroLog: zeroLog,
	}
}
