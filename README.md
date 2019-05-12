# EasyLog

EasyLog is a simple leveled logger for Go. The goal of this project is to provide an easily-used logging tool for small Go applications.

## Install

```sh
go get github.com/xw-guan/easylog
```

or

```sh
git clone git@github.com:xw-guan/easylog.git
```

## Usage

### Log levels

```go
FATAL   1
PANIC   2
ERROR   3
WARN    4
INFO    5
DEBUG   6
TRACE   7
```

### Log and formatting log

The way to use easylog is almost the same as `fmt.Println()` and `fmt.Printf()`

```go
easylog.Info("the way to use log is the same as", "fmt.Println()")
// 2019/04/20 16:01:56 [INFO] the way to use log is the same as fmt.Println()
easylog.Infof("the way to use formatting log is the same as %s", "fmt.Printf()")
// 2019/04/20 16:01:56 [INFO] the way to use formatting log is the same as fmt.Printf()
```

`easylog.Panic()` or `easylog.Panicf()` will call `panic()`

`easylog.Fatal()` or `easylog.Fatalf()` will call `os.Exit(1)`

### Set log level

Logs with levels lower than the setting one are ignored. The default log level is **WARN**

```go
easylog.SetLevel("info")

easylog.Debug("ignored") // ignored
easylog.Info("logged")
```

### Log to file

Giving easylog the path of file, all the logs are going to be writen to it. The file or the directory should be writable. The path can be either relative or absolute. The default output is **stderr**.

```go
filepath := "logs/easylog.log"
err = easylog.SetOutputFile(filepath, false)
if err != nil {
    easylog.Error(err) // cannot open the file
}
easylog.Warn("log to file ", filepath)
// you will see something like 2019/05/03 10:55:57 [WARN] log to file logs/easylog.log in logs/easylog.log
```

You can also set output to an existing writer (NOT RECOMMENDED, bug)

```go
filename := "easylog.log"
file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
    fmt.Println(err)
}
easylog.SetWritter(file)
easylog.Warn("log to file ", filename)
```

### One day one log file

Passing true to the second param of func `easylog.SetOutputFile(filepath string, daily bool)` writes logs to a new file on a new day. yyMMdd_ is added to the file as prefix.

```go
filename := "logs/easylog.log"
if err := easylog.SetOutputFile(filename, true); err != nil {
    easylog.Error(err)
}
easylog.Warn("log to file ", filename)
// on 20190420, log to file logs/190420_easylog.log
// on 20190421, log to file logs/190421_easylog.log
```

### Log with key-values

Odd param as key and the next param as its value

```go
easylog.WithFields(
    "id", 1,
    "username", "Emma",
).Warn("User login")
// 2019/05/03 13:19:18 [WARN] User login
// id : 1
// username : Emma
```

### Log with prefixes and suffixes

```go
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
easylog.WithPrefix(host, pid).WithSuffix(home).Warn("decorate with prefixes and suffixes")
```

## Benchmarks

TODO