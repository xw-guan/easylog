package easylog

import (
	"fmt"
	"io"
	"log"
)

const (
	LvFatal uint8 = 0
	LvPanic uint8 = 1
	LvError uint8 = 2
	LvInfo  uint8 = 3
	LvDebug uint8 = 4
)

var level uint8 = 1

var lvMap = map[string]uint8{
	"fatal": LvFatal,
	"panic": LvPanic,
	"error": LvError,
	"info":  LvInfo,
	"debug": LvDebug,
}

func SetLogLevel(lv string) {
	level = lvMap[lv]
}

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func Debug(msg interface{}) {
	if IsLvEnabled(LvDebug) {
		log.Println(msg)
	}
}

func Debugf(format string, v ...interface{}) {
	if IsLvEnabled(LvDebug) {
		log.Printf(format, v)
	}
}

func DebugPrint(msg interface{}) {
	if IsLvEnabled(LvDebug) {
		fmt.Print(msg)
	}
}

func Info(msg interface{}) {
	if IsLvEnabled(LvInfo) {
		log.Println(msg)
	}
}

func Infof(format string, v ...interface{}) {
	if IsLvEnabled(LvInfo) {
		log.Printf(format, v)
	}
}

func InfoPrint(msg interface{}) {
	if IsLvEnabled(LvInfo) {
		fmt.Print(msg)
	}
}

func Error(msg interface{}) {
	if IsLvEnabled(LvError) {
		log.Println(msg)
	}
}

func Errorf(format string, v ...interface{}) {
	if IsLvEnabled(LvError) {
		log.Printf(format, v)
	}
}

func Panic(msg interface{}) {
	if IsLvEnabled(LvPanic) {
		log.Panicln(msg)
	}
}

func Panicf(format string, v ...interface{}) {
	if IsLvEnabled(LvPanic) {
		log.Printf(format, v)
	}
}

func Fatal(msg interface{}) {
	if IsLvEnabled(LvFatal) {
		log.Fatalln(msg)
	}
}

func Fatalf(format string, v ...interface{}) {
	if IsLvEnabled(LvFatal) {
		log.Printf(format, v)
	}
}

func IsLvEnabled(lv uint8) bool {
	return lv <= level
}
