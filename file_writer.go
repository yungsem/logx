package logx

import (
	"io"
	"os"
	"strings"
	"time"
)

// fileWriter 将日志以普通文本的格式输出到文件中
type fileWriter struct {
	path string
}

// Write 将日志以普通文本的格式输出到文件中
func (fw fileWriter) Write(p []byte) (n int, err error) {
	// 提取日志级别
	level := resolveLevel(string(p))

	// 创建日志文件
	file, err := createLogFile(level, fw.path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// 写入
	return file.Write(p)
}

// NewFileWriter 创建 fileWriter
func NewFileWriter(path string) io.Writer {
	return fileWriter{path: path}
}

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
	case WarnStr:
		fileName = "warn.log"
	case DebugStr:
		fileName = "debug.log"
	case InfoStr:
		fileName = "info.log"
	case ErrorStr:
		fileName = "error.log"
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

// resolveLevel 从日志内容中解析出日志级别
func resolveLevel(msg string) string {
	if strings.Contains(msg, "[ WARN]") {
		return WarnStr
	}
	if strings.Contains(msg, "[DEBUG]") {
		return DebugStr
	}
	if strings.Contains(msg, "[ INFO]") {
		return InfoStr
	}
	if strings.Contains(msg, "[ERROR]") {
		return ErrorStr
	}
	return ""
}
