package mock

func NewWriter() *writer {
	return &writer{}
}

func NewWriteCloser() *writeCloser {
	return &writeCloser{
		writer: NewWriter(),
		closer: NewCloser(),
	}
}

type writer struct {
	data      []byte
	returnN   interface{}
	returnErr error
	log       func(v ...interface{})
}

type writeCloser struct {
	*writer
	*closer
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.data = append(w.data, p...)

	n = len(p)
	if w.returnN != nil {
		n = w.returnN.(int)
	}
	return n, w.returnErr
}

func (w *writer) WriteShouldReturn(n interface{}, err error) {
	if n != nil {
		if _, ok := n.(int); !ok {
			w.log("method SeekShouldReturn expects 'n' to be int or nil")
			return
		}
	}
	w.returnN = n
	w.returnErr = err
}

func (w writer) WrittenData() string {
	return string(w.data)
}
