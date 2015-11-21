package dh

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"time"

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
	proto.Send(w,
		encodings.MarshalBigInt(mod),
		encodings.MarshalBigInt(g),
		encodings.MarshalBigInt(A))

	time.Sleep(5 * time.Second)
	s := encodings.NetstringScanner(in)
	B := big.NewInt(0)
	proto.CheckScan(s)
	B.SetBytes(s.Bytes())

	return 	B.Exp(B, a, mod)
}


func StartServer(in io.Reader, out io.Writer) *big.Int {
	fmt.Println("Listening.")
	s := encodings.NetstringScanner(in)

	proto.CheckScan(s)
	mod := big.NewInt(0)
	mod.SetBytes(s.Bytes())

	proto.CheckScan(s)
	g := big.NewInt(0)
	g.SetBytes(s.Bytes())

	proto.CheckScan(s)
	A := big.NewInt(0)
	A.SetBytes(s.Bytes())

	b, _ := rand.Int(rand.Reader, mod)
	B := big.NewInt(0)
	B.Exp(g, b, mod)
	w := bufio.NewWriter(out)
	proto.Send(w, encodings.MarshalBigInt(B))

	return A.Exp(A, b, mod)
}
