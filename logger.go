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
	// Silent logs no output at all.
	Silent = iota

	// Error logs only errors.
	Error = iota

	// Blocked logs blocking calls and lower.
	Blocked = iota

	// Warn logs warnings and lower.
	Warn = iota

	// Reject logs rejections (e.g., in a firewall) and lower.
	Reject = iota

	// Listen logs listeners and lower.
	Listen = iota

	// Install logs install notifications and lower.
	Install = iota

	// Init logs initialization notifications and lower.
	Init = iota

	// Request logs incoming requests and lower.
	Request = iota

	// Info logs info output and lower.
	Info = iota

	// Verbose logs all verbose output and lower.
	Verbose = iota

	// Debug logs all log output.
	Debug = iota

	max = iota
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
	"verbose": Verbose,
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
	/*      */ "[  verbose  ]",
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

// Level returns a logger instance's log level if logger level allows.
func (l *Logger) Level() int { return l.level }

// SetLevel sets a logger instance's log level if logger level allows.
func (l *Logger) SetLevel(level int) error {
	if level > Debug || level < Silent {
		return fmt.Errorf("logger level must be >= %d and < %d", Silent, max)
	}
	l.level = level
	return nil
}

// Error logs an error if logger level allows.
func (l *Logger) Error(format string, v ...interface{}) {
	out(l.level, Error, format, v...)
}

// Blocked logs a blocked message if logger level allows.
func (l *Logger) Blocked(format string, v ...interface{}) {
	out(l.level, Blocked, format, v...)
}

// Warn logs a warning if logger level allows.
func (l *Logger) Warn(format string, v ...interface{}) {
	out(l.level, Warn, format, v...)
}

// Reject logs a rejection if logger level allows.
func (l *Logger) Reject(format string, v ...interface{}) {
	out(l.level, Reject, format, v...)
}

// Listen logs a listener message if logger level allows.
func (l *Logger) Listen(format string, v ...interface{}) {
	out(l.level, Listen, format, v...)
}

// Install logs an install message if logger level allows.
func (l *Logger) Install(format string, v ...interface{}) {
	out(l.level, Install, format, v...)
}

// Init logs an initialization message if logger level allows.
func (l *Logger) Init(format string, v ...interface{}) {
	out(l.level, Init, format, v...)
}

// Request logs a request if logger level allows.
func (l *Logger) Request(format string, v ...interface{}) {
	out(l.level, Request, format, v...)
}

// Info logs an info message if logger level allows.
func (l *Logger) Info(format string, v ...interface{}) {
	out(l.level, Info, format, v...)
}

// Verbose logs a verbose message if logger level allows.
func (l *Logger) Verbose(format string, v ...interface{}) {
	out(l.level, Verbose, format, v...)
}

// Debug logs a debug message if logger level allows.
func (l *Logger) Debug(format string, v ...interface{}) {
	out(l.level, Debug, format, v...)
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

// MustVerbose logs a verbose message.
func MustVerbose(format string, v ...interface{}) {
	out(max, Verbose, format, v...)
}

// MustDebug logs a debug message.
func MustDebug(format string, v ...interface{}) {
	out(max, Debug, format, v...)
}

// New instantiates a Logger and sets its level. If the level is invalid, it
// returns an error.
func New(level int) (*Logger, error) {
	l := new(Logger)
	if err := l.SetLevel(level); err != nil {
		return nil, err
	}
	return l, nil
}
