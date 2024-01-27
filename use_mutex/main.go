package main

import (
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func stingy() {
	log.Info().Msg("Stingy Start")
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	log.Info().Msg("Stingy Done")
}

func spendy() {
	log.Info().Msg("Spendy Start")
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
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
