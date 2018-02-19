package mock

import (
	"testing"

	"errors"

	assertFactory "github.com/stretchr/testify/assert"
)

func TestNewSeeker(t *testing.T) {
	assert := assertFactory.New(t)
	s := NewSeeker()
	t.Run("check the seeker is initialized and has minimal default setup", func(t *testing.T) {
		assert.NotNil(s.log)
	})
}

func TestSeeker_Seek(t *testing.T) {
	assert := assertFactory.New(t)

	t.Run("should store seek values", func(t *testing.T) {
		s := &seeker{}
		s.Seek(5, 7)
		assert.Equal(int64(5), s.SeekOffset())
		assert.Equal(int(7), s.SeekWhence())
	})

	t.Run("should return passed offset and no error", func(t *testing.T) {
		s := &seeker{}
		offset, err := s.Seek(5, 7)
		assert.Equal(int64(5), offset)
		assert.Nil(err)
	})

	t.Run("should return specified offset and specified error", func(t *testing.T) {
		expectedOffset := int64(3)
		expectedErr := errors.New("<logMsg>")
		s := &seeker{
			returnOffset: expectedOffset,
			returnErr:    expectedErr,
		}

		offset, logMsg := s.Seek(5, 7)
		assert.Equal(expectedOffset, offset)
		assert.Equal(expectedErr, logMsg)
	})

}

func TestSeeker_SeekShouldReturn(t *testing.T) {
	assert := assertFactory.New(t)
	var logMsg string
	s := &seeker{
		// mock log-function to store log to a local variable
		log: func(v ...interface{}) {
			logMsg = v[0].(string)
		},
	}

	t.Run("should set specified values", func(t *testing.T) {
		expectedOffset := int64(6)
		expectedError := errors.New("<err>")
		s.SeekShouldReturn(expectedOffset, expectedError)
		assert.Equal(expectedOffset, s.returnOffset)
		assert.Equal(expectedError, s.returnErr)
		assert.Empty(logMsg)
	})

	t.Run("should set nil values", func(t *testing.T) {
		s.SeekShouldReturn(nil, nil)
		assert.Nil(s.returnOffset)
		assert.Nil(s.returnErr)
		assert.Empty(logMsg)
	})

	t.Run("should fail if offset is not int64 or nil", func(t *testing.T) {
		s.SeekShouldReturn(int(7), errors.New("<err>"))
		assert.Nil(s.returnOffset)
		assert.Nil(s.returnErr)
		assert.Contains(logMsg, "int64 or nil")
	})
}
