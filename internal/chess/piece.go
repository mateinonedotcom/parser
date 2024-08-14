package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toPiece(piece external.Piece) internal.Piece {
	return internal.Piece{
		Type:  toPieceType(piece.Type()),
		Color: toColor(piece.Color()),
	}
}
