package util

import "io"

type ErrorWriter struct {
	w   io.Writer
	Err error
}

func NewErrorWriter(w io.Writer) ErrorWriter {
	return ErrorWriter{
		w: w,
	}
}

func (ew *ErrorWriter) Write(buf []byte) {
	if ew.Err != nil {
		return
	}
	_, ew.Err = ew.w.Write(buf)
}
