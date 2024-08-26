package keymanager

import (
	"fmt"
	"os"
)

type EnvKeyLoader struct {
	EnvVar string
}

func (l *EnvKeyLoader) LoadKey() ([]byte, error) {
	key := os.Getenv(l.EnvVar)
	if key == "" {
		return nil, fmt.Errorf("environment variable %s is not set", l.EnvVar)
	}
	return []byte(key), nil
}
