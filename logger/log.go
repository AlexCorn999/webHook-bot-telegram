package logger

import (
	"log"
	"os"
)

var (
	outFile, _ = os.OpenFile("/logs/webhook.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile    = log.New(outFile, "", 0)
)

func ForError(er error) {
	if er != nil {
		LogFile.Fatalln(er)
	}
}
