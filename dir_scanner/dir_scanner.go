package dir_scanner

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileContent(path string) []byte {
	dat, err := os.ReadFile(path)
	check(err)
	return dat
}
