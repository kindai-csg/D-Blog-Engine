package database

type FileHandler interface {
    Read(string) (string, error)
    Write(string, string) error
}
