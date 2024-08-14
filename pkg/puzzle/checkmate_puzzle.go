package puzzle

import (
	"github.com/mateinonedotcom/puzzle/pkg/chess"
)

type CheckmatePuzzle struct {
	PieceSquares   []chess.PieceSquare `json:"pieceSquares"`
	Turn           chess.Color         `json:"turn"`
	ValidMoves     []chess.Move        `json:"validMoves"`
	CheckmateMoves []chess.Move        `json:"checkmateMoves"`
}
