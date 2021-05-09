package rw

import (
	"io"
)

type MockReader struct {
}

func NewMockReader() *MockReader {
	return &MockReader{}
}

func (m MockReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

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
