package encodings

import (
	"bufio"
	"fmt"
	"io"
)

func isMarshal(data []byte) (off int, total int) {
	var i int
	for i = range data {
		if data[i] == ':' {
			break
		}
	}
	_, err := fmt.Sscanf(string(data[:i]), "%X", &total)
	off = i + 1
	total += off
	if total <= len(data) &&  err == nil {
		return
	} else {
		return -1, 0
	}
}

func MarshalBytes(s []byte) []byte {
	len := []byte(fmt.Sprintf("%X:", len(s)))
	return append(len, s...)
}

func UnmarshalBytes(data []byte) []byte {
	off, total := isMarshal(data)
	if off < 0 {
		return nil
	}
	return data[off:total]
}


func split(data []byte, atEOF bool) (int, []byte, error) {
	size, total := isMarshal(data)
	if size < 0 && atEOF {
		return 0, nil, io.EOF
	} else if size < 0 {
		return 0, nil, nil
	} else {
		return total, data[size:total], nil
	}
}

func NetstringScanner(r io.Reader) (s *bufio.Scanner) {
	s =  bufio.NewScanner(r)
	s.Split(split)
	return s
}
