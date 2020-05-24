package rw

import (
	"io"
)

type MockWriter struct {
}

func NewMockWriter() *MockWriter {
	return &MockWriter{}
}

func (m MockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

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
