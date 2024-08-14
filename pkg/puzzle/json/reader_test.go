package json

import (
	"testing"

	"github.com/mateinonedotcom/puzzle/pkg/puzzle"
)

func TestReader(t *testing.T) {
	r, err := NewReader("./testdata/checkmate-lichess-2012-12-31.jsonl")
	if err != nil {
		t.Fatal(err)
	}

	defer r.Close()

	puzzles := []*puzzle.CheckmatePuzzle{}
	for puzzle := range r.ReadAll() {
		puzzles = append(puzzles, puzzle)
	}

	if len(puzzles) != 47 {
		t.Fatalf("expected 1 game, got %d", len(puzzles))
	}
}
