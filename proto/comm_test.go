package proto

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mmaker/otot/encodings"
)


func TestSendRecv(t *testing.T) {
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)
	s := encodings.NetstringScanner(buf)

	fst := []byte("hello")
	snd := []byte("world")
	thr := []byte("!")
	SendBytes(w, fst, snd)
	SendBytes(w, thr)

	fstgot := RecvBytes(s)
	assert.Equal(t, fstgot, fst)
	sndgot := RecvBytes(s)
	assert.Equal(t, sndgot, snd)
	thrgot := RecvBytes(s)
	assert.Equal(t, thrgot, thr)
}
