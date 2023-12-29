package operation

import (
	"context"
	"io"
	"jamm3e3333/file-upload/cmd/internal/application/dto"
	"jamm3e3333/file-upload/cmd/pkg/gcs"
	"time"
)

type UploadImage struct {
	sw                       gcs.StorageWriter
	createFileTimeoutMinutes time.Duration
}

func NewUploadImage(sw gcs.StorageWriter, t time.Duration) *UploadImage {
	return &UploadImage{sw: sw, createFileTimeoutMinutes: t}
}

func (o *UploadImage) Execute(ctx context.Context, filePath string, dto *dto.CreateImage) error {
	var cancel func()
	ctx, cancel = context.WithTimeout(ctx, time.Minute*o.createFileTimeoutMinutes)
	defer cancel()

	w := o.sw.NewWriter(ctx, filePath, dto.Name)

	// unused written bytes
	_, uploadErr := io.Copy(w, dto.Image)
	if uploadErr != nil {
		return uploadErr
	}

	err := w.Close()
	if err != nil {
		return err
	}

	return nil
}
