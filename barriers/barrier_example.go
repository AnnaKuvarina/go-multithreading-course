package main

import (
	"github.com/rs/zerolog/log"
	"time"
)

func waitOnBarrier(name string, timeToSleep int, barrier *Barrier) {
	for {
		log.Info().Msgf("%s running", name)
		time.Sleep(time.Duration(timeToSleep) * time.Second)
		log.Info().Msgf("%s is waiting on barrier", name)
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go waitOnBarrier("red", 4, barrier)
	go waitOnBarrier("blue", 10, barrier)
	time.Sleep(time.Duration(100) * time.Second)
}
