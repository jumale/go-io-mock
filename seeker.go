package mock

import (
	"log"
)

func NewSeeker() *seeker {
	return &seeker{log: log.Fatal}
}

type seeker struct {
	offset       int64
	whence       int
	returnOffset interface{}
	returnErr    error
	log          func(v ...interface{})
}

func (s *seeker) Seek(offset int64, whence int) (int64, error) {
	s.offset = offset
	s.whence = whence

	if s.returnOffset != nil {
		offset = s.returnOffset.(int64)
	}
	return offset, s.returnErr
}

func (s seeker) SeekOffset() int64 {
	return s.offset
}

func (s seeker) SeekWhence() int {
	return s.whence
}

func (s *seeker) SeekShouldReturn(offset interface{}, err error) {
	if offset != nil {
		if _, ok := offset.(int64); !ok {
			s.log("method SeekShouldReturn expects 'offset' to be int64 or nil")
			return
		}
	}
	s.returnOffset = offset
	s.returnErr = err
}
