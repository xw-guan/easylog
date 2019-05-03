package main

import (
	"fmt"
	"github.com/xw-guan/easylog"
	"os"
	"path"
)

func main() {
	// you can set output to an existing writer
	filename := "easylog.log"
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	easylog.SetWritter(file)
	easylog.Warn("log to file ", filename)

	// or call SetOutputFile()
	filename = "1" + filename
	err = easylog.SetOutputFile(filename, false)
	if err != nil {
		easylog.Error(err)
	}
	easylog.Warn("log to file ", filename)

	// you can use absolute path or relative path
	home, _ := os.UserHomeDir()
	err = easylog.SetOutputFile(path.Join(home, filename), false)
	//err = easylog.SetOutputFile(path.Join("/var/log/easylog", filename), false)
	if err != nil {
		easylog.Error(err)
	}
	easylog.Warn("log to file of absolute path ", filename)

	// see also one-day-one-log-file.go
}
