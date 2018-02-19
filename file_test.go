package mock

import (
	"testing"

	"io"

	assertFactory "github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	assert := assertFactory.New(t)
	f := NewFile("foo")
	assert.Equal("foo", f.name)
	assert.Implements((*io.Reader)(nil), f)
	assert.Implements((*io.Writer)(nil), f)
	assert.Implements((*io.Closer)(nil), f)
	assert.Implements((*io.Seeker)(nil), f)
}

func TestFile_Name(t *testing.T) {
	assert := assertFactory.New(t)
	f := file{}
	assert.Equal("", f.Name())
	f.name = "foo"
	assert.Equal("foo", f.Name())
}
