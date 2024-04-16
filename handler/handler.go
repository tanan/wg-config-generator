package handler

import (
	"io"
	"log/slog"
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
	if err != nil {
		slog.Error("failed to exec command", slog.String("command", cmd.String()))
		return "", err
	}
	return string(out), err
}
