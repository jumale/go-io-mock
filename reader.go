package mock

import "io"

func NewReader() *reader {
	return &reader{hasEOF: true}
}

func NewReadCloser() *readCloser {
	return &readCloser{
		reader: NewReader(),
		closer: NewCloser(),
	}
}

type ReadCase struct {
	Data string
	Err  error
}

type reader struct {
	cases  []ReadCase
	hasEOF bool // if reader should return io.EOF when there is nothing left to read
}

type readCloser struct {
	*reader
	*closer
}

func (r *reader) ReaderHasEOF(hasEOF bool) {
	r.hasEOF = hasEOF
}

func (r *reader) AddReadCase(c ReadCase) {
	r.cases = append(r.cases, c)
}

func (r *reader) Read(p []byte) (n int, err error) {
	if len(r.cases) == 0 {
		if r.hasEOF {
			return 0, io.EOF
		} else {
			return 0, nil
		}
	}

	c := r.cases[0]
	r.cases = r.cases[1:]
	return copy(p, []byte(c.Data)), c.Err
}
