package easylog

import (
	"io"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type EasyConf struct {
	sync.Mutex
	fileConf
	l     *log.Logger // delegate
	level uint8
	flag  int
}

func (conf *EasyConf) Flag() int {
	return conf.flag
}

func (conf *EasyConf) SetFlag(flag int) {
	stdFlags := flag & (Ldate | Ltime | Lmicroseconds | Llongfile | Lshortfile | LUTC)
	easylogFlags := flag & (Llevel)
	conf.Lock()
	conf.l.SetFlags(stdFlags)
	conf.flag = easylogFlags
	conf.Unlock()
}

func (conf *EasyConf) Writer() io.Writer {
	return conf.l.Writer()
}

func (conf *EasyConf) SetWriter(w io.Writer) {
	conf.l.SetOutput(w)
}

func (conf *EasyConf) Level() uint8 {
	return conf.level
}

func (conf *EasyConf) SetLevel(lv string) {
	conf.Lock()
	conf.level = lvAtoi[strings.ToUpper(lv)]
	conf.Unlock()
}

func (conf *EasyConf) isLevelEnabled(lv uint8) bool {
	return conf.level >= lv
}

func (conf *EasyConf) SetOutputFile(filePath string, daily bool) (err error) {
	if filePath == "" {
		filePath = os.Stderr.Name() // default output
	}
	conf.Lock()
	defer conf.Unlock()
	if filePath == conf.path && conf.daily == daily {
		return // nothing changes
	}
	f, old, err := conf.setFile(filePath, daily)
	if err != nil {
		return err
	}
	conf.setNewFileAndCloseOld(f, old)
	return
}

func (conf *EasyConf) delegate() *log.Logger {
	if err := conf.changeDailyFileIfNeed(); err != nil {
		Error("Unable to change output filename:", err)
	}
	return conf.l
}

func (conf *EasyConf) changeDailyFileIfNeed() (err error) {
	if !conf.daily {
		return // daily file disabled
	}
	filePath := conf.LogFilePath()
	if filePath == conf.logFile.Name() {
		return // not need to change
	}
	f, err := OpenFile(filePath)
	if err != nil {
		return err // remain the old logFile
	}
	conf.setNewFileAndCloseOld(f, conf.logFile)
	return nil
}

func (conf *EasyConf) setNewFileAndCloseOld(f, old *os.File) {
	conf.logFile = f
	conf.SetWriter(f)
	Trace("Set output filePath to", f.Name())
	CloseFileSilently(old) // close the old logFile
}

type fileConf struct {
	dir      string
	filename string
	path     string
	logFile  *os.File
	daily    bool
}

func (conf *fileConf) LogFilePath() string {
	if conf.path == os.Stderr.Name() || conf.path == os.Stdout.Name() || !conf.daily {
		return conf.path
	}
	return path.Join(conf.dir, conf.dailyOutputFileName())
}

func (conf *fileConf) dailyOutputFileName() string {
	return time.Now().Format("060102") + "_" + conf.filename
}

func (conf *fileConf) setFile(filePath string, daily bool) (f *os.File, old *os.File, err error) {
	old = conf.logFile
	f, err = OpenFile(filePath)
	if err != nil {
		return nil, old, err
	}
	conf.logFile = f
	conf.path = filePath
	conf.daily = daily
	if daily {
		conf.dir, conf.filename = path.Split(filePath)
	}
	return
}

func OpenFile(path string) (*os.File, error) {
	switch path {
	case os.Stderr.Name():
		return os.Stderr, nil
	case os.Stdout.Name():
		return os.Stdout, nil
	default:
		return os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
}

func CloseFileSilently(f *os.File) {
	if f != nil && f.Name() != os.Stderr.Name() && f.Name() != os.Stdout.Name() {
		Trace("Closing filename", f.Name())
		if err := f.Close(); err != nil {
			Errorf("Unable to close %s, it may cause a memory leak\n", f.Name())
		}
	}
}

func NewStdConfig() EasyConf {
	fconf := fileConf{}
	_, _, _ = fconf.setFile(os.Stderr.Name(), false)
	return EasyConf{
		l: log.New(os.Stderr, "", log.LstdFlags),
		fileConf: fconf,
		level: WARN,
		flag:  Llevel,
	}
}
