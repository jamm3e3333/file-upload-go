package port

import (
	"jamm3e3333/file-upload/cmd/internal/application/dto"
)

type ConvertImage interface {
	ToWebp(img *dto.CreateImage) (*dto.CreateImage, error)
}
