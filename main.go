package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const (
	maxN int = 100
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other OS")
	}

	fmt.Println("Today is", time.Now().Weekday())
	switch today := time.Now().Weekday(); today {
	case time.Monday:
	case time.Tuesday:
	case time.Wednesday:
	case time.Thursday:
	case time.Friday:
	case time.Saturday:
	case time.Sunday:
	default:
	}

}
