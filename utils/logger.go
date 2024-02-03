package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

// Console logger 
type ConsoleLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

// Create a new Console logger 
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warn:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		error: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

// Write DEBUG log to console
func (l *ConsoleLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Write INFO log to console
func (l *ConsoleLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Write WARN log to console
func (l *ConsoleLogger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

// Write ERROR log to console
func (l *ConsoleLogger) Error(v ...interface{}) {
	l.error.Println(v...)
}

// Logger for Log files
type FileLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

// Creates a new file logger 
func NewFileLogger(filename string) *FileLogger {
	logfile, err := os.OpenFile(os.Getenv("LOGFILE"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	fsize, err := strconv.Atoi(os.Getenv("LOGFILESIZE"))
    if err != nil {
        panic(err)
    }
	b, err := strconv.Atoi(os.Getenv("LOGFILEBACKUPS"))
    if err != nil {
        panic(err)
    }
	ldays, err := strconv.Atoi(os.Getenv("LOGFILEDAYS"))
    if err != nil {
        panic(err)
    }

	c := &lumberjack.Logger{
		Filename:   os.Getenv("LOGFILE"),
		MaxSize:    fsize,
		MaxBackups: b,
		MaxAge:     ldays,
	}
	dl := log.New(logfile, "DEBUG: ", log.LstdFlags)
	dl.SetOutput(c)

	il := log.New(logfile, "INFO: ", log.LstdFlags)
	il.SetOutput(c)

	wl := log.New(logfile, "WARN: ", log.LstdFlags)
	wl.SetOutput(c)

	el := log.New(logfile, "ERROR: ", log.LstdFlags)
	el.SetOutput(c)

	return &FileLogger{
		debug: dl,
		info:  il,
		warn:  wl,
		error: el,
	}
}

// Write DEBUG log to file
func (l *FileLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Write INFO log to file
func (l *FileLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Write WARN log to file
func (l *FileLogger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

// Write ERROR log to file
func (l *FileLogger) Error(v ...interface{}) {
	l.error.Println(v...)
}

var Clog *ConsoleLogger
var Flog *FileLogger