package main

import "github.com/xw-guan/easylog"

func main() {
	filename := "easylog.log"
	if err := easylog.SetOutputFile(filename, true); err != nil {
		easylog.Error(err)
	}
	easylog.Warn("log to file ", filename)
	// on 20190420, log to file 190420_easylog.log
	// on 20190421, log to file 190421_easylog.log
}
