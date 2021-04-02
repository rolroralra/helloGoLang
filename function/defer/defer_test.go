package _defer

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func Test(t *testing.T) {
	file, err := os.Open("sample.txt")
	defer file.Close()

	if err != nil {
		panic(err.Error())
	}

	buf := make([]byte, 10)
	for {
		n, e := file.Read(buf)

		if e != nil {
			if e == io.EOF {
				break
			}

			panic(e.Error())
		}

		fmt.Print(string(buf[:n]))
	}
}
