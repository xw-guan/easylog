package easylog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type EasyLogger struct {
	sync.Mutex
	l     *log.Logger // delegate
	level uint8
	flag  int
}

func NewEasyLogger() LeveledLogger {
	return &EasyLogger{l: log.New(os.Stderr, "", log.LstdFlags), level: WARN}
}

func (logger *EasyLogger) BuildMessage(level uint8, msg string, sb *strings.Builder) *strings.Builder {
	if logger.flag&Llevel != 0 {
		sb.WriteString(" [")
		sb.WriteString(lvItoa[level])
		sb.WriteString("] ")
	}
	sb.WriteString(msg)
	return sb
}

func (logger *EasyLogger) Message(level uint8, msg ...interface{}) string {
	var sb strings.Builder
	return logger.BuildMessage(level, fmt.Sprint(msg...), &sb).String()
}

func (logger *EasyLogger) Messagef(level uint8, format string, msg ...interface{}) string {
	var sb strings.Builder
	return logger.BuildMessage(level, fmt.Sprintf(format, msg...), &sb).String()
}

func (logger *EasyLogger) Fatal(msg ...interface{}) {
	if logger.isLevelEnabled(FATAL) {
		log.Fatal(logger.Message(FATAL, msg...))
	}
}

func (logger *EasyLogger) Panic(msg ...interface{}) {
	if logger.isLevelEnabled(PANIC) {
		log.Panic(logger.Message(PANIC, msg...))
	}
}

func (logger *EasyLogger) Error(msg ...interface{}) {
	if logger.isLevelEnabled(ERROR) {
		log.Print(logger.Message(ERROR, msg...))
	}
}

func (logger *EasyLogger) Warn(msg ...interface{}) {
	if logger.isLevelEnabled(WARN) {
		log.Print(logger.Message(WARN, msg...))
	}
}

func (logger *EasyLogger) Info(msg ...interface{}) {
	if logger.isLevelEnabled(INFO) {
		log.Print(logger.Message(INFO, msg...))
	}
}

func (logger *EasyLogger) Debug(msg ...interface{}) {
	if logger.isLevelEnabled(DEBUG) {
		log.Print(logger.Message(DEBUG, msg...))
	}
}

func (logger *EasyLogger) Trace(msg ...interface{}) {
	if logger.isLevelEnabled(TRACE) {
		log.Print(logger.Message(TRACE, msg...))
	}
}

func (logger *EasyLogger) Fatalf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(FATAL) {
		log.Fatal(logger.Messagef(FATAL, format, msg...))
	}
}

func (logger *EasyLogger) Panicf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(PANIC) {
		log.Panic(logger.Messagef(PANIC, format, msg...))
	}
}

func (logger *EasyLogger) Errorf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(ERROR) {
		log.Print(logger.Messagef(ERROR, format, msg...))
	}
}

func (logger *EasyLogger) Warnf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(WARN) {
		log.Print(logger.Messagef(WARN, format, msg...))
	}
}

func (logger *EasyLogger) Infof(format string, msg ...interface{}) {
	if logger.isLevelEnabled(INFO) {
		log.Print(logger.Messagef(INFO, format, msg...))
	}
}

func (logger *EasyLogger) Debugf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(DEBUG) {
		log.Print(logger.Messagef(DEBUG, format, msg...))
	}
}

func (logger *EasyLogger) Tracef(format string, msg ...interface{}) {
	if logger.isLevelEnabled(TRACE) {
		log.Print(logger.Messagef(TRACE, format, msg...))
	}
}

func (logger *EasyLogger) Flag() int {
	return logger.flag
}

func (logger *EasyLogger) SetFlag(flag int) LeveledLogger {
	stdFlags := flag & (Ldate | Ltime | Lmicroseconds | Llongfile | Lshortfile | LUTC)
	easylogFlags := flag & (Llevel)
	logger.Lock()
	logger.l.SetFlags(stdFlags)
	logger.flag = easylogFlags
	logger.Unlock()
	return logger
}

func (logger *EasyLogger) Writer() io.Writer {
	return logger.l.Writer()
}

func (logger *EasyLogger) SetWriter(w io.Writer) LeveledLogger {
	logger.l.SetOutput(w)
	return logger
}

func (logger *EasyLogger) Level() uint8 {
	return logger.level
}

func (logger *EasyLogger) SetLevel(lv string) LeveledLogger {
	logger.Lock()
	logger.level = lvAtoi[strings.ToUpper(lv)]
	logger.Unlock()
	return logger
}

func (logger *EasyLogger) isLevelEnabled(lv uint8) bool {
	return logger.level <= lv
}
