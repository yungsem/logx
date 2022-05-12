package writer

import "github.com/yungsem/logx/tool"

// textWriter 将日志以普通文本的格式输出到文件中
type textWriter struct {
	path string
}

// Write 将日志以普通文本的格式输出到文件中
func (tw textWriter) Write(p []byte) (n int, err error) {
	// 提取日志级别
	level := tool.ResolveLevel(string(p))

	// 创建日志文件
	file, err := createLogFile(level, tw.path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// 写入
	return file.Write(p)
}
