package logger

import (
	"fmt"
	"log"

	"github.com/mgutz/ansi"
)

const (
	Silent = iota
	Error  = iota
	Warn   = iota
	Listen = iota
	Init   = iota
	Info   = iota
	Debug  = iota
	max    = iota
)

type Logger struct{ Level int }

var prefixes = [...]string{
	"",
	ansi.Color("[***error***]", "black+B:red"),
	ansi.Color("[**warning**]", "red:yellow+h"),
	ansi.Color("[ listening ]", "black:cyan+h"),
	/*      */ "[initialized]",
	/*      */ "[information]",
	/*      */ "[   debug   ]"}

func (logger *Logger) Log(level int, message string) {
	if logger.Level == Silent {
		return
	}
	// If an invalid Level has been set, output an error and the message as info.
	if logger.Level > Debug {
		log.Printf("%s logger.Level must be < %d", prefixes[Error], max)
		log.Printf("%s %s", prefixes[Info], message)
		return
	}
	if 0 <= level && level <= logger.Level {
		log.Printf("%s %s", prefixes[level], message)
	}
}

func New(level int) (*Logger, error) {
	if level > Debug {
		return nil, fmt.Errorf("logger.Level must be < %d", max)
	}
	return &Logger{Level: level}, nil
}
