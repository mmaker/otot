package arab

import (
	"io"
	"unicode/utf8"
)

type encoder struct {
	w io.Writer
}

type decoder struct {
	r io.Reader
}

func (e *encoder) Write(p []byte) (i int, err error) {
	var buf string
	for _, c := range p {
		s := string(rune(int(c) + 0x600))
		buf += s
	}
	return e.w.Write([]byte(buf))
}

func (d *decoder) Read(p []byte) (int, error) {
	buf := make([]byte, len(p))
	n, err := d.r.Read(buf)

	var r rune
	var i int = 0
	for j, w := 0, 0; j < n; j += w {
		r, w = utf8.DecodeRune(buf[j:])
		p[i] = byte(int(r) - 0x600)
		i++
	}
	return i, err
}



func NewEncoder(w io.Writer) io.Writer {
	return &encoder{
		w :w,
	}
}

func NewDecoder(r io.Reader) io.Reader {
	return &decoder{
		r: r,
	}
}
