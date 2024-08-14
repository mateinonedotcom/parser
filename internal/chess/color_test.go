package chess

import (
	"testing"

	notnil "github.com/notnil/chess"

	"github.com/mateinonedotcom/puzzle/internal/test"
	"github.com/mateinonedotcom/puzzle/pkg/chess"
)

func TestToColorWhite(t *testing.T) {
	test.Equal(t, toColor(notnil.White), chess.White)
}

func TestToColorBlack(t *testing.T) {
	test.Equal(t, toColor(notnil.Black), chess.Black)
}
