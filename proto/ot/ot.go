package ot

import (
	"bufio"
	"crypto/rand"
	"log"
	"io"
	"math/big"

	"github.com/mmaker/otot/encodings"
	"github.com/mmaker/otot/proto"
)

func StartSender(in io.Reader, out io.Writer, choices []string) {
	w := bufio.NewWriter(out)
	s := encodings.NetstringScanner(in)

	mod, _ := rand.Prime(rand.Reader, proto.BITS)
	g, _ := rand.Int(rand.Reader, mod)
	a, _ := rand.Int(rand.Reader, mod)
	A := big.NewInt(0)
	A.Exp(g, a, mod)

	proto.SendBigInt(w, mod, g, A)

	B := proto.RecvBigInt(s)

	nkey := big.NewInt(0)
	nkey.Exp(B, a, mod)
	k0 := proto.Hash(nkey.Bytes())

	nkey.ModInverse(A, mod)
	nkey.Mul(nkey, B)
	nkey.Exp(nkey, a, mod)
	k1 := proto.Hash(nkey.Bytes())


	e0 := proto.Encrypt(k0, []byte(choices[0]))
	e1 := proto.Encrypt(k1, []byte(choices[1]))
	proto.SendBytes(w, e0, e1)
}


func StartReceiver(in io.Reader, out io.Writer, choice int) (msg string) {
	log.Println("Listening.")
	s := encodings.NetstringScanner(in)
	w := bufio.NewWriter(out)

	mod := proto.RecvBigInt(s)
	g := proto.RecvBigInt(s)
	A := proto.RecvBigInt(s)

	b, _ := rand.Int(rand.Reader, mod)
	B := big.NewInt(0)
	if choice == 0 {
		B.Exp(g, b, mod)
	} else if choice == 1 {
		B.Exp(g, b, mod)
		B.Mul(B, A)
		B.Mod(B, mod)
	}
	proto.SendBigInt(w, B)

	tmp := big.NewInt(0)
	tmp.Exp(A, b, mod)
	k := proto.Hash(tmp.Bytes())

	e0 := proto.RecvBytes(s)
	e1 := proto.RecvBytes(s)

	var data []byte
	if choice == 0 {
		data = proto.Decrypt(k, e0)
	} else {
		data = proto.Decrypt(k, e1)
	}
	msg = string(data)
	return

}
