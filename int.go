package rock

func (I *Int) makeW() {
	I.p.w.c = make(chan []byte, I.Len)
	go postIfClient(I.p.w.c, Tint, I.Name)
}

func (I *Int) makeR() {
	I.p.r.c = make(chan []byte, I.Len)
	go getIfClient(I.p.r.c, Tint, I.Name)
}

func (I *Int) makeNIfServer() {
	if IsClient {
		return
	}
	I.p.n.c = make(chan int)
}

func (I *Int) add() {
	intDict.Lock()
	if intDict.m == nil {
		intDict.m = map[string]*Int{}
	}
	if _, found := intDict.m[I.Name]; !found {
		intDict.m[I.Name] = I
	}
	intDict.Unlock()
}

func (I *Int) to(i int) {
	if IsClient {
		I.p.w.c <- int2bytes(i)
		return
	}
	for {
		<-I.p.n.c
		I.p.w.c <- int2bytes(i)
		if len(I.p.n.c) == 0 {
			break
		}
	}
}

func (I *Int) from() int {
	return bytes2int(<-I.p.r.c)
}

func (I *Int) S() chan<- int {
	c := make(chan int, I.Len)
	go started.Do(getAndOrPostIfServer)
	I.add()
	I.p.w.Do(I.makeW)
	I.p.n.Do(I.makeNIfServer)
	go func() {
		I.to(0)
		i := <-c
		close(c)
		I.to(i)
	}()
	return c
}

func (I *Int) R() <-chan int {
	c := make(chan int, I.Len)
	go started.Do(getAndOrPostIfServer)
	I.add()
	I.p.r.Do(I.makeR)
	go func() {
		I.from()
		c <- I.from()
		close(c)
	}()
	return c
}
