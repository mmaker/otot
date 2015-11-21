package dh

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	"github.com/mmaker/otot/encodings"
	"github.com/mmaker/otot/proto"
)

func StartClient(in io.Reader, out io.Writer) *big.Int {
	mod, _ := rand.Prime(rand.Reader, proto.BITS)
	g, _ := rand.Int(rand.Reader, mod)
	a, _ := rand.Int(rand.Reader, mod)
	A := big.NewInt(0)
	A.Exp(g, a, mod)

	w := bufio.NewWriter(out)
	s := encodings.NetstringScanner(in)

	proto.SendBigInt(w, mod, g, A)
	B := proto.RecvBigInt(s)
	return 	B.Exp(B, a, mod)
}

func StartServer(in io.Reader, out io.Writer) *big.Int {
	fmt.Println("Listening.")
	s := encodings.NetstringScanner(in)
	w := bufio.NewWriter(out)

	mod := proto.RecvBigInt(s)
	g := proto.RecvBigInt(s)
	A := proto.RecvBigInt(s)

	b, _ := rand.Int(rand.Reader, mod)
	B := big.NewInt(0)
	B.Exp(g, b, mod)
	proto.SendBigInt(w, B)

	return A.Exp(A, b, mod)
}
