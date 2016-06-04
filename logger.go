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
	Silent  = iota
	Error   = iota
	Blocked = iota
	Warn    = iota
	Reject  = iota
	Listen  = iota
	Install = iota
	Init    = iota
	Request = iota
	Info    = iota
	Debug   = iota
	max     = iota
)

type Logger struct{ level int }

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

func (logger *Logger) Level() int { return logger.level }

func (logger *Logger) SetLevel(level int) error {
	if level > Debug || level < Silent {
		return fmt.Errorf("logger level must be >= %d and < %d", Silent, max)
	}
	logger.level = level
	return nil
}

func (logger *Logger) Error(format string, v ...interface{}) {
	out(logger.level, Error, format, v...)
}

func (logger *Logger) Blocked(format string, v ...interface{}) {
	out(logger.level, Blocked, format, v...)
}

func (logger *Logger) Warn(format string, v ...interface{}) {
	out(logger.level, Warn, format, v...)
}

func (logger *Logger) Reject(format string, v ...interface{}) {
	out(logger.level, Reject, format, v...)
}

func (logger *Logger) Listen(format string, v ...interface{}) {
	out(logger.level, Listen, format, v...)
}

func (logger *Logger) Install(format string, v ...interface{}) {
	out(logger.level, Install, format, v...)
}

func (logger *Logger) Init(format string, v ...interface{}) {
	out(logger.level, Init, format, v...)
}

func (logger *Logger) Request(format string, v ...interface{}) {
	out(logger.level, Request, format, v...)
}

func (logger *Logger) Info(format string, v ...interface{}) {
	out(logger.level, Info, format, v...)
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	out(logger.level, Debug, format, v...)
}

func MustError(format string, v ...interface{}) {
	out(max, Error, format, v...)
}

func MustBlocked(format string, v ...interface{}) {
	out(max, Blocked, format, v...)
}

func MustWarn(format string, v ...interface{}) {
	out(max, Warn, format, v...)
}

func MustReject(format string, v ...interface{}) {
	out(max, Reject, format, v...)
}

func MustListen(format string, v ...interface{}) {
	out(max, Listen, format, v...)
}

func MustInstall(format string, v ...interface{}) {
	out(max, Install, format, v...)
}

func MustInit(format string, v ...interface{}) {
	out(max, Init, format, v...)
}

func MustRequest(format string, v ...interface{}) {
	out(max, Request, format, v...)
}

func MustInfo(format string, v ...interface{}) {
	out(max, Info, format, v...)
}

func MustDebug(format string, v ...interface{}) {
	out(max, Debug, format, v...)
}

func New(level int) (*Logger, error) {
	logger := new(Logger)
	if err := logger.SetLevel(level); err != nil {
		return nil, err
	} else {
		return logger, nil
	}
}
