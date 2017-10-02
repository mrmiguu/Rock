package main

import (
	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/iota/rest"
)

func main() {
	var msg rest.String

	done := load.New(`msg.Post("howdy!")`)
	msg.Post("howdy!")
	done <- true
}
