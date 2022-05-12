package writer

import (
	"github.com/rs/zerolog"
	"io"
	"logx/tool"
)

// jsonWriter 实现了 zerolog 中的 LevelWriter 接口
// 将日志以 json 的格式输出到文件中
type jsonWriter struct {
	path string
	io.Writer
}

// WriteLevel 将日志以 json 的格式输出到文件中
func (jw jsonWriter) WriteLevel(zeroLevel zerolog.Level, p []byte) (n int, err error) {
	// 转换 level
	level := tool.ConvertZeroLevel(zeroLevel)

	// 创建日志文件
	file, err := createLogFile(level, jw.path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// 写入
	return file.Write(p)
}
