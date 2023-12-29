package port

import (
	"context"
	"jamm3e3333/file-upload/cmd/internal/application/dto"
)

type UploadImage interface {
	Execute(ctx context.Context, filePath string, dto *dto.CreateImage) error
}
