package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/vito/invaders"
)

func main() {
	var seed = flag.Int64("seed", time.Now().UnixNano(), "initial seed")

	r := rand.New(rand.NewSource(*seed))

	fmt.Println("seed:", *seed)

	invader := invaders.Invader{}

	for {
		invader.Set(r)

		fmt.Print(invader)

		time.Sleep(time.Second)
	}
}
