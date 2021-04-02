package file

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	SampleFilePath = "sample.txt"
)

func TestPrintFile(t *testing.T) {
	fmt.Println("=========================")
	PrintFile(SampleFilePath)
	fmt.Println("=========================")
}

func TestPrintFile2(t *testing.T) {
	fmt.Println("=========================")
	PrintFile2(SampleFilePath)
	fmt.Println("=========================")
}

func TestReadFile(t *testing.T) {
	want, _ := ioutil.ReadFile(SampleFilePath)
	got := ReadFile(SampleFilePath)

	if len(got) != len(want) {
		panic(fmt.Sprintf("want size : %v, got size: %v\n", len(want), len(got)))
	}

	for i := range got {
		if got[i] != want[i] {
			fmt.Printf("got: %v, want: %v\n", got[i], want[i])
			panic("Test Failed")
		}
	}
}
