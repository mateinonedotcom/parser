package pgn

import (
	"testing"

	"github.com/mateinonedotcom/puzzle/internal/chess"
)

func TestReader(t *testing.T) {
	r, err := NewReader("./testdata/lichess-2012-12-31.pgn")
	if err != nil {
		t.Fatal(err)
	}

	defer r.Close()

	games := []*chess.Game{}
	for game := range r.ReadAll() {
		games = append(games, game)
	}

	if len(games) != 135 {
		t.Fatalf("expected 1 game, got %d", len(games))
	}
}
