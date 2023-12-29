package controller

import (
	"context"
	"fmt"
	"jamm3e3333/file-upload/cmd/internal/application/command"
	httperror "jamm3e3333/file-upload/cmd/internal/ui/error"
	"jamm3e3333/file-upload/cmd/internal/ui/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandleCreateImage interface {
	Handle(ctx context.Context, cm *command.CreateImage) error
}

type CreateImage struct {
	ge                *gin.Engine
	handleCreateImage HandleCreateImage
}

func NewCreateImage(h HandleCreateImage) *CreateImage {
	return &CreateImage{handleCreateImage: h}
}

func (c *CreateImage) Register(ge *gin.Engine) {
	ge.POST(
		"/image",
		middleware.CheckMaxBodySize,
		middleware.RetrieveImage,
		middleware.CheckImgContentType,
		c.create,
	)
}

func (c *CreateImage) create(ctx *gin.Context) {
	var fp middleware.ImageUploadParam

	imgCtx := ctx.MustGet(middleware.ImageCtxKey)
	fp = imgCtx.(middleware.ImageUploadParam)

	f, err := fp.Image.Open()
	defer func() {
		err := f.Close()
		fmt.Printf("couldn't close the uploaded file: %s", err.Error())
	}()

	if f == nil {
		ctx.JSON(http.StatusBadRequest, httperror.NewBadRequestError(fmt.Errorf("uploaded file is empty")))
	}

	ct := fp.Image.Header.Get(middleware.ContentType)

	cm := &command.CreateImage{
		Image: f,
		Name:  fp.Image.Filename,
		Type:  ct,
	}

	err = c.handleCreateImage.Handle(ctx, cm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httperror.NewInternalError(err))
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}
