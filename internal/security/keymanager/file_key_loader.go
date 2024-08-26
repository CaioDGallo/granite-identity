package keymanager

import (
	"os"
)

type FileKeyLoader struct {
	Path string
}

func (l *FileKeyLoader) LoadKey() ([]byte, error) {
	key, err := os.ReadFile(l.Path)
	if err != nil {
		return nil, err
	}
	return key, nil
}
