package command

import "io"

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

// CreateFile TODO: Add file name and file type
type CreateFile struct {
	File File
}
