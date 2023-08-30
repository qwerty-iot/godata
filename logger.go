package godata

import (
	"fmt"
	"net/http"
	"time"
)

const (
	LogLevelError string = "error"
	LogLevelWarn  string = "warn"
	LogLevelInfo  string = "info"
	LogLevelDebug string = "debug"
)

type LogFunc func(ts time.Time, level string, req *http.Request, err error, log string)

var logFunc LogFunc = defaultLogFunc
var logLevel int = 0

func SetLogFunc(lf LogFunc) {
	logFunc = lf
}

func SetLogLevel(level string) {
	switch level {
	case LogLevelError:
		logLevel = 1
	case LogLevelWarn:
		logLevel = 2
	case LogLevelInfo:
		logLevel = 3
	case LogLevelDebug:
		logLevel = 4
	default:
		logLevel = 0
	}
}

func defaultLogFunc(ts time.Time, level string, req *http.Request, err error, l string) {
	loc := ""
	if req != nil && len(req.RemoteAddr) != 0 {
		loc = fmt.Sprintf("%s][%s][%s", req.RemoteAddr, req.Method, req.URL.RequestURI())
	}
	if err != nil {
		fmt.Println(" [" + level + "] [" + loc + "] " + l + "(err: " + err.Error() + ")")
	} else {
		fmt.Println(" [" + level + "] [" + loc + "] " + l)
	}
}

func logError(req *http.Request, err error, f string, args ...interface{}) {
	if logLevel < 1 {
		return
	}
	logFunc(time.Now(), LogLevelError, req, err, fmt.Sprintf(f, args...))
}

func logWarn(req *http.Request, err error, f string, args ...interface{}) {
	if logLevel < 2 {
		return
	}
	logFunc(time.Now(), LogLevelWarn, req, err, fmt.Sprintf(f, args...))
}

func logDebug(req *http.Request, err error, f string, args ...interface{}) {
	if logLevel < 4 {
		return
	}
	logFunc(time.Now(), LogLevelDebug, req, err, fmt.Sprintf(f, args...))
}
