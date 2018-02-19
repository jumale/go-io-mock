package mock

func NewFile(name string) *file {
	return &file{
		name:   name,
		reader: NewReader(),
		writer: NewWriter(),
		closer: NewCloser(),
		seeker: NewSeeker(),
	}
}

type file struct {
	name string
	*reader
	*writer
	*closer
	*seeker
}

func (f file) Name() string {
	return f.name
}
