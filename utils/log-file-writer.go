package utils

import (
	"os"
	"path/filepath"
	"time"
)

const logDir = "/var/log/go-server"
const logFileFormat = "2006-01-02.log"

const logFileCheckInterval = time.Minute

type logFileWriter struct {
	file       *os.File
	lastCheck  time.Time
	checkEvery time.Duration
}

func NewLogFileWriter() *logFileWriter {
	return &logFileWriter{
		checkEvery: logFileCheckInterval,
	}
}

func (w *logFileWriter) OpenLogFile(logFilePath string) error {
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	w.file = file
	return nil
}

func (w *logFileWriter) getCorrectLogFile() error {
	logFileName := time.Now().Format(logFileFormat)
	logFilePath := filepath.Join(logDir, logFileName)
	if w.file == nil {
		w.OpenLogFile(logFilePath)
	} else if w.file.Name() != logFilePath {
		w.file.Close()
		w.OpenLogFile(logFilePath)
	}
	return nil
}

func (w *logFileWriter) Write(p []byte) (n int, err error) {
	// 检查是否需要切换日志文件
	if time.Since(w.lastCheck) > w.checkEvery {
		w.lastCheck = time.Now()
		if err := w.getCorrectLogFile(); err != nil {
			return 0, err
		}
	}
	// 写入日志数据
	return w.file.Write(p)
}
