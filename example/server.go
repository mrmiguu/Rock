package main

import (
	"time"

	"github.com/mrmiguu/Loading"
	"github.com/mrmiguu/rock"
)

func main() {
	var i rock.Int

	strt := rock.Int{}
	done := load.New("starting")
	<-strt.R()
	done <- true

	timeout := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		timeout <- true
	}()

	var n int
	start := time.Now()
	for {
		select {
		case <-timeout:
			println("timeout!", int(float64(time.Since(start).Nanoseconds())/2000000/float64(n)), "ms")
			return
		case n = <-i.R():
			println(n)
			i.S() <- n + 1
		}
	}
}
