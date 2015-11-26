package dh

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/mmaker/otot/encodings"
	"github.com/mmaker/otot/proto"
)

func StartClient(c *encodings.TConn) *big.Int {
	mod, _ := rand.Prime(rand.Reader, proto.BITS)
	g, _ := rand.Int(rand.Reader, mod)
	a, _ := rand.Int(rand.Reader, mod)
	A := big.NewInt(0)
	A.Exp(g, a, mod)

	c.SendBigInt(mod, g, A)
	B := c.RecvBigInt()
	return 	B.Exp(B, a, mod)
}

func StartServer(c *encodings.TConn) *big.Int {
	fmt.Println("Listening.")
	mod := c.RecvBigInt()
	g := c.RecvBigInt()
	A := c.RecvBigInt()

	b, _ := rand.Int(rand.Reader, mod)
	B := big.NewInt(0)
	B.Exp(g, b, mod)
	c.SendBigInt(B)

	return A.Exp(A, b, mod)
}
