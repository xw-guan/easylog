package easylog

import (
	"os"
	"path"
	"sync"
	"time"
)

// Wrap a leveled logger and provide better log file management
type FileLogger struct {
	sync.Mutex
	LeveledLogger
	dir     string
	file    string
	daily   bool
	logFile *os.File
}

func NewFileLogger(l LeveledLogger) (fl *FileLogger) {
	fl = new(FileLogger)
	fl.LeveledLogger = l
	return
}

func (logger *FileLogger) SetOutputFile(dir string, file string, daily bool) (err error) {
	logger.Lock()
	defer logger.Unlock()
	if logger.dir == dir && logger.file == file && logger.daily == daily {
		return // nothing changes
	}
	logger.dir = dir
	logger.file = file
	logger.daily = daily

	//file += ".log"
	if daily {
		file = time.Now().Format("060102") + "_" + file
	}
	f, err := os.OpenFile(path.Join(dir, file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer silentlyCloseFile(logger.logFile)
	logger.logFile = f
	logger.SetWriter(f)

	Debugf("Set output file to %s\n", f.Name())
	return
}

func SetOutputFile(dir string, file string, daily bool) (err error) {
	fl, ok := stdLogger.(*FileLogger)
	if ok {
		return fl.SetOutputFile(dir, file, daily)
	}
	fl = NewFileLogger(stdLogger)
	if err = fl.SetOutputFile(dir, file, daily); err == nil {
		stdLogger = fl
	}
	return
}

func silentlyCloseFile(f *os.File) {
	if f != nil {
		Tracef("Closing file %s\n", f.Name())
		if err := f.Close(); err != nil {
			Errorf("Unable to close %s, it may cause a memory leak\n", f.Name())
		}
	}
}
