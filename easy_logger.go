package easylog

import (
	"fmt"
	"strings"
)

type EasyLogger struct {
	EasyConf
	buildMessage func(func(*strings.Builder), *strings.Builder)
}

func OriginalMessage(msg ...interface{}) func(sb *strings.Builder) {
	return func(sb *strings.Builder) {
		sb.WriteString(fmt.Sprint(msg...))
	}
}

func OriginalMessagef(format string, msg ...interface{}) func(sb *strings.Builder) {
	return func(sb *strings.Builder) {
		sb.WriteString(fmt.Sprintf(format, msg...))
	}
}

func NewEasyLogger() EasyLogger {
	return EasyLogger{
		EasyConf: NewStdConfig(),
		buildMessage: func(msg func(*strings.Builder), sb *strings.Builder) {
			msg(sb)
		},
	}
}

func (logger *EasyLogger) levelMessage(level uint8, sb *strings.Builder) {
	if logger.flag|Llevel != 0 && lvItoa[level] != "" {
		sb.WriteString("[")
		sb.WriteString(lvItoa[level])
		sb.WriteString("] ")
	}
}

func (logger *EasyLogger) Message(level uint8, msg ...interface{}) *strings.Builder {
	var sb strings.Builder
	logger.levelMessage(level, &sb)
	logger.buildMessage(OriginalMessage(msg...), &sb)
	return &sb
}

func (logger *EasyLogger) Messagef(level uint8, format string, msg ...interface{}) *strings.Builder {
	var sb strings.Builder
	logger.levelMessage(level, &sb)
	logger.buildMessage(OriginalMessagef(format, msg...), &sb)
	return &sb
}

func (logger *EasyLogger) Fatal(msg ...interface{}) {
	if logger.isLevelEnabled(FATAL) {
		logger.delegate().Fatal(logger.Message(FATAL, msg...))
	}
}

func (logger *EasyLogger) Panic(msg ...interface{}) {
	if logger.isLevelEnabled(PANIC) {
		logger.delegate().Panic(logger.Message(PANIC, msg...))
	}
}

func (logger *EasyLogger) Error(msg ...interface{}) {
	if logger.isLevelEnabled(ERROR) {
		logger.delegate().Print(logger.Message(ERROR, msg...))
	}
}

func (logger *EasyLogger) Warn(msg ...interface{}) {
	if logger.isLevelEnabled(WARN) {
		logger.delegate().Print(logger.Message(WARN, msg...))
	}
}

func (logger *EasyLogger) Info(msg ...interface{}) {
	if logger.isLevelEnabled(INFO) {
		logger.delegate().Print(logger.Message(INFO, msg...))
	}
}

func (logger *EasyLogger) Debug(msg ...interface{}) {
	if logger.isLevelEnabled(DEBUG) {
		logger.delegate().Print(logger.Message(DEBUG, msg...))
	}
}

func (logger *EasyLogger) Trace(msg ...interface{}) {
	if logger.isLevelEnabled(TRACE) {
		logger.delegate().Print(logger.Message(TRACE, msg...))
	}
}

func (logger *EasyLogger) Fatalf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(FATAL) {
		logger.delegate().Fatal(logger.Messagef(FATAL, format, msg...))
	}
}

func (logger *EasyLogger) Panicf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(PANIC) {
		logger.delegate().Panic(logger.Messagef(PANIC, format, msg...))
	}
}

func (logger *EasyLogger) Errorf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(ERROR) {
		logger.delegate().Print(logger.Messagef(ERROR, format, msg...))
	}
}

func (logger *EasyLogger) Warnf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(WARN) {
		logger.delegate().Print(logger.Messagef(WARN, format, msg...))
	}
}

func (logger *EasyLogger) Infof(format string, msg ...interface{}) {
	if logger.isLevelEnabled(INFO) {
		logger.delegate().Print(logger.Messagef(INFO, format, msg...))
	}
}

func (logger *EasyLogger) Debugf(format string, msg ...interface{}) {
	if logger.isLevelEnabled(DEBUG) {
		logger.delegate().Print(logger.Messagef(DEBUG, format, msg...))
	}
}

func (logger *EasyLogger) Tracef(format string, msg ...interface{}) {
	if logger.isLevelEnabled(TRACE) {
		logger.delegate().Print(logger.Messagef(TRACE, format, msg...))
	}
}