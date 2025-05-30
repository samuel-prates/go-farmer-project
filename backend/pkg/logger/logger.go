// pkg/logger/logger.go
package logger

import (
	"fmt"
	"log"
	"os"
)

// LogLevel represents the severity of the log message
type LogLevel int

const (
	// InfoLevel is used for general information
	InfoLevel LogLevel = iota
	// WarnLevel is used for warnings that might cause issues
	WarnLevel
	// ErrorLevel is used for errors that need attention
	ErrorLevel
	// FatalLevel is used for critical errors that require the application to exit
	FatalLevel
)

// Logger is a simple logger with level support
type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

// New creates a new Logger instance
func New() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime),
		fatalLogger: log.New(os.Stderr, "[FATAL] ", log.Ldate|log.Ltime),
	}
}

// formatMessage formats the log message with timestamp and additional context
func formatMessage(message string, args ...interface{}) string {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	return message
}

// Info logs an info level message
func (l *Logger) Info(message string, args ...interface{}) {
	l.infoLogger.Println(formatMessage(message, args...))
}

// Warn logs a warning level message
func (l *Logger) Warn(message string, args ...interface{}) {
	l.warnLogger.Println(formatMessage(message, args...))
}

// Error logs an error level message
func (l *Logger) Error(message string, args ...interface{}) {
	l.errorLogger.Println(formatMessage(message, args...))
}

// Fatal logs a fatal level message and exits the application
func (l *Logger) Fatal(message string, args ...interface{}) {
	l.fatalLogger.Println(formatMessage(message, args...))
	os.Exit(1)
}

// Default logger instance
var defaultLogger = New()

// GetDefaultLogger returns the default logger instance
func GetDefaultLogger() *Logger {
	return defaultLogger
}

// Info logs an info level message using the default logger
func Info(message string, args ...interface{}) {
	defaultLogger.Info(message, args...)
}

// Warn logs a warning level message using the default logger
func Warn(message string, args ...interface{}) {
	defaultLogger.Warn(message, args...)
}

// Error logs an error level message using the default logger
func Error(message string, args ...interface{}) {
	defaultLogger.Error(message, args...)
}

// Fatal logs a fatal level message and exits the application using the default logger
func Fatal(message string, args ...interface{}) {
	defaultLogger.Fatal(message, args...)
}