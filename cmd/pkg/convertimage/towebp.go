package convertimage

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

type Convert interface {
	ToWebp(i io.Reader) (io.Reader, error)
}

type Converter struct{}

func NewImgConverter() *Converter {
	return &Converter{}
}

func (*Converter) ToWebp(i io.Reader) (io.Reader, error) {
	img, _, err := image.Decode(i)
	if err != nil {
		return nil, err
	}

	encOpt, optErr := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 50)
	if optErr != nil {
		return nil, err
	}

	// TODO: is it possible to process and convert the file by chunks?
	buf := new(bytes.Buffer)

	encErr := webp.Encode(buf, img, encOpt)
	if encErr != nil {
		return nil, encErr
	}

	return bytes.NewReader(buf.Bytes()), nil
}
