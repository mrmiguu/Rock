package rock

func (I *Int) makeW() {
	I.p.w.c = make(chan []byte, I.Len)
	go postIfClient(I.p.w.c, Tint, I.Name)
}

func (I *Int) makeR() {
	I.p.r.c = make(chan []byte, I.Len)
	go getIfClient(I.p.r.c, Tint, I.Name)
}

func (I *Int) makeN() {
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

func (I *Int) To(i int) {
	go started.Do(getAndOrPostIfServer)

	I.add()

	I.p.w.Do(I.makeW)
	if IsClient {
		I.p.w.c <- int2bytes(i)
		return
	}

	I.p.n.Do(I.makeN)
	for {
		<-I.p.n.c
		I.p.w.c <- int2bytes(i)
		if len(I.p.n.c) == 0 {
			break
		}
	}
}

func (I *Int) From() int {
	go started.Do(getAndOrPostIfServer)

	I.add()

	I.p.r.Do(I.makeR)
	return bytes2int(<-I.p.r.c)
}
