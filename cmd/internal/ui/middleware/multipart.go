package middleware

import (
	httperror "jamm3e3333/file-upload/cmd/internal/ui/error"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type imageContentType string

const (
	ImagePng  imageContentType = "image/png"
	ImageJpeg imageContentType = "image/jpeg"
	ImageGif  imageContentType = "image/gif"
)

var imgContentTypeToContentType = map[imageContentType]imageContentType{
	ImagePng:  ImagePng,
	ImageJpeg: ImageJpeg,
	ImageGif:  ImageGif,
}

const (
	ContentType = "Content-Type"
	ImageCtxKey = "image"
)

type ImageUploadParam struct {
	Image multipart.FileHeader `form:"image" binding:"required"`
}

func RetrieveImage(ctx *gin.Context) {
	var fp ImageUploadParam
	err := ctx.MustBindWith(&fp, binding.FormMultipart)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httperror.NewBadRequestError(err))
		return
	}

	ctx.Set(ImageCtxKey, fp)
	ctx.Next()
}

func CheckImgContentType(ctx *gin.Context) {
	var fp ImageUploadParam
	ctxImg := ctx.MustGet(ImageCtxKey)

	fp = ctxImg.(ImageUploadParam)

	fileCt := fp.Image.Header.Get(ContentType)
	if _, exists := imgContentTypeToContentType[imageContentType(fileCt)]; !exists {
		ctx.AbortWithStatusJSON(http.StatusUnsupportedMediaType, httperror.NewUnsupportedImgContentTypeError(fileCt))
		return
	}

	ctx.Next()
}
