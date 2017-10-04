package rock

func (S *String) makeW() {
	S.p.w.c = make(chan []byte, S.Len)
	go postIfClient(S.p.w.c, Tstring, S.Name)
}

func (S *String) makeR() {
	S.p.r.c = make(chan []byte, S.Len)
	go getIfClient(S.p.r.c, Tstring, S.Name)
}

func (S *String) makeNIfServer() {
	if IsClient {
		return
	}
	S.p.n.c = make(chan int)
}

func (S *String) add() {
	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string]*String{}
	}
	if _, found := stringDict.m[S.Name]; !found {
		stringDict.m[S.Name] = S
	}
	stringDict.Unlock()
}

func (S *String) to(s string) {
	if IsClient {
		S.p.w.c <- []byte(s)
		return
	}
	for {
		<-S.p.n.c
		S.p.w.c <- []byte(s)
		if len(S.p.n.c) == 0 {
			break
		}
	}
}

func (S *String) from() string {
	return string(<-S.p.r.c)
}

func (S *String) S() chan<- string {
	c := make(chan string, S.Len)
	go started.Do(getAndOrPostIfServer)
	S.add()
	S.p.w.Do(S.makeW)
	S.p.n.Do(S.makeNIfServer)
	go func() {
		S.to("")
		i := <-c
		close(c)
		S.to(i)
	}()
	return c
}

func (S *String) R() <-chan string {
	c := make(chan string, S.Len)
	go started.Do(getAndOrPostIfServer)
	S.add()
	S.p.r.Do(S.makeR)
	go func() {
		S.from()
		c <- S.from()
		close(c)
	}()
	return c
}
