package rock

func (S *String) makeW() {
	S.p.w.c = make(chan []byte, S.Len)
	go postIfClient(S.p.w.c, Tstring, S.Name)
}

func (S *String) makeR() {
	S.p.r.c = make(chan []byte, S.Len)
	go getIfClient(S.p.r.c, Tstring, S.Name)
}

func (S *String) makeN() {
	S.p.n.c = make(chan int)
}

func (S *String) SelSend(s string) chan<- interface{} {
	send := make(chan interface{})
	go func() { S.To(s); <-send; close(send) }()
	return send
}

func (S *String) SelRecv() <-chan string {
	recv := make(chan string)
	go func() { recv <- S.From(); close(recv) }()
	return recv
}

func (S *String) To(s string) {
	go started.Do(getAndOrPostIfServer)

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string]*String{}
	}
	if _, found := stringDict.m[S.Name]; !found {
		stringDict.m[S.Name] = S
	}
	stringDict.Unlock()

	S.p.w.Do(S.makeW)
	if IsClient {
		S.p.w.c <- []byte(s)
		return
	}

	S.p.n.Do(S.makeN)
	for {
		<-S.p.n.c
		S.p.w.c <- []byte(s)
		if len(S.p.n.c) == 0 {
			break
		}
	}
}

func (S *String) From() string {
	go started.Do(getAndOrPostIfServer)

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string]*String{}
	}
	if _, found := stringDict.m[S.Name]; !found {
		stringDict.m[S.Name] = S
	}
	stringDict.Unlock()

	S.p.r.Do(S.makeR)
	return string(<-S.p.r.c)
}
