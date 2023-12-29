package internal

import (
	"jamm3e3333/file-upload/cmd/app/config"
	"jamm3e3333/file-upload/cmd/internal/application/handler"
	"jamm3e3333/file-upload/cmd/internal/infrastructure/gcs/operation"
	"jamm3e3333/file-upload/cmd/internal/ui/controller"
	"jamm3e3333/file-upload/cmd/pkg/convertimage"
	"jamm3e3333/file-upload/cmd/pkg/gcs"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ge *gin.Engine, storageCfg *config.GcStorage) {
	// deps
	cs := gcs.NewCloudStorage(storageCfg)
	conv := convertimage.NewImgConverter()

	// operations
	upi := operation.NewUploadImage(cs, time.Duration(storageCfg.UploadFileTimeoutMinutes))
	lfo := operation.NewListImage(cs)
	coio := operation.NewConvertImage(conv)

	// handlers
	cfh := handler.NewCreateImageHandler(upi, coio)
	lfh := handler.NewListImage(lfo)

	// controllers
	cuf := controller.NewCreateImage(cfh)
	clf := controller.NewImageFile(lfh)

	cuf.Register(ge)
	clf.Register(ge)
}
