package main

import (
	"io"
	"log"
	"os"
)

// newLogger : ロガー生成
func newLogger() *log.Logger {
	var w io.Writer
	switch os.Getenv("ENV") {
	default:
		w = os.Stdout
	}
	return log.New(w, "[Akindo] ", log.LstdFlags)
}
