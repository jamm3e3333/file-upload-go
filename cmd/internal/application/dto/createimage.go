package dto

import "io"

type CreateImage struct {
	Image io.Reader
	Name  string
	Type  string
}
