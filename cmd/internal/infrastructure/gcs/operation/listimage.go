package operation

import (
	"context"
	"jamm3e3333/file-upload/cmd/internal/application/dto"
	"jamm3e3333/file-upload/cmd/pkg/gcs"
	"path"
)

type ListImage struct {
	lf gcs.FilesInfo
}

func NewListImage(lf gcs.FilesInfo) *ListImage {
	return &ListImage{lf: lf}
}

// Execute TODO: add cache
func (o *ListImage) Execute(ctx context.Context) ([]*dto.ListImage, error) {
	f, err := o.lf.Files(ctx)
	images := make([]*dto.ListImage, 0, len(f))
	for _, v := range f {
		images = append(images, &dto.ListImage{
			Name:      o.imageNameFromPath(v.Path),
			SignedUrl: v.SignedUrl,
		})
	}

	if err != nil {
		return nil, err
	}

	return images, nil
}

func (o *ListImage) imageNameFromPath(imagePath string) string {
	return path.Base(imagePath)
}
