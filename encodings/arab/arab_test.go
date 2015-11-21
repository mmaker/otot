package arab

import (
	"bytes"
	"testing"
)

func TestEncoder(t *testing.T) {
	m := []byte("hello world!")
	buf := new(bytes.Buffer)
	e := NewEncoder(buf)
	d := NewDecoder(buf)
	e.Write(m)

	data := make([]byte, 1000)
	n, err := d.Read(data)
	if n != len(m) || err != nil {
	 	t.Fail()
	}
	if bytes.Compare(data[:n], m) != 0{
	 	t.Fail()
	}

}

func TestByteRange(t *testing.T) {
	buf := new(bytes.Buffer)
	e := NewEncoder(buf)
	d := NewDecoder(buf)

	for i := 0; i != 255; i++ {
		s := []byte{byte(i)}
		e.Write(s)
	}
	data := make([]byte, 1000)
	n, err := d.Read(data)
	if n != 255 || err != nil {
		t.Errorf("Read %d bytes (%s)", n, err)
	}
	for i, c := range data[:255] {
		if int(i) != int(c) {
			t.Errorf("'%x' != '%x'", int(c), int(i))
		}
	}

}
