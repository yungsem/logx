package writer

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"logx/tool"
	"os"
	"strings"
	"time"
)

const (

	// 输出类型：文件/标准输出
	fileOutput = "file"
	stdOutput  = "std"
	// 输出格式：text/json
	textFormat = "text"
	jsonFormat = "json"
)

// createLogFile 根据日志等级创建日志文件
func createLogFile(level string, path string) (*os.File, error) {

	// 目录分割符
	sep := string(os.PathSeparator)

	// 获取当前日期
	now := time.Now()

	// 中间目录名
	today := now.Format("2006-01-02")

	// 文件名前缀
	todayTight := now.Format("20060102")

	// 文件名后缀
	var fileName string
	switch level {
	case tool.LevelError:
		fileName = "error.log"
	case tool.LevelInfo:
		fileName = "info.log"
	case tool.LevelDebug:
		fileName = "debug.log"
	case tool.LevelWarn:
		fileName = "warn.log"
	}

	// 目录路径
	dirPath := path + sep + today + sep

	// 文件路径
	filePath := dirPath + todayTight + "_" + fileName

	// 创建目录，如果目录已经存在，则什么都不做
	err := os.MkdirAll(dirPath, os.ModeDir)
	if err != nil {
		return nil, err
	}

	// 创建文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// NewStdWriter 创建输出 text 格式的日志到标准输出的 writer
func NewStdWriter() io.Writer {
	return newStdWriter(textFormat)
}

// NewStdJsonWriter 创建输出 json 格式的日志到标准输出的 writer
func NewStdJsonWriter() io.Writer {
	return newStdWriter(jsonFormat)
}

// NewFileWriter 创建输出 text 格式的日志到文件的 writer
func NewFileWriter(path string) io.Writer {
	return newFileWriter(path, textFormat)
}

// NewFileJsonWriter 创建输出 json 格式的日志到文件的 writer
func NewFileJsonWriter(path string) io.Writer {
	return newFileWriter(path, jsonFormat)
}

// newStdWriter 创建用于输出日志到标准输出的 writer
func newStdWriter(formatType string) io.Writer {
	if formatType == textFormat {
		// 输出格式：text
		// 创建 ConsoleWriter
		consoleWriter := newConsoleWriter()
		// 输出到标准输出
		consoleWriter.Out = os.Stdout
		return consoleWriter
	} else {
		// 输出格式：json
		return os.Stdout
	}
}

// newFileWriter 创建用于输出日志到文件的 writer
func newFileWriter(path string, formatType string) io.Writer {
	if formatType == textFormat {
		// 输出格式：text
		// 创建 ConsoleWriter
		consoleWriter := newConsoleWriter()
		// 输出到文件
		tw := textWriter{
			path: path,
		}
		consoleWriter.Out = tw
		return consoleWriter
	} else {
		// 输出格式：json
		return jsonWriter{
			path: path,
		}
	}
}

// newConsoleWriter 创建 ConsoleWriter
func newConsoleWriter() zerolog.ConsoleWriter {
	cw := zerolog.ConsoleWriter{
		NoColor:    true,
		TimeFormat: "2006-01-02 15:04:05.999",
	}
	cw.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[%5s]", i))
	}
	return cw
}
