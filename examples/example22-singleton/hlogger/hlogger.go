package hlogger

import (
	"log"
	"os"
	"sync"
)

type HydraLogger struct {
	filename string
	*log.Logger
}

var hlogger *HydraLogger
var once sync.Once

func createLogger(fname string) *HydraLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &HydraLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra: ", log.Lshortfile),
	}
}

func GetInstance() *HydraLogger {
	once.Do(func() {
		hlogger = createLogger("hydralogger.log")
	})
	return hlogger
}
