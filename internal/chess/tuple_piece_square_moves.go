package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toTuplePieceSquareMoves(piece external.Piece, square external.Square, position *external.Position) internal.TuplePieceSquareMoves {
	return internal.TuplePieceSquareMoves{
		Piece:  toPiece(piece),
		Square: toSquare(square),
		Moves:  toTupleSquareIsCheckmateArray(square, position),
	}
}

func toTuplePieceSquareMovesArray(
	mapSquarePiece map[external.Square]external.Piece,
	position *external.Position,
) []internal.TuplePieceSquareMoves {
	tuplePieceSquareMovesArray := make([]internal.TuplePieceSquareMoves, 0, len(mapSquarePiece))
	for square, piece := range mapSquarePiece {
		tuplePieceSquareMoves := toTuplePieceSquareMoves(piece, square, position)

		tuplePieceSquareMovesArray = append(tuplePieceSquareMovesArray, tuplePieceSquareMoves)
	}

	return tuplePieceSquareMovesArray
}
