package logger

import (
	"errors"
	"os"
	"path"
	"strings"
)

type files struct {
	logFile *os.File
	errFile *os.File
}

// Logger Logger 结构体
// level为等级,filePath为存储路径,fileName为文件名,file里放的是两个 *os.File
type Logger struct {
	level       Level
	filePath    string
	fileName    string
	logFullPath string
	errFullPath string
	console     bool
	files
}

// Level 等级(自定义)类型
type Level uint16

const (
	// UNKNOWN 未知
	UNKNOWN Level = iota
	// DEBUG 调试
	DEBUG
	// INFO 正常输出
	INFO
	// WARNING 警报
	WARNING
	// ERROR 错误
	ERROR
)

func levelGetter(l string) (Level, error) {
	l = strings.ToLower(l)
	switch l {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	default:
		return UNKNOWN, errors.New("传入了一个不知道的等级类型")
	}
}

// NewLogger 返回 Logger 结构体
// 第一参数为限制等级,字符串,debug、info、warning、error,不限大小写
// 第二个参数为输出路径
// 第三个参数为输出的文件名
// 第四个参数是个 bool 值,决定输出的结果是展现在控制台上还是写入到文件里
func NewLogger(l string, filePath, fileName string, console bool) *Logger {
	lev, err := levelGetter(l)
	if err != nil {
		panic("Logger创建失败！")
	}

	log := &Logger{
		level:    lev,
		filePath: filePath,
		fileName: fileName,
		console:  console,
	}
	log.initLogger()

	return log
}

// logger 的初始化,启动 logFile 和 errFile 的 IO
func (l *Logger) initLogger() {
	logFileName := path.Join(l.filePath, l.fileName)
	errFileName := logFileName + ".err"

	l.logFullPath = logFileName
	l.errFullPath = errFileName
	fileObj, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("open logFile error")
	}
	l.logFile = fileObj

	fileObj, err = os.OpenFile(errFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("open logFile error")
	}
	l.errFile = fileObj
}

// Close 使用完这个日志器后关闭 logFile 和 errFile 的 IO
func (l *Logger) Close() {
	l.logFile.Close()
	l.errFile.Close()
}
