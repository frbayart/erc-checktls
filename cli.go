// cli.go

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fDebug         bool
	fType          string
	fOutput        string
	fSiteName      string
	fIgnoreImirhil bool
	fVerbose       bool
	fReallyVerbose bool
)

const (
	cliUsage = `%s version %s
Usage: %s [-hvVI] [-t text|csv] [-o file] file[.json]

`
)

// Usage string override.
var Usage = func() {
	fmt.Fprintf(os.Stderr, cliUsage, MyName, MyVersion, MyName)
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&fOutput, "o", "-", "Save into file (default stdout)")
	flag.StringVar(&fType, "t", "text", "Type of report")
	flag.StringVar(&fSiteName, "S", "", "Display that site")
	flag.BoolVar(&fIgnoreImirhil, "I", false, "Do not fetch tls.imirhil.fr grade")
	flag.BoolVar(&fDebug, "D", false, "Debug mode")
	flag.BoolVar(&fVerbose, "v", false, "Verbose mode")
	flag.BoolVar(&fReallyVerbose, "V", false, "More verbose mode")
}
