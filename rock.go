package rock

func (S *String) To(s string) {
	go getAndOrPostIfServer()

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string]*String{}
	}
	if _, found := stringDict.m[S.Name]; !found {
		stringDict.m[S.Name] = S
	}
	stringDict.Unlock()

	S.p.w.Lock()
	if S.p.w.c == nil {
		S.p.w.c = make(chan []byte, S.Len)
		go postIfClient(S.p.w.c, Tstring, S.Name)
	}
	S.p.w.Unlock()

	if IsClient {
		S.p.w.c <- []byte(s)
		return
	}

	S.p.n.Lock()
	if S.p.n.c == nil {
		S.p.n.c = make(chan int)
	}
	N := S.p.n.c
	S.p.n.Unlock()

	for {
		<-N
		S.p.w.c <- []byte(s)
		if len(N) == 0 {
			break
		}
	}
}

func (S *String) From() string {
	go getAndOrPostIfServer()

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string]*String{}
	}
	if _, found := stringDict.m[S.Name]; !found {
		stringDict.m[S.Name] = S
	}
	stringDict.Unlock()

	S.p.w.Lock()
	if S.p.r.c == nil {
		S.p.r.c = make(chan []byte, S.Len)
		go getIfClient(S.p.r.c, Tstring, S.Name)
	}
	S.p.w.Unlock()

	return string(<-S.p.r.c)
}
