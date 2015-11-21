package ot

import (
	"net"
	"strings"
	"sync"
	"testing"
)

func TestOT(t *testing.T) {
	fstc, sndc := net.Pipe()

	var wg sync.WaitGroup
	wg.Add(2)

	choices := []string{"first", "second"}
	var got string
	go func() {
		StartSender(fstc, sndc, choices)
		wg.Done()
	}()
	go func () {
		got = StartReceiver(fstc, sndc, 1)
		wg.Done()
	}()

	wg.Wait()
	if strings.Compare(got, "second") != 0 {
		t.Errorf("Error: got '%s')", got)
	}

}
