package emojii

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	m := []byte("hello world!")
	buf := new(bytes.Buffer)
	e := NewEncoder(buf)
	d := NewDecoder(buf)
	e.Write(m)

	data := make([]byte, 1000)
	n, err := d.Read(data)
	assert.Nil(t, err)
	assert.Equal(t, data[:n], m)
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
	assert.Nil(t, err)
	assert.Equal(t, n, 255)
	for i, c := range data[:255] {
		if int(i) != int(c) {
			t.Errorf("'%x' != '%x'", int(c), int(i))
		}
	}

}
