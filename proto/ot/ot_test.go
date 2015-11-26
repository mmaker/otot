package ot

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mmaker/otot/encodings"
)

func TestOT(t *testing.T) {
	in, out := net.Pipe()
	c1 := encodings.NewTConn(in, out)
	c2 := encodings.NewTConn(in, out)

	var wg sync.WaitGroup
	wg.Add(2)

	choices := []string{"first", "second"}
	var got string
	go func() {
		StartSender(c1, choices)
		wg.Done()
	}()
	go func () {
		got = StartReceiver(c2, 1)
		wg.Done()
	}()

	wg.Wait()
	assert.Equal(t, got, "second")
}
