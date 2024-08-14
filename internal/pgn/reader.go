package pgn

import (
	"os"

	"github.com/mateinonedotcom/puzzle/internal/chess"
	notnil "github.com/notnil/chess"
)

type Reader struct {
	file *os.File
}

func NewReader(name string) (*Reader, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Reader{
		file: file,
	}, nil
}

func (r *Reader) ReadAll() <-chan *chess.Game {
	c := make(chan *chess.Game)
	go func() {
		scanner := notnil.NewScanner(r.file)
		for scanner.Scan() {
			game := scanner.Next()
			c <- chess.NewGame(game)
		}

		close(c)
	}()

	return c
}

func (r *Reader) Close() error {
	return r.file.Close()
}
