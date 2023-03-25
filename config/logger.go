package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writter io.Writer
}

func ThrowNewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writter: writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Formatted Enabled Logs
func (l *Logger) DebugF(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) InfoF(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) WarnF(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) ErrorF(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
