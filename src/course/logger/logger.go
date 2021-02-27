package logger

// 声明日志写入器
type LogWriter interface {
	Write(data interface{}) error
}

// 日志器
type Logger struct {
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		_ = writer.Write(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}
