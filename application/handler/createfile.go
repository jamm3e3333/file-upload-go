package handler

import (
	"context"
	"fmt"
	"io"
	"jamm3e3333/file-upload/application/command"
	"os"
	"path/filepath"
)

type CreateFile struct {
	cm *command.CreateFile
}

// Handle TODO: Should add domain, maybe (file entity with method injection)
func Handle(ctx context.Context, cm *command.CreateFile) error {
	var f *os.File
	path, err := filepath.Abs("files/file.jpg")

	if err != nil {
		return err
	}
	f, err = os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println("cannot close the file")
		}
	}()

	_, err = io.Copy(f, cm.File)
	if err != nil {
		return err
	}

	return nil
}
