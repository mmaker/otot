package ecdh

import (
	"math/big"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mmaker/otot/encodings"
)



func TestECDH(t *testing.T) {
	in, out := net.Pipe()
	c1 := encodings.NewTConn(in, out)
	c2 := encodings.NewTConn(in, out)

	var wg sync.WaitGroup
	wg.Add(2)
	var Sx, Sy, Cx, Cy *big.Int
	go func() {
		Sx, Sy = StartServer(c1)
		wg.Done()
	}()
	go func() {
		Cx, Cy = StartClient(c2)
		wg.Done()
	}()

	wg.Wait()
	assert.Equal(t, Sx, Cx)
	assert.Equal(t, Sy, Cy)
}
