package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toPieceType(pieceType external.PieceType) internal.PieceType {
	switch pieceType {
	case external.King:
		return internal.King
	case external.Queen:
		return internal.Queen
	case external.Rook:
		return internal.Rook
	case external.Bishop:
		return internal.Bishop
	case external.Knight:
		return internal.Knight
	case external.Pawn:
		return internal.Pawn
	default:
		panic("Invalid piece type: " + pieceType.String())
	}
}
