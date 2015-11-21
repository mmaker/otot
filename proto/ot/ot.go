package ot

import (
	"bufio"
	"crypto/rand"
	"log"
	"io"
	"math/big"
	"time"

	"github.com/mmaker/otot/encodings"
	"github.com/mmaker/otot/proto"
)

func StartSender(in io.Reader, out io.Writer, choices []string) {
	s := encodings.NetstringScanner(in)

	mod, _ := rand.Prime(rand.Reader, proto.BITS)
	g, _ := rand.Int(rand.Reader, mod)

	a, _ := rand.Int(rand.Reader, mod)
	A := big.NewInt(0)
	A.Exp(g, a, mod)

	w := bufio.NewWriter(out)
	proto.Send(w,
		encodings.MarshalBigInt(mod),
		encodings.MarshalBigInt(g),
		encodings.MarshalBigInt(A))

	log.Println("Now waiting for receiver…")
	time.Sleep(10 * time.Second)
	proto.CheckScan(s)
	B := big.NewInt(0)
	B.SetBytes(s.Bytes())

	tmp := big.NewInt(0)
	k0 := proto.Hash(tmp.Exp(B, a, mod).Bytes())
	tmp.ModInverse(A, mod)
	tmp.Mul(tmp, B)
	tmp.Exp(tmp, a, mod)
	k1 := proto.Hash(tmp.Bytes())

	e0 := encodings.MarshalBytes(proto.Encrypt(k0, []byte(choices[0])))
	e1 := encodings.MarshalBytes(proto.Encrypt(k1, []byte(choices[1])))
	proto.Send(w, e0, e1)
}


func StartReceiver(in io.Reader, out io.Writer, choice int) (msg string) {
	log.Println("Listening.")
	s := encodings.NetstringScanner(in)

	proto.CheckScan(s)
	mod := big.NewInt(0)
	mod.SetBytes(s.Bytes())

	proto.CheckScan(s)
	g := big.NewInt(0)
	g.SetBytes(s.Bytes())

	proto.CheckScan(s)
	A := big.NewInt(0)
	A.SetBytes(s.Bytes())

	b, _ := rand.Int(rand.Reader, mod)
	B := big.NewInt(0)
	if choice == 0 {
		B.Exp(g, b, mod)
	} else if choice == 1 {
		B.Exp(g, b, mod)
		B.Mul(B, A)
		B.Mod(B, mod)
	}
	w := bufio.NewWriter(out)
	proto.Send(w, encodings.MarshalBigInt(B))

	tmp := big.NewInt(0)
	k := proto.Hash(tmp.Exp(A, b, mod).Bytes())

	proto.CheckScan(s)
	e0 := s.Bytes()
	proto.CheckScan(s)
	e1 := s.Bytes()

	var data []byte
	if choice == 0 {
		data = proto.Decrypt(k, e0)
	} else {
		data = proto.Decrypt(k, e1)
	}
	msg = string(data)
	return

}
