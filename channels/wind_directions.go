package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func main() {
	textChannel := make(chan string)
	metarChannel := make(chan []string)
	windsChannel := make(chan []string)
	resultChannel := make(chan [8]int)
	// Change to array, each metar report is a separate item in the array
	go parseToArray(textChannel, metarChannel)
	// Extract wind direction
	go extractWindDirection(metarChannel, windsChannel)
	// Assign to N. NE, E, SE, S, SW, W, NW
	go mineWindDistribution(windsChannel, resultChannel)

	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	startTime := time.Now()

	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}
		text := string(data)
		textChannel <- text
	}
	close(textChannel)
	results := <-resultChannel
	elapsed := time.Since(startTime)
	fmt.Printf("%v\n", results)
	fmt.Printf("Processing took %s", elapsed)
}
