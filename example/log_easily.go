package main

import (
	"errors"
	"github.com/xw-guan/easylog"
)

func main() {
	DoSomething := func() error {
		return errors.New("error occurred while do something")
	}

	if err := DoSomething(); err != nil {
		easylog.Error(err) // 2019/04/20 16:01:56 [ERROR] error occurred while do something
	}

	// the default level is warn
	easylog.Info("you cannot see this message")
	easylog.SetLevel("info")
	easylog.Info("now you see the message") // 2019/04/20 16:01:56 [INFO] now you see the message

	easylog.Infof("the way to use this formatting log is the same as %s", "fmt.Printf()")

	//easylog.Panic("easylog.Panic will call panic()")
	easylog.Fatal("fatal will call os.Exit(1)") // 2019/04/20 16:01:56 [FATAL] fatal will call os.Exit(1)

}