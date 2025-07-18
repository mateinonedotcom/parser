package chess

type TuplePieceSquareMoves struct {
	Piece  Piece                    `json:"piece"`
	Square Square                   `json:"square"`
	Moves  []TupleSquareIsCheckmate `json:"moves"`
}
