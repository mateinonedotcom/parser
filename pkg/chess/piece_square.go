package chess

type PieceSquare struct {
	Piece  Piece  `json:"piece"`
	Square Square `json:"square"`
}
