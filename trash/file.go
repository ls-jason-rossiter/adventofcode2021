package trash

import "os"

func MustLoadFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return f
}
