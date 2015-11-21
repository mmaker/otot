package proto

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/mmaker/otot/encodings"
)


func assertByteEqual(t *testing.T, a, b []byte) {
	if bytes.Compare(a, b) != 0 {
		t.Errorf("'%x' != '%x'", a, b)
	}
}

func TestSendRecv(t *testing.T) {
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)
	s := encodings.NetstringScanner(buf)

//	fst := []byte("hello")
//	snd := []byte("world")
	thr := []byte("!")
//	SendBytes(w, fst, snd)
	SendBytes(w, thr)

//	fstgot := RecvBytes(s)
//	assertByteEqual(t, fstgot, fst)
//	sndgot := RecvBytes(s)
	//	assertByteEqual(t, sndgot, snd)
	thrgot := RecvBytes(s)
	assertByteEqual(t, thrgot, thr)
}
