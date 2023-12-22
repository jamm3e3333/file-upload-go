package controller

import (
	"context"
	"fmt"
	"jamm3e3333/file-upload/application/command"
	httperror "jamm3e3333/file-upload/ui/error"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type fileUploadParam struct {
	File multipart.FileHeader `form:"file" binding:"required"`
}

type HandleCreateFile func(ctx context.Context, cm *command.CreateFile) error

type FileUpload struct {
	ge               *gin.Engine
	handleCreateFile HandleCreateFile
}

func NewFileUpload(h HandleCreateFile) *FileUpload {

	return &FileUpload{handleCreateFile: h}
}

func (c *FileUpload) Register(ge *gin.Engine) {
	ge.POST("/file", c.upload)
}

func (c *FileUpload) upload(ctx *gin.Context) {
	var fp fileUploadParam

	err := ctx.MustBindWith(&fp, binding.FormMultipart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httperror.NewBadRequestError(err))
		return
	}

	f, err := fp.File.Open()

	// HandleCreateFile TODO: check the error in the command
	if f == nil {
		ctx.JSON(http.StatusBadRequest, httperror.NewBadRequestError(fmt.Errorf("uploaded file is empty")))
	}
	cm := &command.CreateFile{File: f}

	err = c.handleCreateFile(ctx, cm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httperror.NewInternalError(err))
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}
