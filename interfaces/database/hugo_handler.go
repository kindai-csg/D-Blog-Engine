package database

type HugoHandler interface {
    Deploy() error
    Update() error
}
