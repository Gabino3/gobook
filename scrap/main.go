package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	playingWithLogs()
}


func playingWithLogs() {
	x := 1
	log.SetLevel(log.DebugLevel)
	l := log.WithFields(log.Fields{"x": x})
	l.Debug("hello")
	time.Sleep(1 * time.Second)
	y := 2
	l = l.WithField("y", y)
	l.Debug("hello")
}