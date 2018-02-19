package mock

import (
	"testing"

	"io"

	"errors"

	assertFactory "github.com/stretchr/testify/assert"
)

func TestNewWriter(t *testing.T) {
	assert := assertFactory.New(t)
	w := NewWriter()
	assert.Equal(&writer{}, w)
}

func TestNewWriteCloser(t *testing.T) {
	assert := assertFactory.New(t)
	wc := NewWriteCloser()
	assert.Implements((*io.Writer)(nil), wc)
	assert.Implements((*io.Closer)(nil), wc)
}

func TestWriter_Write(t *testing.T) {
	assert := assertFactory.New(t)

	t.Run("should append data multiple times, and return added length and empty error", func(t *testing.T) {
		w := &writer{}

		n, err := w.Write([]byte("<foo>"))
		assert.Equal(5, n)
		assert.Nil(err)
		assert.Equal("<foo>", w.WrittenData())

		n, err = w.Write([]byte("<en>"))
		assert.Equal(4, n)
		assert.Nil(err)
		assert.Equal("<foo><en>", w.WrittenData())
	})

	t.Run("should return specified length and error", func(t *testing.T) {
		expectedN := 7
		expectedError := errors.New("<err>")
		w := &writer{returnN: expectedN, returnErr: expectedError}

		n, err := w.Write([]byte("foo"))
		assert.Equal(expectedN, n)
		assert.Equal(expectedError, err)
		assert.Equal("foo", w.WrittenData())

	})
}

func TestWriter_WriteShouldReturn(t *testing.T) {
	assert := assertFactory.New(t)
	var logMsg string
	w := &writer{
		// mock log-function to store log to a local variable
		log: func(v ...interface{}) {
			logMsg = v[0].(string)
		},
	}

	t.Run("should set specified values", func(t *testing.T) {
		expectedN := int(6)
		expectedError := errors.New("<err>")
		w.WriteShouldReturn(expectedN, expectedError)
		assert.Equal(expectedN, w.returnN)
		assert.Equal(expectedError, w.returnErr)
		assert.Empty(logMsg)
	})

	t.Run("should set nil values", func(t *testing.T) {
		w.WriteShouldReturn(nil, nil)
		assert.Nil(w.returnN)
		assert.Nil(w.returnErr)
		assert.Empty(logMsg)
	})

	t.Run("should fail if 'n' is not int or nil", func(t *testing.T) {
		w.WriteShouldReturn(int64(7), errors.New("<err>"))
		assert.Nil(w.returnN)
		assert.Nil(w.returnErr)
		assert.Contains(logMsg, "int or nil")
	})
}
