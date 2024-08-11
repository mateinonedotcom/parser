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
		if game.Method() != chess.Checkmate {
			continue
		}

		moveHistory := game.MoveHistory()
		moveLast := moveHistory[len(moveHistory)-1]

		fmt.Println(moveLast.PrePosition.Board().Draw())
		fmt.Println("Turn: ", moveLast.PrePosition.Turn())
		fmt.Println("Moves: ", moveLast.PrePosition.ValidMoves())
		fmt.Println("")

		checkmates := []*chess.Move{}
		for _, move := range moveLast.PrePosition.ValidMoves() {
			position := moveLast.PrePosition.Update(move)
			if position.Status() == chess.Checkmate {
				checkmates = append(checkmates, move)
			}
		}

		fmt.Println("Checkmates: ", checkmates)

		break
	}
}
