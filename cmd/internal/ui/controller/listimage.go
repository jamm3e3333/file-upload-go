package controller

import (
	"context"
	"jamm3e3333/file-upload/cmd/internal/application/dto"
	httperror "jamm3e3333/file-upload/cmd/internal/ui/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListImageHandler interface {
	Handle(ctx context.Context) ([]*dto.ListImage, error)
}

type ListImage struct {
	h ListImageHandler
}

type listImageData struct {
	Name string `json:"name" required:"true"`
	Url  string `json:"url" required:"true"`
}

type listImageResponse struct {
	Data []listImageData `json:"data"`
}

func NewImageFile(h ListImageHandler) *ListImage {
	return &ListImage{h: h}
}

func (c *ListImage) Register(ge *gin.Engine) {
	ge.GET("/image", c.list)
}

func (c *ListImage) list(ctx *gin.Context) {
	f, err := c.h.Handle(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httperror.NewInternalError(err))
		return
	}

	files := make([]listImageData, 0, len(f))
	for _, v := range f {
		file := listImageData{Name: v.Name, Url: v.SignedUrl}
		files = append(files, file)
	}

	ctx.JSON(http.StatusOK, &listImageResponse{Data: files})
}
