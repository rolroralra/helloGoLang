package main

import (
	"rsc.io/quote"
	_ "rsc.io/quote/v3"
)

func Hello() string {
	//return "Hello, world."
	return quote.Hello()
}
