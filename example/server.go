package main

import (
	"time"

	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/rock"
)

func main() {
	var msg rock.String

	start := rock.String{Name: "start"}
	done := load.New("starting")
	start.From()
	done <- true

	then := time.Now()
	for range [100]int{} {
		msg.To("Hello")
	}
	println(int(float64(time.Since(then).Nanoseconds())/100.0)/1000000, `ms (To)`)

	then = time.Now()
	for range [100]int{} {
		msg.From()
	}
	println(int(float64(time.Since(then).Nanoseconds())/100.0)/1000000, `ms (From)`)
}
