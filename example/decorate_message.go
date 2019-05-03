package main

import (
	"github.com/xw-guan/easylog"
	"os"
	"strconv"
	"strings"
)

func main() {
	host := func(builder *strings.Builder) {
		h, _ := os.Hostname()
		builder.WriteString("Host: ")
		builder.WriteString(h)
		builder.WriteString(" ")
	}
	pid := func(builder *strings.Builder) {
		builder.WriteString("Pid: ")
		builder.WriteString(strconv.Itoa(os.Getpid()))
		builder.WriteString(" ")
	}
	home := func(builder *strings.Builder) {
		dir, _ := os.UserHomeDir()
		builder.WriteString(" ")
		builder.WriteString(dir)
	}
	easylog.WithPrefix(host, pid).WithSuffix(home).Warn("decorate")
}
