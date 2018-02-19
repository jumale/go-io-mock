package mock

import (
	"testing"

	"io"

	"errors"

	assertFactory "github.com/stretchr/testify/assert"
)

func TestNewReader(t *testing.T) {
	assert := assertFactory.New(t)
	r := NewReader()
	assert.True(r.hasEOF, "by default reader will return io.EOF, when there is nothing to read anymore")
}

func TestNewReadCloser(t *testing.T) {
	assert := assertFactory.New(t)
	rc := NewReadCloser()
	assert.Implements((*io.Reader)(nil), rc)
	assert.Implements((*io.Closer)(nil), rc)
}

func TestReader_Read(t *testing.T) {
	assert := assertFactory.New(t)

	t.Run("should return 0 and nil if there is nothing to read and reader does not have EOF", func(t *testing.T) {
		r := reader{cases: []ReadCase{}, hasEOF: false}
		var buf []byte
		n, err := r.Read(buf)
		assert.Equal(0, n)
		assert.Nil(err)
	})

	t.Run("should 0 and io.EOF it there is nothing to read and reader has EOF", func(t *testing.T) {
		r := reader{cases: []ReadCase{}, hasEOF: true}
		var buf []byte
		n, err := r.Read(buf)
		assert.Equal(0, n)
		assert.Equal(io.EOF, err)
	})

	t.Run("should read cases one by one", func(t *testing.T) {
		r := reader{
			cases: []ReadCase{
				{Data: "foo"},
				{Data: "bart", Err: errors.New("<err>")},
			},
		}

		buf := make([]byte, 3)
		n, err := r.Read(buf)
		assert.EqualValues("foo", buf)
		assert.Equal(3, n)
		assert.Nil(err)

		buf = make([]byte, 4)
		n, err = r.Read(buf)
		assert.EqualValues("bart", buf)
		assert.Equal(4, n)
		assert.Equal("<err>", err.Error())
	})
}

func TestReader_AddReadCase(t *testing.T) {
	assert := assertFactory.New(t)
	r := &reader{}
	assert.Len(r.cases, 0)

	r.AddReadCase(ReadCase{Data: "foo"})
	assert.Len(r.cases, 1)
	assert.Equal("foo", r.cases[0].Data)
}

func TestReader_ReaderHasEOF(t *testing.T) {
	assert := assertFactory.New(t)
	r := &reader{}
	assert.False(r.hasEOF)

	r.ReaderHasEOF(true)
	assert.True(r.hasEOF)
}
