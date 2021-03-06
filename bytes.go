package rock

func (B *Bytes) makeW() {
	B.p.w.c = make(chan []byte, B.Len)
	go postIfClient(B.p.w.c, Tbytes, B.Name)
}

func (B *Bytes) makeR() {
	B.p.r.c = make(chan []byte, B.Len)
	go getIfClient(B.p.r.c, Tbytes, B.Name)
}

func (B *Bytes) makeNIfServer() {
	if IsClient {
		return
	}
	B.p.n.c = make(chan int)
}

func (B *Bytes) add() {
	bytesDict.Lock()
	if bytesDict.m == nil {
		bytesDict.m = map[string]*Bytes{}
	}
	if _, found := bytesDict.m[B.Name]; !found {
		bytesDict.m[B.Name] = B
	}
	bytesDict.Unlock()
}

func (B *Bytes) to(b []byte) {
	if IsClient {
		B.p.w.c <- b
		return
	}
	for {
		<-B.p.n.c
		B.p.w.c <- b
		if len(B.p.n.c) == 0 {
			break
		}
	}
}

func (B *Bytes) from() []byte {
	return <-B.p.r.c
}

func (B *Bytes) S() chan<- []byte {
	c := make(chan []byte, B.Len)
	go started.Do(getAndOrPostIfServer)
	B.add()
	B.p.w.Do(B.makeW)
	B.p.n.Do(B.makeNIfServer)
	go func() {
		B.to([]byte{0})
		i := <-c
		close(c)
		B.to(i)
	}()
	return c
}

func (B *Bytes) R() <-chan []byte {
	c := make(chan []byte, B.Len)
	go started.Do(getAndOrPostIfServer)
	B.add()
	B.p.r.Do(B.makeR)
	go func() {
		B.from()
		c <- B.from()
		close(c)
	}()
	return c
}
