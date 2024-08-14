package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toPieceSquare(square external.Square, piece external.Piece) internal.PieceSquare {
	return internal.PieceSquare{
		Piece:  toPiece(piece),
		Square: toSquare(square),
	}
}

func toPieceSquares(squarePieceMap map[external.Square]external.Piece) []internal.PieceSquare {
	pieceSquares := make([]internal.PieceSquare, 0, len(squarePieceMap))
	for square, piece := range squarePieceMap {
		pieceSquare := toPieceSquare(square, piece)
		pieceSquares = append(pieceSquares, pieceSquare)
	}

	return pieceSquares
}
