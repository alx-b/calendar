package logger

import (
	"log"
	"os"
)

// Globally create the logger so it can be import to other modules/files.
var Log = CreateLogger()

type Logger struct {
	File  *os.File
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Panic *log.Logger
}

// CreateLogger open Write-Only log file (create file if it does not exist).
// Append errors to the log.
// Returns Logger struct
func CreateLogger() Logger {
	file, err := os.OpenFile(
		"logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	return Logger{
		File:  file,
		Info:  log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile),
		Warn:  log.New(file, "WARN: ", log.LstdFlags|log.Lshortfile),
		Error: log.New(file, "ERRO: ", log.LstdFlags|log.Lshortfile),
		Panic: log.New(file, "PANI: ", log.LstdFlags|log.Lshortfile),
	}
}

// CloseFile calls the os.File Close function.
func (l *Logger) CloseFile() {
	if err := l.File.Close(); err != nil {
		log.Fatal(err)
	}
}

// SyncFile calls os.File Sync function.
func (l *Logger) SyncFile() {
	if err := l.File.Sync(); err != nil {
		log.Fatal(err)
	}
}
