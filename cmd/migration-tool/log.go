/*
 * @Author: LinkLeong link@icewhale.org
 * @Date: 2022-08-30 22:15:30
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-08-30 22:15:47
 * @FilePath: /CasaOS/cmd/migration-tool/log.go
 * @Description:
 * @Website: https://www.casaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package main

import (
	"log"
	"os"
)

type Logger struct {
	DebugMode bool

	_debug *log.Logger
	_info  *log.Logger
	_error *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		DebugMode: false,
		_debug:    log.New(os.Stdout, "DEBUG: ", 0),
		_info:     log.New(os.Stdout, "", 0),
		_error:    log.New(os.Stderr, "ERROR: ", 0),
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.DebugMode {
		l._debug.Printf(format, v...)
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l._info.Printf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l._error.Printf(format, v...)
}
