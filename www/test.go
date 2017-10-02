package main

import "github.com/mrmiguu/iota/rest"

func main() {
	var msg rest.String

	msg.Post("howdy!")
}
