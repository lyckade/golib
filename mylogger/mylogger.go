/*
Package mylogger is a simple abstraction for the logger
*/
package mylogger

import (
	"log"
	"os"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	CRITICAL
	FATAL
)

//NewFileLogger returns a simple Logger which creates a
//Logger which writes into a log file
func NewFileLogger(fileName string, prefix string) *log.Logger {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		f, _ = os.Create(fileName)
		//panic(fmt.Sprintf("NewFileLogger can not be created, because file could not be opened: %s", err))
	}
	//defer f.Close()
	logger := log.New(f, prefix, log.Lshortfile|log.Ldate|log.Ltime)
	return logger
}
