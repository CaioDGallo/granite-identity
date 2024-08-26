package keymanager

import (
	"sync"
)

type keyManager struct {
	err  error
	key  []byte
	once sync.Once
}

var instance *keyManager

func LoadKey(loader KeyLoader) ([]byte, error) {
	if instance == nil {
		instance = &keyManager{}
	}

	instance.once.Do(func() {
		instance.key, instance.err = loader.LoadKey()
	})

	return instance.key, instance.err
}

func GetKey() []byte {
	if instance == nil {
		return nil
	}

	return instance.key
}
