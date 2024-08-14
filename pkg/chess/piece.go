package chess

type Piece struct {
	Type  PieceType `json:"type"`
	Color Color     `json:"color"`
}
