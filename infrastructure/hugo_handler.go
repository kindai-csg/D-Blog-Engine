package infrastructure

import (
    "os/exec"
    "fmt"
)

type HugoHandler struct {
    path string
}

func NewHugoHandler(path string) *HugoHandler {
    hugoHandler := HugoHandler{
        path,
    }
    return &hugoHandler
}

func (handler *HugoHandler) Deploy() error {
    cmd := exec.Command("hugo")
    cmd.Dir = handler.path
    err := cmd.Run()
    if err != nil {
        return err
    }
    cmd = exec.Command("git", "add", ".")
    cmd.Dir = handler.path + "/public"
    fmt.Println(cmd.Dir)
    err = cmd.Run()
    if err != nil {
        return err
    }
    fmt.Println("check")
    cmd = exec.Command("git", "commit", "-m", ":rocket: update article")
    cmd.Dir = handler.path + "/public"
    err = cmd.Run()
    if err != nil {
        return err
    }
    cmd = exec.Command("git", "push", "origin", "master")
    cmd.Dir = handler.path + "/public"
    err = cmd.Run()
    if err != nil {
        return err
    }
    cmd = exec.Command("git", "add", ".")
    cmd.Dir = handler.path
    err = cmd.Run()
    if err != nil {
        return err
    }
    cmd = exec.Command("git", "commit", "-m", ":rocket: update article")
    cmd.Dir = handler.path
    err = cmd.Run()
    if err != nil {
        return err
    }
    cmd = exec.Command("git", "push", "origin", "master")
    cmd.Dir = handler.path
    return cmd.Run()
}

func (handler *HugoHandler) Update() error {
    cmd := exec.Command("git", "pull", "origin", "master")
    cmd.Dir = handler.path
    return cmd.Run()
}
