package handler

import (
	"io"
	"os/exec"
)

type Handler interface{}

type handler struct {
	Command
	WorkDir string
}

func NewHandler(cmd Command, dir string) Handler {
	return &handler{
		Command: cmd,
		WorkDir: dir,
	}
}

type Command interface {
	ExecCommand(cmd *exec.Cmd, stdin io.Reader) (string, error)
}

type command struct{}

func NewCommand() Command {
	return &command{}
}

func (h command) ExecCommand(cmd *exec.Cmd, stdin io.Reader) (string, error) {
	cmd.Stdin = stdin
	out, err := cmd.CombinedOutput()
	return string(out), err
}
