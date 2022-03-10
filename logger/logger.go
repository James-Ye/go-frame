package logger

import (
	"io"
	"os"

	"github.com/James-Ye/go-frame/path_mgr"
)

var logfile string

func Init(appname string) {
	appPath := path_mgr.GetAppPath()

	logfile = appPath
	logfile += "\\log_"
	logfile += appname
	logfile += ".log"
}

func writelog(message string) {
	var f *os.File
	defer f.Close()
	var err error
	if _, err = os.Stat(logfile); os.IsNotExist(err) {
		if f, err = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND, 0666); err != nil {
			panic(err)
		}
	} else if f, err = os.OpenFile(logfile, os.O_APPEND, 0666); err != nil {
		panic(err)
	}
	_, err = io.WriteString(f, message) //写入文件(字符串)
	if err != nil {
		panic(err)
	}
}

func Trace(format string, a ...interface{}) {
	writelog(format)
}

func Debug(format string, a ...interface{}) {

}

func Error(format string, a ...interface{}) {

}

func SetDebug(set bool) {

}
