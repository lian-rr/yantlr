package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

const prefix = "Y∆NTLR"

type Logger struct {
	Logger *log.Logger
	OutFD  *os.File
}

func NewLogger(out io.Writer) *Logger {
	p := os.Getpid()

	return &Logger{
		Logger: log.New(out, fmt.Sprintf("(%d) %s ", p, prefix),
			log.Ldate|log.Lmicroseconds),
	}
}

func (L *Logger) Fatal(s string, args ...interface{}) { L.Logger.Fatalf("[FATAL] "+s, args...) }
func (L *Logger) Panic(s string, args ...interface{}) { L.Logger.Panicf("[FATAL] "+s, args...) }
func (L *Logger) Error(s string, args ...interface{}) {
	L.Logger.Printf("[ERROR] "+s, args...)
	debug.PrintStack()
}
func (L *Logger) Info(s string, args ...interface{}) { L.Logger.Printf("[INFO] "+s, args...) }
func (L *Logger) Access(r *http.Request) {
	L.Logger.Printf("[REQ] %s - %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
}
