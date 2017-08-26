package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	x := 1
	log.SetLevel(log.DebugLevel)
	l := log.WithFields(log.Fields{"x": x})
	l.Debug("hello")
	time.Sleep(1 * time.Second)
	l.Debug("hello")
}
