package main

import "github.com/xw-guan/easylog"

func main() {
	easylog.WithFields(
		"id", 1,
		"username", "Emma",
	).Warn("User login")
	// 2019/05/03 13:19:18 [WARN] User login
	// id : 1
	// username : Emma
}
