package logger

import (
	"io"
	"os"
)

var logfile string

func Init(appname string) {
	logfile = "C:\\ProgramData\\"
	logfile += "\\"
	logfile += appname
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
