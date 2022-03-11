package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/James-Ye/go-frame/goroutine_mgr"
	"golang.org/x/sys/windows"
)

var g_log *log
var once sync.Once

func GetInstance() *log {
	once.Do(func() {
		g_log = &log{}
	})
	return g_log
}

type log struct {
	logfile string
	level   LEVEL
}

type LEVEL = uint32

const (
	LOG_DEBUG LEVEL = 0 //调试信息
	LOG_TRACE LEVEL = 1 //调试信息
	LOG_WARN  LEVEL = 2 //警告
	LOG_ERROR LEVEL = 3 //错误
	LOG_NONE  LEVEL = 4 //不记录日志
)

var level_string = [4]string{"Debug", "Trace", "Warn", "Error"}

func (l *log) setpath(logpath string) {
	l.logfile = logpath
}

func (l *log) setLevel(level uint32) {
	l.level = level
}

func (l *log) writelog(message string) {
	var f *os.File
	defer f.Close()
	var err error
	if _, err = os.Stat(l.logfile); os.IsNotExist(err) {
		if f, err = os.OpenFile(l.logfile, os.O_CREATE|os.O_APPEND, 0666); err != nil {
			panic(err)
		}
	} else if f, err = os.OpenFile(l.logfile, os.O_APPEND, 0666); err != nil {
		panic(err)
	}
	_, err = io.WriteString(f, message+"\n") //写入文件(字符串)
	if err != nil {
		panic(err)
	}
}

func (l *log) report(kind string, msg string) {
	s := fmt.Sprintf("%s (pid:%d, tid:%d) %s->%s\r\n",
		time.Now().Format("2006-01-02 15:04:05"),
		windows.GetCurrentProcessId(),
		goroutine_mgr.GetCurrentGoid(),
		kind,
		msg)
	l.writelog(s)
}

func (l *log) log(level uint32, format string, a ...interface{}) {
	if l.level > level {
		return
	}

	if len(a) == 0 {
		l.report(level_string[level], format)
	} else {
		message := fmt.Sprintf(format, a...)
		l.report(level_string[level], message)
	}
}

func SetPath(appname string) {
	GetInstance().setpath(appname)
}

func SetLevel(level LEVEL) {
	GetInstance().setLevel(level)
}

func Trace(format string, a ...interface{}) {
	GetInstance().log(LOG_TRACE, format, a...)
}

func Debug(format string, a ...interface{}) {
	GetInstance().log(LOG_DEBUG, format, a...)
}

func Error(format string, a ...interface{}) {
	GetInstance().log(LOG_ERROR, format, a...)
}

func Warn(format string, a ...interface{}) {
	GetInstance().log(LOG_WARN, format, a...)
}
