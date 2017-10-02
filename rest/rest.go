package rest

func (S *String) Post(s string) {
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

	S.w.c <- []byte(s)
}

func (S *String) Get() string {
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
