package infrastructure

import (
    "io/ioutil"
    "os"
)

type FileHandler struct {}

func NewFileHandler() *FileHandler {
    return &FileHandler{}
}

func (handler *FileHandler) Read(filepath string) (string, error) {
    bytes, err := ioutil.ReadFile(filepath)
    if err != nil {
            return "", err
    }
    return string(bytes), nil
}

func (handler *FileHandler) Write(filepath string, text string) error {
    fp, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
            return err
    }
    defer fp.Close()
    fp.WriteString(text)
    return nil
}
