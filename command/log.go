package command

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Log struct {
	DirName string
}

func (log *Log) SetDirName(name string) {
	log.DirName = name
}

const (
	ChangeLog = iota
	DetailLog
	ErrorLog
)

func (log *Log) GetDirName() string {
	if len(log.DirName) != 0 {
		return log.DirName
	}

	//Mon Jan 2 15:04:05 -0700 MST 2006
	return time.Now().Format("01-02-15-04-05")
}

func (log *Log) initStandardFile() (changeLog, detailLog, errLog *os.File) {
	dir := log.GetDirName()
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

func (log *Log) GetLog(index int) *os.File {

	logOnce.Do(func() {
		changeLog, detailLog, errorLog = log.initStandardFile()
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

func (log *Log) WriteLog(index int, piece ...interface{}) {
	file := log.GetLog(index)
	str := strings.Trim(fmt.Sprintf("%v", piece), "[]")
	n, err := file.WriteString(str + "\n")
	fmt.Println(n ,err)
}

func (log *Log) Close() {
	if changeLog == nil {
		return
	}

	changeLog.Close()
	detailLog.Close()
	errorLog.Close()
}
