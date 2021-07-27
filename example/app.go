package main

import (
	"time"

	"github.com/bcowtech/gosfmt"
)

func main() {
	rng := gosfmt.NewSFMT()
	rng.Seed(time.Now().UnixNano())
}
