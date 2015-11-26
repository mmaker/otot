package ot

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/mmaker/otot/encodings"
	"github.com/mmaker/otot/proto"
)

func StartSender(c *encodings.TConn, choices []string) {
	mod, _ := rand.Prime(rand.Reader, proto.BITS)
	g, _ := rand.Int(rand.Reader, mod)
	a, _ := rand.Int(rand.Reader, mod)
	A := big.NewInt(0)
	A.Exp(g, a, mod)
	c.SendBigInt(mod, g, A)

	B := c.RecvBigInt()

	nkey := big.NewInt(0)
	nkey.Exp(B, a, mod)
	k0 := proto.Hash(nkey.Bytes())

	nkey.ModInverse(A, mod)
	nkey.Mul(nkey, B)
	nkey.Exp(nkey, a, mod)
	k1 := proto.Hash(nkey.Bytes())


	e0 := proto.Encrypt(k0, []byte(choices[0]))
	e1 := proto.Encrypt(k1, []byte(choices[1]))
	c.SendBytes(e0, e1)
}


func StartReceiver(c *encodings.TConn, choice int) (msg string) {
	log.Println("Listening.")

	mod := c.RecvBigInt()
	g := c.RecvBigInt()
	A := c.RecvBigInt()

	b, _ := rand.Int(rand.Reader, mod)
	B := big.NewInt(0)
	if choice == 0 {
		B.Exp(g, b, mod)
	} else if choice == 1 {
		B.Exp(g, b, mod)
		B.Mul(B, A)
		B.Mod(B, mod)
	}
	c.SendBigInt(B)

	tmp := big.NewInt(0)
	tmp.Exp(A, b, mod)
	k := proto.Hash(tmp.Bytes())

	e0 := c.RecvBytes()
	e1 := c.RecvBytes()

	var data []byte
	if choice == 0 {
		data = proto.Decrypt(k, e0)
	} else {
		data = proto.Decrypt(k, e1)
	}
	msg = string(data)
	return

}
