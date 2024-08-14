package json

import (
	"encoding/json"
	"os"

	"github.com/mateinonedotcom/puzzle/pkg/puzzle"
)

type Reader struct {
	file    *os.File
	decoder *json.Decoder
}

func NewReader(name string) (*Reader, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)

	return &Reader{
		file:    file,
		decoder: decoder,
	}, nil
}

func (r *Reader) ReadAll() <-chan *puzzle.CheckmatePuzzle {
	c := make(chan *puzzle.CheckmatePuzzle)
	go func() {
		for r.decoder.More() {
			var puzzle puzzle.CheckmatePuzzle
			if err := r.decoder.Decode(&puzzle); err != nil {
				panic(err)
			}

			c <- &puzzle
		}

		close(c)
	}()

	return c
}

func (r *Reader) Close() error {
	return r.file.Close()
}
