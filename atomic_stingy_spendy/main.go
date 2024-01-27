package main

import (
	"github.com/rs/zerolog/log"
	"sync/atomic"
	"time"
)

var (
	money int32 = 100
)

func stingy() {
	log.Info().Msg("Stingy Start")
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, 10)
		time.Sleep(1 * time.Millisecond)
	}
	log.Info().Msg("Stingy Done")
}

func spendy() {
	log.Info().Msg("Spendy Start")
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, -10)
		time.Sleep(1 * time.Millisecond)
	}
	log.Info().Msg("Spendy Done")
}

func main() {
	go spendy()
	go stingy()
	time.Sleep(3000 * time.Millisecond)
	log.Info().Msgf("Result = %d", money)
}
