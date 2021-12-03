package trash

import (
	"bufio"
	"io"
	"os"
)

func MustLoadFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return f
}

func ReaderToStrings(f io.Reader) (ss []string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}
	return
}
