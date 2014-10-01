// +build linux

package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"strings"
)

const (
	dsnPattern = "file:%s?cache=shared&mode=%s"
)

var (
	dbDSNro string
	dbDSNrw string
	dbPath  string
	fsPath  string
	locale  = "en_US"
)

func init() {
	// See if we could make sense of current locale
	l := os.Getenv("LANG")
	if len(l) > 0 {
		l := strings.Split(l, ".")[0]
		if len(l) > 0 {
			locale = l
		}
	}

	dbDSNro = fmt.Sprintf(dsnPattern, "/var/local/cc.db", "ro")
	dbDSNrw = fmt.Sprintf(dsnPattern, "/var/local/cc.db", "rw")

	dbPath = filepath.Join("/mnt/us/documents", config.RelRoot)
	if !strings.HasSuffix(dbPath, string(os.PathSeparator)) {
		dbPath = dbPath + string(os.PathSeparator)
	}
	fsPath = dbPath
}

func prepareLog() {
	if !debug {
		lw, err := syslog.New(syslog.LOG_NOTICE, "KAL")
		if err == nil {
			log.SetOutput(lw)
		}
	}
}
