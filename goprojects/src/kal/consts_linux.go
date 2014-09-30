// +build linux

package main

import (
	"log"
	"log/syslog"
	"os"
	"strings"
)

const (
	dbDSNro = "file:/var/local/cc.db?cache=shared&mode=ro"
	dbDSNrw = "file:/var/local/cc.db?cache=shared&mode=rw"
	dbPath  = "/mnt/us/documents/mybooks/"
	fsPath  = "/mnt/us/documents/mybooks/"
)

var locale string = "en_US"

func init() {
	// See if we could make sense of current locale
	l := os.Getenv("LANG")
	if len(l) > 0 {
		l := strings.Split(l, ".")[0]
		if len(l) > 0 {
			locale = l
		}
	}
}

func prepareLog() {
	if !debug {
		lw, err := syslog.New(syslog.LOG_NOTICE, "KAL")
		if err == nil {
			log.SetOutput(lw)
		}
	}
}
