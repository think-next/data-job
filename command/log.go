package command

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type log struct {
	DirName string
}

var localLog *log
var localLogOnce sync.Once

func GetCmdLog() *log {
	localLogOnce.Do(func() {
		localLog = &log{}
	})

	return localLog
}

func (localLog *log) SetDirName(name string) {
	localLog.DirName = name
}

const (
	ChangeLog = iota
	DetailLog
	ErrorLog
)

func (localLog *log) getDirName() string {
	if len(localLog.DirName) != 0 {
		return localLog.DirName
	}

	//Mon Jan 2 15:04:05 -0700 MST 2006
	return time.Now().Format("01-02-15-04-05")
}

func (localLog *log) initStandardFile() (changeLog, detailLog, errLog *os.File) {
	dir := localLog.getDirName()
	os.Mkdir("./"+dir, os.ModePerm)

	changePath := fmt.Sprintf("./%s/change.txt", dir)
	detailPath := fmt.Sprintf("./%s/detail.txt", dir)
	errPath := fmt.Sprintf("./%s/error.txt", dir)

	var err error
	changeLog, err = os.Create(changePath)
	if err != nil {
		panic("create change.txt failure")
	}

	detailLog, err = os.Create(detailPath)
	if err != nil {
		panic("create detail.txt failure")
	}

	errLog, err = os.Create(errPath)
	if err != nil {
		panic("create error.txt failure")
	}
	return
}

var logOnce sync.Once
var changeLog, detailLog, errorLog *os.File

func (localLog *log) GetLog(index int) *os.File {

	logOnce.Do(func() {
		changeLog, detailLog, errorLog = localLog.initStandardFile()
	})

	switch index {
	case ChangeLog:
		return changeLog
	case DetailLog:
		return detailLog
	case ErrorLog:
		return errorLog
	}

	panic(fmt.Sprintf("error index %d", index))
}

func (localLog *log) WriteLog(index int, piece ...interface{}) {
	file := localLog.GetLog(index)
	str := strings.Trim(fmt.Sprintf("%v", piece), "[]")
	n, err := file.WriteString(str + "\n")
	fmt.Println(n, err)
}

func (localLog *log) Close() {
	if changeLog == nil {
		return
	}

	changeLog.Close()
	detailLog.Close()
	errorLog.Close()
}
