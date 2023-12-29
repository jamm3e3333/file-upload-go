package command

import "io"

type Image interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

type CreateImage struct {
	Image Image
	Name  string
	Type  string
}
