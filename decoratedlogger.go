package easylog

import (
	"fmt"
	"strings"
)

type DecoratedLogger struct {
	LeveledLogger
	prefixes *MessageContainer
	suffixes *MessageContainer
}

type MessageContainer []func(*strings.Builder)

func (container *MessageContainer) Add(f func(*strings.Builder)) {
	*container = append(*container, f)
}

func (container *MessageContainer) BuildMessage(sb *strings.Builder) {
	for _, f := range *container {
		f(sb)
	}
}

// Override LeveledLogger.BuildMessage
func (logger *DecoratedLogger) BuildMessage(level uint8, msg string, sb *strings.Builder) *strings.Builder {
	logger.prefixes.BuildMessage(sb)
	logger.LeveledLogger.BuildMessage(level, msg, sb)
	logger.suffixes.BuildMessage(sb)
	return sb
}

func (logger *DecoratedLogger) AddPrefix(f func(*strings.Builder)) *DecoratedLogger {
	logger.prefixes.Add(f)
	return logger
}

func (logger *DecoratedLogger) AddSuffix(f func(*strings.Builder)) *DecoratedLogger {
	logger.suffixes.Add(f)
	return logger
}

func decorateWithFields(delegate LeveledLogger, fields ...interface{}) *DecoratedLogger {
	logger := DecoratedLogger{LeveledLogger: delegate}
	logger.AddSuffix(func(sb *strings.Builder) {
		for i, val := range fields {
			sb.WriteString(fmt.Sprint(val))
			if i%2 == 0 {
				sb.WriteString(" : ")
			} else {
				sb.WriteString("\n")
			}
		}
	})
	return &logger
}

func WithFields(fields ...interface{}) *DecoratedLogger {
	return decorateWithFields(stdLogger, fields...)
}

func AddPrefix(f func(*strings.Builder)) *DecoratedLogger {
	logger := DecoratedLogger{LeveledLogger: stdLogger}
	return logger.AddPrefix(f)
}

func AddSuffix(f func(*strings.Builder)) *DecoratedLogger {
	logger := DecoratedLogger{LeveledLogger: stdLogger}
	return logger.AddSuffix(f)
}
