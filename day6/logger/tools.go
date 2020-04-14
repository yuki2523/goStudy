package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

func checkFileSize(f *os.File) bool {
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		panic("无法获取当前文件具体信息,解析文件大小失败")
	}
	if fileInfo.Size() >= 1*1024 { // 文件大小大于 1kb 那么需要重新切割文件
		return true
	}
	return false
}

// 按大小切割文件,每次切割时给原文件搭配时间重命名,之后再接着写
func (l *Logger) spitFile(flag bool) { // flag 是判断现在是切割一般的 .log 文件 还是切割 log.err 文件
	var (
		fp       = l.filePath
		fn       = l.fileName
		fullname string
	)
	if flag {
		l.logFile.Close()
		fullname = l.logFullPath
	} else {
		l.errFile.Close()
		fullname = l.errFullPath
	}

	newName := path.Join(fp, time.Now().Format("060102150405")+fn)
	if !flag {
		newName += ".err"
	}

	os.Rename(fullname, newName)
	fileObj, err := os.OpenFile(fullname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic("新切割出来的文件打开失败")
	}

	if flag {
		l.logFile = fileObj
	} else {
		l.errFile = fileObj
	}
}
