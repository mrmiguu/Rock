package rock

import "errors"

func (E *Error) makeW() {
	E.p.w.c = make(chan []byte, E.Len)
	go postIfClient(E.p.w.c, Terror, E.Name)
}

func (E *Error) makeR() {
	E.p.r.c = make(chan []byte, E.Len)
	go getIfClient(E.p.r.c, Terror, E.Name)
}

func (E *Error) makeNIfServer() {
	if IsClient {
		return
	}
	E.p.n.c = make(chan int)
}

func (E *Error) add() {
	errorDict.Lock()
	if errorDict.m == nil {
		errorDict.m = map[string]*Error{}
	}
	if _, found := errorDict.m[E.Name]; !found {
		errorDict.m[E.Name] = E
	}
	errorDict.Unlock()
}

func (E *Error) to(e error) {
	if IsClient {
		E.p.w.c <- []byte(e.Error())
		return
	}
	for {
		<-E.p.n.c
		E.p.w.c <- []byte(e.Error())
		if len(E.p.n.c) == 0 {
			break
		}
	}
}

func (E *Error) from() error {
	return errors.New(string(<-E.p.r.c))
}

func (E *Error) S() chan<- error {
	c := make(chan error, E.Len)
	go started.Do(getAndOrPostIfServer)
	E.add()
	E.p.w.Do(E.makeW)
	E.p.n.Do(E.makeNIfServer)
	go func() {
		E.to(errors.New(""))
		i := <-c
		close(c)
		E.to(i)
	}()
	return c
}

func (E *Error) R() <-chan error {
	c := make(chan error, E.Len)
	go started.Do(getAndOrPostIfServer)
	E.add()
	E.p.r.Do(E.makeR)
	go func() {
		E.from()
		c <- E.from()
		close(c)
	}()
	return c
}
