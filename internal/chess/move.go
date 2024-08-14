package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toMove(move *external.Move) internal.Move {
	return internal.Move{
		From: toSquare(move.S1()),
		To:   toSquare(move.S2()),
	}
}

func toMoves(moves []*external.Move) []internal.Move {
	internalMoves := make([]internal.Move, 0, len(moves))
	for _, move := range moves {
		internalMove := toMove(move)
		internalMoves = append(internalMoves, internalMove)
	}

	return internalMoves
}
