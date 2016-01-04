package encodings

import (
	"bufio"
	"encoding/hex"
	"io"
	"log"
	"math/big"
)


type TConn struct {
	s *bufio.Scanner
	w *bufio.Writer
}

func check(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func (t *TConn) checkScan() {
	if !t.s.Scan() {
		log.Fatalf("Impossible to read. (%s)", t.s.Err())
	}
}

func (t *TConn) SendBytes(msgs ...[]byte) {
	var err error
	for _, msg := range msgs {
		buf := make([]byte, hex.EncodedLen(len(msg)))
		hex.Encode(buf, msg)

		msg = MarshalBytes(buf)
		_, err = t.w.Write(msg)
		check(err)
	}
	err = t.w.Flush()
	check(err)
}

func (t *TConn) RecvBytes() []byte {
	t.checkScan()
	recv := t.s.Bytes()
	buf := make([]byte, hex.DecodedLen(len(recv)))
	hex.Decode(buf, recv)
	return buf
}


func (t *TConn) RecvBigInt() *big.Int {
	p := t.RecvBytes()
	n := big.NewInt(0)
	n.SetBytes(p)
	return n
}

func (t *TConn) SendBigInt(ns...*big.Int) {
	ss := make([][]byte, len(ns))
	for i, n := range ns {
		ss[i] = n.Bytes()
	}
	t.SendBytes(ss...)
}


// XXX. this function should be used by check, scanCheck, and other methods that
// might abort.
func (t *TConn) Abort(err error) {
	log.Fatalf("%s\n", err)
}

func NewTConn(in io.Reader, out io.Writer) *TConn {
	return &TConn{
		s: NetstringScanner(in),
		w: bufio.NewWriter(out),
	}
}
