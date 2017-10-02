package rock

func (S *String) To(s string) {
	getAndOrPostIfServer()

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string]*String{}
	}
	if _, found := stringDict.m[S.Name]; !found {
		stringDict.m[S.Name] = S
	}
	stringDict.Unlock()

	S.w.Lock()
	if S.w.c == nil {
		S.w.c = make(chan []byte, S.Len)
		go postIfClient(S.w.c, Tstring, S.Name)
	}
	S.w.Unlock()

	if IsClient {
		S.w.c <- []byte(s)
		return
	}

	S.n.Lock()
	if S.n.c == nil {
		S.n.c = make(chan int)
	}
	N := S.n.c
	S.n.Unlock()
	for {
		<-N
		S.w.c <- []byte(s)
		if len(N) == 0 {
			break
		}
	}
}

func (S *String) From() string {
	getAndOrPostIfServer()

	stringDict.Lock()
	if stringDict.m == nil {
		stringDict.m = map[string]*String{}
	}
	if _, found := stringDict.m[S.Name]; !found {
		stringDict.m[S.Name] = S
	}
	stringDict.Unlock()

	S.w.Lock()
	if S.r.c == nil {
		S.r.c = make(chan []byte, S.Len)
		go getIfClient(S.r.c, Tstring, S.Name)
	}
	S.w.Unlock()

	return string(<-S.r.c)
}
