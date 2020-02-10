package lvm

import (
	stdlog "log"
	"os"
)

type logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

var log logger = stdlog.New(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile)

func SetLogger(l logger) {
	log = l
}
