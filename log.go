package main

import (
	"fmt"
	"os"
	"time"
)

type LogLevel uint8

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

func (ll LogLevel) String() string {
	switch ll {
	case LogLevelDebug:
		return "[DEBUG]"
	case LogLevelInfo:
		return "[INFO]"
	case LogLevelWarn:
		return "[WARN]"
	case LogLevelError:
		return "[ERROR]"
	case LogLevelFatal:
		return "[FATAL]"
	default:
		return "???"
	}
}

func logf(level LogLevel, format string, a ...any) {
	log(level, fmt.Sprintf(format, a...))
}

func log(level LogLevel, a any) {
	fmt.Printf("% -7s [%s] %v\n", level, time.Now().Format(time.RFC1123), a)
	if level == LogLevelFatal {
		os.Exit(1)
	}
}
