package chess

import (
	internal "github.com/mateinonedotcom/puzzle/pkg/chess"

	external "github.com/notnil/chess"
)

func toColor(c external.Color) internal.Color {
	switch c {
	case external.White:
		return internal.White
	case external.Black:
		return internal.Black
	default:
		panic("Invalid color: " + c.Name())
	}
}
