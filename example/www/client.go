package main

import (
	"github.com/mrmiguu/rock"
)

func main() {
	var msg rock.String

	start := rock.String{Name: "start"}
	start.To("!")

	for range [100]int{} {
		msg.From()
	}
	for range [100]int{} {
		msg.To("World")
	}
}
