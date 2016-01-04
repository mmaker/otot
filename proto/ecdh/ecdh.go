package ecdh

import (
	"crypto/rand"
	"crypto/elliptic"
	"log"
	"math/big"

	"github.com/mmaker/otot/encodings"
)

var E elliptic.Curve = elliptic.P521()

func StartClient(c *encodings.TConn) (*big.Int, *big.Int) {
	priv, x, y, err := elliptic.GenerateKey(E, rand.Reader)
	if err != nil {
		c.Abort(err)
	}
	// XXX. we don't really need to send both coordinates. Send y and one
	// bit is enough.
	c.SendBigInt(x, y)

	Sx := c.RecvBigInt()
	Sy := c.RecvBigInt()
	return E.ScalarMult(Sx, Sy, priv)
}


func StartServer(c *encodings.TConn) (*big.Int, *big.Int) {
	log.Println("Listening.")
	priv, x, y, err := elliptic.GenerateKey(E, rand.Reader)
	if err != nil {
		c.Abort(err)
	}
	Cx := c.RecvBigInt()
	Cy := c.RecvBigInt()
	c.SendBigInt(x, y)

	return E.ScalarMult(Cx, Cy, priv)
}
