package easylog

import (
	"log"
	"strings"
	"sync"
)

// Levels
const (
	FATAL uint8 = iota
	PANIC
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

// Level strings
const (
	StrFATAL = "FATAL"
	StrPANIC = "PANIC"
	StrERROR = "ERROR"
	StrWARN  = "WARN"
	StrINFO  = "INFO"
	StrDEBUG = "DEBUG"
	StrTRACE = "TRACE"
)

// Map level strings to uint8 values
var lvMap = map[string]uint8{
	StrFATAL: FATAL,
	StrPANIC: PANIC,
	StrERROR: ERROR,
	StrWARN:  WARN,
	StrINFO:  INFO,
	StrDEBUG: DEBUG,
	StrTRACE: TRACE,
}

//type Logger interface {
//	Fatal(msg ...interface{})
//	Panic(msg ...interface{})
//	Error(msg ...interface{})
//	Warn(msg ...interface{})
//	Debug(msg ...interface{})
//	Trace(msg ...interface{})
//	Fatalf(format string, msg ...interface{})
//	Panicf(format string, msg ...interface{})
//	Errorf(format string, msg ...interface{})
//	Warnf(format string, msg ...interface{})
//	Debugf(format string, msg ...interface{})
//	Tracef(format string, msg ...interface{})
//	Level() uint8
//	SetLevel(level string)
//}

type EasyLogger struct {
	sync.Mutex
	level uint8
}

var lg EasyLogger

func init() {

}

func SetLevel(lv string) {
	lg.level = lvMap[strings.ToUpper(lv)]
}

func Logln(lv uint8, msg interface{}) {
	if IsLvEnabled(lv) {
		log.Println(msg)
	}
}

func Fatal(msg interface{}) {
	if IsLvEnabled(FATAL) {
		log.Fatalln(msg)
	}
}

func Fatalf(format string, v ...interface{}) {
	if IsLvEnabled(FATAL) {
		log.Printf(format, v)
	}
}

func Panic(msg interface{}) {
	if IsLvEnabled(PANIC) {
		log.Panicln(msg)
	}
}

func Panicf(format string, v ...interface{}) {
	if IsLvEnabled(PANIC) {
		log.Printf(format, v)
	}
}

func Error(msg ...interface{}) {
	if IsLvEnabled(ERROR) {
		log.Println(msg)
	}
}

func Errorf(format string, v ...interface{}) {
	if IsLvEnabled(ERROR) {
		log.Printf(format, v)
	}
}

func Warn(msg ...interface{}) {
	if IsLvEnabled(WARN) {
		log.Println(msg)
	}
}

func Warnf(format string, v ...interface{}) {
	if IsLvEnabled(WARN) {
		log.Printf(format, v)
	}
}

func Info(msg interface{}) {
	if IsLvEnabled(INFO) {
		log.Println(msg)
	}
}

func Infof(format string, v ...interface{}) {
	if IsLvEnabled(INFO) {
		log.Printf(format, v)
	}
}


func Debug(msg interface{}) {
	if IsLvEnabled(DEBUG) {
		log.Println(msg)
	}
}

func Debugf(format string, v ...interface{}) {
	if IsLvEnabled(DEBUG) {
		log.Printf(format, v)
	}
}

func Trace(msg interface{}) {
	if IsLvEnabled(TRACE) {
		log.Println(msg)
	}
}

func Tracef(format string, v ...interface{}) {
	if IsLvEnabled(TRACE) {
		log.Printf(format, v)
	}
}

func IsLvEnabled(lv uint8) bool {
	return lv <= lg.level
}
