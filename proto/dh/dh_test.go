package dh

import (
	"io"
	"math/big"
	"net"
	"sync"
	"testing"
)



func TestDH(t *testing.T) {
	var fst io.Reader
	var snd io.Writer
	fst, snd = net.Pipe()
//	fst = emojii.NewDecoder(fst)
//	snd = emojii.NewEncoder(snd)

	var wg sync.WaitGroup
	wg.Add(2)

	var sK, cK *big.Int
	go func() {
		sK = StartServer(fst, snd)
		wg.Done()
	}()

	go func() {
		cK = StartClient(fst, snd)
		wg.Done()
	}()

	wg.Wait()
	if sK.Cmp(cK) != 0 {
		t.Errorf("%s != %s", sK, cK)
	}
}
