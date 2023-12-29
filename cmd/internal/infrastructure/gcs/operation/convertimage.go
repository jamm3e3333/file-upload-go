package operation

import (
	"jamm3e3333/file-upload/cmd/internal/application/dto"
	"jamm3e3333/file-upload/cmd/pkg/convertimage"
	"regexp"
)

type ConvertImage struct {
	convert convertimage.Convert
}

func NewConvertImage(convert convertimage.Convert) *ConvertImage {
	return &ConvertImage{convert: convert}
}

func (o *ConvertImage) ToWebp(img *dto.CreateImage) (*dto.CreateImage, error) {
	convImg, err := o.convert.ToWebp(img.Image)
	if err != nil {
		return nil, err
	}

	return &dto.CreateImage{
		Name:  o.toWebpImgName(img.Name),
		Image: convImg,
		Type:  img.Type,
	}, nil
}

func (o *ConvertImage) toWebpImgName(name string) string {
	re := regexp.MustCompile(`\.[a-zA-Z0-9]+$`)
	if re.MatchString(name) {
		return re.ReplaceAllString(name, "."+"webp")
	}
	return name + "." + "webp"
}
