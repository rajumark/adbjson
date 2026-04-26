package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Logger is a structured logger
type Logger struct {
	level  LogLevel
	output *log.Logger
}

var (
	// Global logger instance
	globalLogger *Logger
)

// Init initializes the global logger
func Init(level LogLevel) {
	globalLogger = &Logger{
		level:  level,
		output: log.New(os.Stderr, "", 0),
	}
}

// Get returns the global logger instance
func Get() *Logger {
	if globalLogger == nil {
		Init(INFO) // Default to INFO level
	}
	return globalLogger
}

// SetLevel changes the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// logEntry represents a structured log entry
type logEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Context   map[string]interface{} `json:"context,omitempty"`
}

// shouldLog determines if a message should be logged based on level
func (l *Logger) shouldLog(level LogLevel) bool {
	return level >= l.level
}

// levelToString converts LogLevel to string
func levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	default:
		return "unknown"
	}
}

// log is the internal logging method
func (l *Logger) log(level LogLevel, message string, context map[string]interface{}) {
	if !l.shouldLog(level) {
		return
	}

	entry := logEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     levelToString(level),
		Message:   message,
		Context:   context,
	}

	jsonBytes, err := json.Marshal(entry)
	if err != nil {
		// Fallback to simple log if JSON marshaling fails
		l.output.Printf("[%s] %s", levelToString(level), message)
		return
	}

	l.output.Println(string(jsonBytes))
}

// Debug logs a debug message
func (l *Logger) Debug(message string, context map[string]interface{}) {
	l.log(DEBUG, message, context)
}

// Info logs an info message
func (l *Logger) Info(message string, context map[string]interface{}) {
	l.log(INFO, message, context)
}

// Warn logs a warning message
func (l *Logger) Warn(message string, context map[string]interface{}) {
	l.log(WARN, message, context)
}

// Error logs an error message
func (l *Logger) Error(message string, context map[string]interface{}) {
	l.log(ERROR, message, context)
}

// Debugf logs a debug message with formatted string
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...), nil)
}

// Infof logs an info message with formatted string
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...), nil)
}

// Warnf logs a warning message with formatted string
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...), nil)
}

// Errorf logs an error message with formatted string
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...), nil)
}
