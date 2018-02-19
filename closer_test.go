package mock

import (
	"testing"

	"errors"

	assertFactory "github.com/stretchr/testify/assert"
)

func TestNewCloser(t *testing.T) {
	assert := assertFactory.New(t)
	c := NewCloser()
	assert.Equal(&closer{}, c)
}

func TestCloser_Close(t *testing.T) {
	assert := assertFactory.New(t)

	t.Run("close without error", func(t *testing.T) {
		c := &closer{}

		assert.False(c.closed)
		err := c.Close()

		assert.Nil(err)
		assert.True(c.closed)
	})

	t.Run("close with error", func(t *testing.T) {
		expectedError := errors.New("<expectedError>")
		c := &closer{err: expectedError}

		assert.False(c.closed)
		err := c.Close()

		assert.Equal(expectedError, err)
		assert.False(c.closed)
	})
}

func TestCloser_IsClosed(t *testing.T) {
	assert := assertFactory.New(t)
	c := &closer{closed: false}
	assert.False(c.IsClosed())

	c.closed = true
	assert.True(c.IsClosed())
}

func TestCloser_CloseShouldReturn(t *testing.T) {
	assert := assertFactory.New(t)
	c := &closer{}
	assert.Nil(c.err)

	expectedError := errors.New("<err>")
	c.CloseShouldReturn(expectedError)
	assert.Equal(expectedError, c.err)
}
