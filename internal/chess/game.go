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
	if !isWhiteTurn(secondToLastPosition) || !hasCheckmate(secondToLastPosition) {
		return nil
	}

	mapSquarePiece := secondToLastPosition.Board().SquareMap()

	return &puzzle.CheckmatePuzzle{
		Position: toTuplePieceSquareMovesArray(mapSquarePiece, secondToLastPosition),
	}
}

func secondToLastPosition(game *notnil.Game) *notnil.Position {
	history := game.MoveHistory()
	return history[len(history)-1].PrePosition
}

func isWhiteTurn(position *notnil.Position) bool {
	return position.Turn() == notnil.White
}

func hasCheckmate(position *notnil.Position) bool {
	for _, move := range position.ValidMoves() {
		position := position.Update(move)
		if position.Status() == notnil.Checkmate {
			return true
		}
	}

	return false
}
