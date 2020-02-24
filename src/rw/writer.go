package rw

import (
	"io"
)

type BadWriter struct {
	err error
}

func NewBadWriter() io.Writer {

	return NewBadWriterWithErr(BadWriterErr)
}

func NewBadWriterWithErr(err error) io.Writer {

	return &BadWriter{
		err: err,
	}
}

func (b *BadWriter) Write(_ []byte) (n int, err error) {
	return 0, b.err
}
