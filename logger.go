package franky

import (
	"fmt"
)

// Logger is a basic interface for use in many projects
type Logger interface {
	Printf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})
}

type noopLogger struct{}

// NoopLogger implements the logger interface
// but will not do anything
var NoopLogger = &noopLogger{}

func (l *noopLogger) Printf(string, ...interface{}) {}
func (l *noopLogger) Infof(string, ...interface{})  {}
func (l *noopLogger) Errorf(string, ...interface{}) {}
func (l *noopLogger) Debugf(string, ...interface{}) {}
func (l *noopLogger) Warnf(string, ...interface{})  {}

type defaultLogger struct{}

// DefaultLogger is a default way of interacting with
// logging system
var DefaultLogger = &defaultLogger{}

func (l *defaultLogger) Printf(s string, args ...interface{}) {
	fmt.Printf(
		fmt.Sprintf("%s\n", s),
		args...,
	)
}
func (l *defaultLogger) Infof(s string, args ...interface{}) {
	fmt.Printf(
		fmt.Sprintf("[INFO] %s\n", s),
		args...,
	)
}
func (l *defaultLogger) Errorf(s string, args ...interface{}) {
	fmt.Printf(
		fmt.Sprintf("[ERROR] %s\n", s),
		args...,
	)
}

func (l *defaultLogger) Debugf(s string, args ...interface{}) {
	fmt.Printf(
		fmt.Sprintf("[DEBUG] %s\n", s),
		args...,
	)
}

func (l *defaultLogger) Warnf(s string, args ...interface{}) {
	fmt.Printf(
		fmt.Sprintf("[WARN] %s\n", s),
		args...,
	)
}
