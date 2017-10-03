package main

import (
	"github.com/mrmiguu/rock"
)

func main() {
	var i rock.Int

	for range [1000]int{} {
		I := i.From()
		println(I)
		i.To(I - 1)
	}

	select {}
}
