package main

import (
	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/iota/rest"
)

func main() {
	var msg rest.String

	done := load.New("msg.Get()")
	println(msg.Get())
	done <- true
}
