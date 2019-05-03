package easylog

import (
	"fmt"
	"strings"
)

type DecoratedLogger struct {
	EasyLogger
	prefixes []func(*strings.Builder)
	suffixes []func(*strings.Builder)
}

func NewDecoratedLogger(delegate EasyLogger) DecoratedLogger {
	logger := DecoratedLogger{
		EasyLogger: delegate,
	}
	return logger
}

func (logger *DecoratedLogger) WithPrefix(prefix ...func(*strings.Builder)) *DecoratedLogger {
	logger.prefixes = append(logger.prefixes, prefix...)
	logger.buildMessage = logger.buildMessageFunc()
	return logger
}

func (logger *DecoratedLogger) WithSuffix(suffix ...func(*strings.Builder)) *DecoratedLogger {
	logger.suffixes = append(logger.suffixes, suffix...)
	logger.buildMessage = logger.buildMessageFunc()
	return logger
}

func (logger *DecoratedLogger) WithFields(fields ...interface{}) *DecoratedLogger {
	logger.suffixes = append(logger.suffixes, func(sb *strings.Builder) {
		sb.WriteString("\n")
		for i, val := range fields {
			sb.WriteString(fmt.Sprint(val))
			if i%2 == 0 {
				sb.WriteString(" : ")
			} else {
				sb.WriteString("\n")
			}
		}
	})
	logger.buildMessage = logger.buildMessageFunc()
	return logger
}

func (logger *DecoratedLogger) buildMessageFunc() func(func(*strings.Builder), *strings.Builder) {
	return func(msg func(*strings.Builder), sb *strings.Builder) {
		for _, prefix := range logger.prefixes {
			if prefix != nil {
				prefix(sb)
			}
		}
		msg(sb)
		for _, suffix := range logger.suffixes {
			if suffix != nil {
				suffix(sb)
			}
		}
	}
}
func WithPrefix(prefix ...func(*strings.Builder)) *DecoratedLogger {
	logger := NewDecoratedLogger(stdLogger)
	return logger.WithPrefix(prefix...)
}

func WithSuffix(suffix ...func(*strings.Builder)) *DecoratedLogger {
	logger := NewDecoratedLogger(stdLogger)
	return logger.WithSuffix(suffix...)
}

func WithFields(fields ...interface{}) *DecoratedLogger {
	logger := NewDecoratedLogger(stdLogger)
	return logger.WithFields(fields...)
}