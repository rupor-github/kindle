// +build windows

package main

const (
	dbDSNro = "file:Q:/kindle/projects/goprojects/bin/test/cc.db?cache=shared&mode=ro"
	dbDSNrw = "file:Q:/kindle/projects/goprojects/bin/test/cc.db?cache=shared&mode=rw"
	dbPath  = "/mnt/us/documents/mybooks/"
	fsPath  = "Q:\\kindle\\projects\\goprojects\\bin\\test\\documents\\mybooks\\"
)

var locale string = "en_US"

func prepareLog() {
}
