package util

import (
	"io"
	"log"
	"os"
)

var Log	*log.Logger

func init() {
	var logpath =  GetCurrentDirectory() + "/log/spr.log"

	file, err := os.OpenFile(logpath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	Log = log.New(io.MultiWriter(os.Stderr, file),
		"[DEBUG]", log.LstdFlags | log.Lshortfile)
}
