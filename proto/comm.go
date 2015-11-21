package proto

import (
	"bufio"
	"log"
	"encoding/hex"
	"math/big"

	"github.com/mmaker/otot/encodings"
)

const BITS = 256

func Check(err error) {
	if err != nil {
		log.Fatal("%s", err)
	}
}


func RecvBytes(s *bufio.Scanner) []byte {
	CheckScan(s)
	recv := s.Bytes()
	buf := make([]byte, hex.DecodedLen(len(recv)))
	hex.Decode(buf, recv)
	return buf
}

func SendBytes(w *bufio.Writer, msgs ...[]byte) {
	var err error
	for _, msg := range msgs {
		buf := make([]byte, hex.EncodedLen(len(msg)))
		hex.Encode(buf, msg)

		msg = encodings.MarshalBytes(buf)
		_, err = w.Write(msg)
		Check(err)
	}
	err = w.Flush()
	Check(err)
}

func CheckScan(s *bufio.Scanner) {
	if !s.Scan() {
		log.Fatalf("Impossible to read. (%s)", s.Err())
	}
}

func RecvBigInt(s *bufio.Scanner) *big.Int{
	p := RecvBytes(s)
	n := big.NewInt(0)
	n.SetBytes(p)
	return n
}

func SendBigInt(w *bufio.Writer, ns...*big.Int) {
	ss := make([][]byte, len(ns))
	for i, n := range ns {
		ss[i] = n.Bytes()
	}
	SendBytes(w, ss...)
}
