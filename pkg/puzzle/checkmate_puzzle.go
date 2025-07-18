package puzzle

import (
	"github.com/mateinonedotcom/puzzle/pkg/chess"
)

type CheckmatePuzzle struct {
	Position []chess.TuplePieceSquareMoves `json:"position"`
}
