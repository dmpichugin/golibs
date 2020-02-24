package rw

import (
	"io"
)

func NewBadReader() io.Reader {

	return NewBadReaderWithErr(BadReaderErr)
}

func NewEOFReader() io.Reader {

	return NewBadReaderWithErr(io.EOF)
}

type BadReader struct {
	err error
}

func NewBadReaderWithErr(err error) io.Reader {

	return &BadReader{
		err: err,
	}
}

func (b *BadReader) Read(_ []byte) (n int, err error) {
	return 0, b.err
}
