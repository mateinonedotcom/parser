package main

import (
	"fmt"

	"github.com/mateinonedotcom/puzzle/internal/pgn"
	"github.com/mateinonedotcom/puzzle/internal/puzzle/json"
)

const fileToRead = "internal/pgn/testdata/lichess-2012-12-31.pgn"
const fileToWrite = "pkg/puzzle/json/testdata/checkmate-lichess-2012-12-31.jsonl"

func main() {
	fmt.Println("Hello, world.")

	reader, err := pgn.NewReader(fileToRead)
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	writer, err := json.NewWriter(fileToWrite)
	if err != nil {
		panic(err)
	}

	defer writer.Close()

	for game := range reader.ReadAll() {
		puzzle := game.ToCheckmatePuzzle()
		if puzzle == nil {
			continue
		}

		writer.Write(puzzle)
	}

}
