package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// Packages
	"github.com/fatih/color"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Logger handles logging at different levels based on verbose flag
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	infoColor   *color.Color
	debugColor  *color.Color
	errorColor  *color.Color
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewLogger creates a new logger instance
func NewLogger(verbose bool) *Logger {
	flags := log.Lmsgprefix
	logger := &Logger{
		errorLogger: log.New(os.Stderr, "ERROR: ", flags),
		infoLogger:  log.New(os.Stdout, "", flags),
		infoColor:   color.New(color.Bold),
		debugColor:  color.New(color.FgCyan),
		errorColor:  color.New(color.FgRed),
	}
	if verbose {
		logger.debugLogger = log.New(os.Stdout, "DEBUG: ", flags)
	}
	return logger
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Debug logs debug messages (only when verbose)
func (l *Logger) Debug(v ...interface{}) {
	if l.debugLogger != nil {
		l.debugColor.Fprint(os.Stdout, v...)
		fmt.Fprintln(os.Stdout)
	}
}

// Info logs informational messages (always shown)
func (l *Logger) Info(v ...interface{}) {
	l.infoColor.Fprint(os.Stdout, v...)
	fmt.Fprintln(os.Stdout)
}

// Infof logs formatted informational messages (only when verbose)
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoColor.Fprintf(os.Stdout, format, v...)
}

// Error logs error messages (always shown)
func (l *Logger) Error(v ...interface{}) {
	l.errorColor.Fprint(os.Stderr, v...)
	fmt.Fprintln(os.Stderr)
}

// Errorf logs formatted error messages (always shown)
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorColor.Fprintf(os.Stderr, format, v...)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// logging middleware
func logging(next http.Handler, logger *Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("%s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
