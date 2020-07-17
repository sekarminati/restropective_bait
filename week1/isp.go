package main

type ReadWriter interface {
	Read() string
	Write() string
}

type Reader struct {
}

type Writer struct {
}

func (r Reader) Read() string {
	return "Reading"
}

func (w Writer) Write() string {
	return "Writing"
}
