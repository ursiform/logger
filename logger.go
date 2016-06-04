// Copyright 2016 Afshin Darian. All rights reserved.
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

// Package logger provides a simple logging interface that supports multiple
// levels of logging. It goes up to 11.
package logger

import (
	"fmt"
	"log"

	"github.com/mgutz/ansi"
)

const (
	// No log output at all.
	Silent = iota
	// Only errors are logged.
	Error = iota
	// Blocking calls and lower are logged.
	Blocked = iota
	// Warnings and lower are logged.
	Warn = iota
	// Rejections (e.g., in a firewall) and lower are logged.
	Reject = iota
	// Listeners and lower are logged.
	Listen = iota
	// Install notifications and lower are logged.
	Install = iota
	// Initialization notifications and lower are logged.
	Init = iota
	// Incoming requests and lower are logged.
	Request = iota
	// Info output and lower are logged.
	Info = iota
	// All log output is shown.
	Debug = iota
	max   = iota
)

// Logger only outputs log messages that are equal or lower than its log level.
type Logger struct{ level int }

// LogLevel translates human-readable log level names to their int values.
var LogLevel = map[string]int{
	"silent":  Silent,
	"error":   Error,
	"blocked": Blocked,
	"warn":    Warn,
	"reject":  Reject,
	"listen":  Listen,
	"install": Install,
	"init":    Init,
	"request": Request,
	"info":    Info,
	"debug":   Debug}

var prefixes = [...]string{
	"",
	ansi.Color("[***error***]", "black:red"),
	ansi.Color("[**blocked**]", "255+b:165"),
	ansi.Color("[**warning**]", "red:yellow+h"),
	ansi.Color("[ rejection ]", "125+b:208"),
	ansi.Color("[ listening ]", "black:cyan+h"),
	/*      */ "[ installed ]",
	/*      */ "[initialized]",
	/*      */ "[  request  ]",
	/*      */ "[information]",
	/*      */ "[   debug   ]"}

func out(loggerLevel, messageLevel int, format string, v ...interface{}) {
	if loggerLevel == Silent || messageLevel == Silent {
		return
	}
	message := fmt.Sprintf(format, v...)
	if messageLevel > Silent && messageLevel <= loggerLevel {
		log.Printf("%s %s", prefixes[messageLevel], message)
	}
}

// Level returns a logger instance's log level.
func (logger *Logger) Level() int { return logger.level }

// SetLevel sets a logger instance's log level.
func (logger *Logger) SetLevel(level int) error {
	if level > Debug || level < Silent {
		return fmt.Errorf("logger level must be >= %d and < %d", Silent, max)
	}
	logger.level = level
	return nil
}

// Error logs an error.
func (logger *Logger) Error(format string, v ...interface{}) {
	out(logger.level, Error, format, v...)
}

// Blocked logs a blocked message.
func (logger *Logger) Blocked(format string, v ...interface{}) {
	out(logger.level, Blocked, format, v...)
}

// Warn logs a warning.
func (logger *Logger) Warn(format string, v ...interface{}) {
	out(logger.level, Warn, format, v...)
}

// Reject logs a rejection.
func (logger *Logger) Reject(format string, v ...interface{}) {
	out(logger.level, Reject, format, v...)
}

// Listen logs a listener message.
func (logger *Logger) Listen(format string, v ...interface{}) {
	out(logger.level, Listen, format, v...)
}

// Install logs an install message.
func (logger *Logger) Install(format string, v ...interface{}) {
	out(logger.level, Install, format, v...)
}

// Init logs an initialization message.
func (logger *Logger) Init(format string, v ...interface{}) {
	out(logger.level, Init, format, v...)
}

// Request logs a request.
func (logger *Logger) Request(format string, v ...interface{}) {
	out(logger.level, Request, format, v...)
}

// Info logs an info message.
func (logger *Logger) Info(format string, v ...interface{}) {
	out(logger.level, Info, format, v...)
}

// Debug logs a debug message.
func (logger *Logger) Debug(format string, v ...interface{}) {
	out(logger.level, Debug, format, v...)
}

// MustError logs an error.
func MustError(format string, v ...interface{}) {
	out(max, Error, format, v...)
}

// MustBlocked logs a blocked message.
func MustBlocked(format string, v ...interface{}) {
	out(max, Blocked, format, v...)
}

// MustWarn logs a warning.
func MustWarn(format string, v ...interface{}) {
	out(max, Warn, format, v...)
}

// MustReject logs a rejection.
func MustReject(format string, v ...interface{}) {
	out(max, Reject, format, v...)
}

// MustListen logs a listener message.
func MustListen(format string, v ...interface{}) {
	out(max, Listen, format, v...)
}

// MustInstall logs an install message.
func MustInstall(format string, v ...interface{}) {
	out(max, Install, format, v...)
}

// MustInit logs an initialization message.
func MustInit(format string, v ...interface{}) {
	out(max, Init, format, v...)
}

// MustRequest logs a request.
func MustRequest(format string, v ...interface{}) {
	out(max, Request, format, v...)
}

// MustInfo logs an info message.
func MustInfo(format string, v ...interface{}) {
	out(max, Info, format, v...)
}

// MustDebug logs a debug message.
func MustDebug(format string, v ...interface{}) {
	out(max, Debug, format, v...)
}

// New instantiates a Logger and sets its level. If the level is invalid, it
// returns an error.
func New(level int) (*Logger, error) {
	logger := new(Logger)
	if err := logger.SetLevel(level); err != nil {
		return nil, err
	} else {
		return logger, nil
	}
}
