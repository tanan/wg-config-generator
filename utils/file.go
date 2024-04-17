package utils

import (
	"fmt"
	"log/slog"
	"os"
)

func CreateFile(path string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create file : %s", path))
		return nil, err
	}
	return f, nil
}
