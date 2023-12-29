package handler

import (
	"context"
	"jamm3e3333/file-upload/cmd/internal/application/command"
	"jamm3e3333/file-upload/cmd/internal/application/dto"
	"jamm3e3333/file-upload/cmd/internal/application/port"
)

type CreateImageHandler struct {
	upi  port.UploadImage
	covi port.ConvertImage
}

func NewCreateImageHandler(upi port.UploadImage, covi port.ConvertImage) *CreateImageHandler {
	return &CreateImageHandler{upi: upi, covi: covi}
}

// Handle TODO: Should add domain? (file entity with method injection)
func (h *CreateImageHandler) Handle(ctx context.Context, cm *command.CreateImage) error {
	d := &dto.CreateImage{
		Image: cm.Image,
		Name:  cm.Name,
		Type:  cm.Type,
	}

	convImg, err := h.covi.ToWebp(d)
	if err != nil {
		return err
	}

	// TODO: resolve the path by id (user)
	imgErr := h.upi.Execute(ctx, "img/user1", convImg)
	if imgErr != nil {
		return imgErr
	}

	return nil
}
