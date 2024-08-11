package main

import (
	"fmt"
	"os"

	"github.com/notnil/chess"
)

func main() {
	fmt.Println("Hello, world.")

	f, err := os.Open("pgn/lichess/2012/12/31/lichess-2012-12-31.pgn")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := chess.NewScanner(f)
	for scanner.Scan() {
		game := scanner.Next()

		fmt.Println(game)
	}
}
