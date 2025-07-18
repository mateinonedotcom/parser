package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toTupleSquareIsCheckmate(
	move *external.Move,
	position *external.Position,
) internal.TupleSquareIsCheckmate {
	newPosition := position.Update(move)
	isCheckmate := newPosition.Status() == external.Checkmate

	return internal.TupleSquareIsCheckmate{
		Square:      toSquare(move.S2()),
		IsCheckmate: isCheckmate,
	}
}

func toTupleSquareIsCheckmateArray(
	square external.Square,
	position *external.Position,
) []internal.TupleSquareIsCheckmate {
	tupleSquareIsCheckmateArray := make([]internal.TupleSquareIsCheckmate, 0, len(position.ValidMoves()))
	for _, move := range position.ValidMoves() {
		if move.S1() != square {
			continue
		}

		tupleSquareIsCheckmate := toTupleSquareIsCheckmate(move, position)
		tupleSquareIsCheckmateArray = append(tupleSquareIsCheckmateArray, tupleSquareIsCheckmate)
	}

	return tupleSquareIsCheckmateArray
}
