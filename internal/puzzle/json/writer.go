package json

import (
	"encoding/json"
	"os"
)

type Writer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewWriter(name string) (*Writer, error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	encoder := json.NewEncoder(file)

	return &Writer{
		file:    file,
		encoder: encoder,
	}, nil
}

func (w *Writer) Write(a any) error {
	return w.encoder.Encode(a)
}

func (w *Writer) Close() error {
	return w.file.Close()
}
