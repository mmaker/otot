package dh

import (
	"math/big"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mmaker/otot/encodings"
)



func TestDH(t *testing.T) {
	in, out := net.Pipe()
	c1 := encodings.NewTConn(in, out)
	c2 := encodings.NewTConn(in, out)
//	fst = emojii.NewDecoder(fst)
//	snd = emojii.NewEncoder(snd)

	var wg sync.WaitGroup
	wg.Add(2)

	var sK, cK *big.Int
	go func() {
		sK = StartServer(c1)
		wg.Done()
	}()

	go func() {
		cK = StartClient(c2)
		wg.Done()
	}()

	wg.Wait()
	assert.Equal(t, cK, sK)
}
