package chess

import (
	"testing"

	notnil "github.com/notnil/chess"

	"github.com/mateinonedotcom/puzzle/internal/test"
	"github.com/mateinonedotcom/puzzle/pkg/chess"
)

func TestToPieceTypeKing(t *testing.T) {
	test.Equal(t, toPieceType(notnil.King), chess.King)
}

func TestToPieceTypeQueen(t *testing.T) {
	test.Equal(t, toPieceType(notnil.Queen), chess.Queen)
}

func TestToPieceTypeRook(t *testing.T) {
	test.Equal(t, toPieceType(notnil.Rook), chess.Rook)
}

func TestToPieceTypeBishop(t *testing.T) {
	test.Equal(t, toPieceType(notnil.Bishop), chess.Bishop)
}

func TestToPieceTypeKnight(t *testing.T) {
	test.Equal(t, toPieceType(notnil.Knight), chess.Knight)
}

func TestToPieceTypePawn(t *testing.T) {
	test.Equal(t, toPieceType(notnil.Pawn), chess.Pawn)
}
