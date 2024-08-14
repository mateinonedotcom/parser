package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toSquare(square external.Square) internal.Square {
	return internal.Square{
		File: int(square.File()),
		Rank: int(square.Rank()),
	}
}
