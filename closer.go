package mock

func NewCloser() *closer {
	return &closer{}
}

type closer struct {
	closed bool
	err    error
}

func (c *closer) Close() error {
	if c.err == nil {
		c.closed = true
	}
	return c.err
}

func (c closer) IsClosed() bool {
	return c.closed
}

func (c *closer) CloseShouldReturn(err error) {
	c.err = err
}
