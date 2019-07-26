package logger

import (
	"log"
	"os"
)

var (
	Log *log.Logger
)

const (
	VERSION = "0.13"
)

//https://tutorialedge.net/golang/the-go-init-function/
// very nice !
// will be called only once. on the first import of this package.
func init() {

	var file, err1 = os.Create("app.log")

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile :  ./app.log " + "VERSION :" + VERSION)
}
