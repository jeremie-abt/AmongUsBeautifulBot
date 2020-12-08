package log

import (
	"log"
	"os"
)

var (
	Debuglog = log.New(os.Stderr, "[DEBUG]", log.Ltime|log.Lshortfile)
	Infolog  = log.New(os.Stdout, "[INFO]", log.Ltime|log.Lshortfile)
	Warnlog  = log.New(os.Stderr, "[WARN]", log.Ltime|log.Lshortfile)
	Buildlog = log.New(os.Stderr, "[BUILD]", log.Ltime|log.Lshortfile)
	Errlog   = log.New(os.Stderr, "[ERROR]", log.Ltime|log.Lshortfile)
)
