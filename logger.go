package easylog

import (
	"io"
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
var lvAtoi = map[string]uint8{
	StrFATAL: FATAL,
	StrPANIC: PANIC,
	StrERROR: ERROR,
	StrWARN:  WARN,
	StrINFO:  INFO,
	StrDEBUG: DEBUG,
	StrTRACE: TRACE,
}

// Map level uint8 values to strings
var lvItoa = map[uint8]string{
	FATAL: StrFATAL,
	PANIC: StrPANIC,
	ERROR: StrERROR,
	WARN:  StrWARN,
	INFO:  StrINFO,
	DEBUG: StrDEBUG,
	TRACE: StrTRACE,
}

// These flags define which text to prefix to each log entry generated by the LeveledLogger,
// including flags inherited from std pkg log and flags defined by EasyLog (after LUTC).
const (
	Ldate         = 1 << iota // the date in the local time zone: 2009/01/23
	Ltime                     // the time in the local time zone: 01:23:23
	Lmicroseconds             // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                 // full file name and line number: /a/b/c/d.go:23
	Lshortfile                // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                      // if Ldate or Ltime is set, use UTC rather than the local time zone
	Llevel                    // level of log: [ERROR] [INFO]
	//Ldaily                    // one log file one day
)

type LeveledLogger interface {
	Fatal(msg ...interface{})
	Panic(msg ...interface{})
	Error(msg ...interface{})
	Warn(msg ...interface{})
	Info(msg ...interface{})
	Debug(msg ...interface{})
	Trace(msg ...interface{})
	Fatalf(format string, msg ...interface{})
	Panicf(format string, msg ...interface{})
	Errorf(format string, msg ...interface{})
	Warnf(format string, msg ...interface{})
	Infof(format string, msg ...interface{})
	Debugf(format string, msg ...interface{})
	Tracef(format string, msg ...interface{})
	Flag() int
	SetFlag(flag int) LeveledLogger
	Writer() io.Writer
	SetWriter(w io.Writer) LeveledLogger
	Level() uint8
	SetLevel(lv string) LeveledLogger
	isLevelEnabled(lv uint8) bool
}

var lg LeveledLogger

func SetLevel(lv string) {
	lg.SetLevel(lv)
}

func SetWritter(w io.Writer) {
	lg.SetWriter(w)
}

func SetFlag(flag int) {
	lg.SetFlag(flag)
}

func Fatal(msg ...interface{}) {
	lg.Fatal(msg...)
}

func Fatalf(format string, msg ...interface{}) {
	lg.Fatalf(format, msg...)
}

func Panic(msg ...interface{}) {
	lg.Panic(msg...)
}

func Panicf(format string, msg ...interface{}) {
	lg.Panicf(format, msg...)
}

func Error(msg ...interface{}) {
	lg.Error(msg...)
}

func Errorf(format string, msg ...interface{}) {
	lg.Errorf(format, msg...)
}

func Warn(msg ...interface{}) {
	lg.Warn(msg...)
}

func Warnf(format string, msg ...interface{}) {
	lg.Warnf(format, msg...)
}

func Info(msg ...interface{}) {
	lg.Info(msg...)
}

func Infof(format string, msg ...interface{}) {
	lg.Infof(format, msg)
}

func Debug(msg ...interface{}) {
	lg.Debug(msg...)
}

func Debugf(format string, msg ...interface{}) {
	lg.Debugf(format, msg)
}

func Trace(msg ...interface{}) {
	lg.Trace(msg...)
}

func Tracef(format string, msg ...interface{}) {
	lg.Tracef(format, msg)
}
