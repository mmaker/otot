package encodings

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSendRecv(t *testing.T) {
	buf := new(bytes.Buffer)
	c := NewTConn(buf, buf)

	fst := []byte("hello")
	snd := []byte("world")
	thr := []byte("!")
	c.SendBytes(fst, snd)
	c.SendBytes(thr)

	fstgot := c.RecvBytes()
	assert.Equal(t, fstgot, fst)
	sndgot := c.RecvBytes()
	assert.Equal(t, sndgot, snd)
	thrgot := c.RecvBytes()
	assert.Equal(t, thrgot, thr)
}
