package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func PrintFile(filepath string) {
	file, err := os.Open(filepath)
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

func ReadFile(filePath string) []byte {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err.Error())
	}

	var bufSize int64 = 100

	if fileInfo, e := file.Stat(); e != nil {
		panic(e.Error())
	} else if fileSize := fileInfo.Size(); bufSize < fileSize {
		bufSize = fileSize
	}

	buf := make([]byte, bufSize)

	if _, e := file.Read(buf); e != nil {
		panic(e.Error())
	}

	return buf
}

func PrintFile2(filepath string) {
	buf, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(err.Error())
	}

	fmt.Print(string(buf))
}
