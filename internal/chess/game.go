package chess

import (
	"github.com/mateinonedotcom/puzzle/pkg/puzzle"
	notnil "github.com/notnil/chess"
)

type Game struct {
	game *notnil.Game
}

func NewGame(game *notnil.Game) *Game {
	return &Game{
		game: game,
	}
}

func (g *Game) ToCheckmatePuzzle() *puzzle.CheckmatePuzzle {
	secondToLastPosition := secondToLastPosition(g.game)
	checkmateMoves := checkmateMoves(secondToLastPosition)
	if len(checkmateMoves) == 0 {
		return nil
	}

	squareMap := secondToLastPosition.Board().SquareMap()
	turn := secondToLastPosition.Turn()
	validMoves := validMoves(secondToLastPosition)

	return &puzzle.CheckmatePuzzle{
		PieceSquares:   toPieceSquares(squareMap),
		Turn:           toColor(turn),
		ValidMoves:     toMoves(validMoves),
		CheckmateMoves: toMoves(checkmateMoves),
	}
}

func secondToLastPosition(game *notnil.Game) *notnil.Position {
	history := game.MoveHistory()
	return history[len(history)-1].PrePosition
}

func validMoves(position *notnil.Position) []*notnil.Move {
	return position.ValidMoves()
}

func checkmateMoves(position *notnil.Position) []*notnil.Move {
	checkmates := []*notnil.Move{}
	for _, move := range position.ValidMoves() {
		position := position.Update(move)
		if position.Status() == notnil.Checkmate {
			checkmates = append(checkmates, move)
		}
	}

	return checkmates
}
