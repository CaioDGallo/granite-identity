package keymanager

type KeyLoader interface {
    LoadKey() ([]byte, error)
}
