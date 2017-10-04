package main

import (
	"github.com/mrmiguu/rock"
)

func main() {
	var i rock.Int

	strt := rock.Int{}
	strt.S() <- 0

	n := 1
	for range [1000]int{} {
		i.S() <- n
		n = <-i.R()
		println(n)
	}

	select {}
}
