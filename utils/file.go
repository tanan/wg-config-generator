package utils

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

func CreateFile(path string, mode fs.FileMode) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(mode))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create file : %s", path))
		return nil, err
	}

	err = os.Chmod(path, mode)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to change file permission : %s", path))
		return nil, err
	}
	return f, nil
}

func Makedir(path string, perm fs.FileMode) error {
	dir := filepath.Dir(path)
	if info, err := os.Stat(dir); os.IsExist(err) {
		if info.Mode() == perm {
			return nil
		}
		if err := os.Chmod(dir, perm); err != nil {
			slog.Error(fmt.Sprintf("failed to change directory permission : %s", path))
			return err
		}
	}
	return os.Mkdir(dir, perm)
}
