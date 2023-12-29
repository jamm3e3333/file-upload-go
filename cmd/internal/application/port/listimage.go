package port

import (
	"context"
	"jamm3e3333/file-upload/cmd/internal/application/dto"
)

type ListImage interface {
	Execute(ctx context.Context) ([]*dto.ListImage, error)
}
