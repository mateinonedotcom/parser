package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toSquare(square external.Square) internal.Square {
	return internal.Square{
		File: square.File().String(),
		Rank: square.Rank().String(),
	}
}
