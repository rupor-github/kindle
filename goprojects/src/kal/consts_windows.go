// +build windows

package main

import (
	"fmt"
	"kal/Godeps/_workspace/src/bitbucket.org/kardianos/osext"
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
	exep, _ := osext.ExecutableFolder()
	p := filepath.Join(exep, "test\\cc.db")

	dbDSNro = fmt.Sprintf(dsnPattern, p, "ro")
	dbDSNrw = fmt.Sprintf(dsnPattern, p, "rw")

	fsPath = filepath.Join(exep, "test\\documents", config.RelRoot)
	if !strings.HasSuffix(fsPath, string(os.PathSeparator)) {
		fsPath = fsPath + string(os.PathSeparator)
	}

	dbPath = filepath.ToSlash(filepath.Join("/mnt/us/documents", config.RelRoot))
	if !strings.HasSuffix(dbPath, string(os.PathSeparator)) {
		dbPath = dbPath + string(os.PathSeparator)
	}
	dbPath = filepath.ToSlash(dbPath)
}

func prepareLog() {
}
