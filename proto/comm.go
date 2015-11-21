package proto

import (
	"bufio"
	"log"
)

const BITS = 1024

func Check(err error) {
	if err != nil {
		log.Fatal("%s", err)
	}
}



func Send(w *bufio.Writer, msgs ...[]byte) {
	var err error
	for _, msg := range msgs {
		_, err = w.Write(msg)
		Check(err)
	}
	err = w.Flush()
	Check(err)
}


func CheckScan(s *bufio.Scanner) {
	if !s.Scan() {
		log.Fatalf("Impossible to read. (%s)", s.Err())
	}
}
