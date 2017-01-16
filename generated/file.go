package generated

import "io"

type File struct {
	Reader   io.Reader
	Name     string
	MimeType string
	Size     int64
}
