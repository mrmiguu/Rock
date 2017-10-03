package main

import (
	"github.com/mrmiguu/rock"
)

func main() {
	var i rock.Int

	I := -1
	for range [1000]int{} {
		println(I)
		i.To(I)
		I = i.From()
	}

	select {}
}
