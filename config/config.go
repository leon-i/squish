package config

import (
	"fmt"
)

type optimizerConfig struct {
	Quality     uint
	Destination string
}

var SquishConfig = new(optimizerConfig)

func SetValues(quality uint, destination string) {
	if quality == 0 || !isValidQuality(quality) {
		SquishConfig.Quality = 75
	} else {
		SquishConfig.Quality = quality
	}

	if destination == "" || len(destination) >= 256 {
		SquishConfig.Destination = "squished"
	} else {
		SquishConfig.Destination = destination
	}

	fmt.Println("\nimage quality set to", SquishConfig.Quality)
	fmt.Println("\nimage(s) will be output to ./" + SquishConfig.Destination)
	fmt.Println("")
}

func isValidQuality(q uint) (valid bool) {
	if q < 0 || q > 100 {
		return false
	}

	return true
}