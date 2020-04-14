package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// 处理业务的函数,输出信息的部分都在这里
func (l *Logger) messageGetter(level Level, format string, args ...interface{}) {
	now := time.Now()

	// 判断等级并给输出等级 currentLevel 赋值, DEBUG 对应 debug
	// INFO 对应 info ... , 这部分就是做一个对应
	var currentLevel string
	switch level {
	case DEBUG:
		currentLevel = "DEBUG"
	case INFO:
		currentLevel = "INFO"
	case WARNING:
		currentLevel = "WARNING"
	case ERROR:
		currentLevel = "ERROR"
	}

	// 下面这部分使用 runtime.Caller 获取打印日志的位置、文件和函数相关信息
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("无法解析当前位置的具体信息")
	}

	if l.level <= level { // 先判断初始化时的日志等级,根据等级决定该句日志是否写入
		if l.console { // 判断是否输出在控制台
			fmt.Printf("[%s] [%s]\t[fileName:%s  function:%s  line:%d]\t", now.Format("2006/01/02 15:04:05"), currentLevel, path.Base(file), runtime.FuncForPC(pc).Name(), line)
			fmt.Printf(format, args...)
		} else { // 如果不是输出在控制台,那么向文件里输出
			if checkFileSize(l.logFile) { // 如果需要切割文件
				// l.logFile.Close()
				// newName := path.Join(l.filePath, time.Now().Format("06-01-02 15:04:05 ") + l.fileName)
				// os.Rename(l.logFullPath, newName)
				// fileObj, err := os.OpenFile(l.logFullPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
				// if err != nil {
				// 	panic("新切割出来的文件打开失败")
				// }
				// l.logFile = fileObj
				l.spitFile(true)
			} else {
				fmt.Fprintf(l.logFile, "[%s] [%s]\t[fileName:%s  function:%s  line:%d]\t", now.Format("2006/01/02 15:04:05"), currentLevel, path.Base(file), runtime.FuncForPC(pc).Name(), line)
				fmt.Fprintf(l.logFile, format, args...)
			}

			if level == ERROR { // 如果是 error 等级的,再向 error 等级的文件再输出一份信息,和上面统一的 l.logFile 不在一起处理
				if checkFileSize(l.errFile) {
					l.spitFile(false)
				} else {
					fmt.Fprintf(l.errFile, "[%s] [%s]\t[fileName:%s  function:%s  line:%d]\t", now.Format("2006/01/02 15:04:05"), currentLevel, path.Base(file), runtime.FuncForPC(pc).Name(), line)
					fmt.Fprintf(l.errFile, format, args...)
				}
			}
		}
	}
}

// Debug debug等级,使用方式同 fmt.Printf
func (l *Logger) Debug(format string, args ...interface{}) {
	l.messageGetter(DEBUG, format, args...)
}

// Info info等级,使用方式同 fmt.Printf
func (l *Logger) Info(format string, args ...interface{}) {
	l.messageGetter(INFO, format, args...)
}

// Warning warning等级,使用方式同 fmt.Printf
func (l *Logger) Warning(format string, args ...interface{}) {
	l.messageGetter(WARNING, format, args...)
}

// Error error等级,使用方式同 fmt.Printf
func (l *Logger) Error(format string, args ...interface{}) {
	l.messageGetter(ERROR, format, args...)
}
