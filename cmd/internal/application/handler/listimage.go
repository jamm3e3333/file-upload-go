package handler

import (
	"context"
	"jamm3e3333/file-upload/cmd/internal/application/dto"
	"jamm3e3333/file-upload/cmd/internal/application/port"
)

type ListImage struct {
	lf port.ListImage
}

func NewListImage(lf port.ListImage) *ListImage {
	return &ListImage{
		lf: lf,
	}
}

func (h *ListImage) Handle(ctx context.Context) ([]*dto.ListImage, error) {
	return h.lf.Execute(ctx)
}
