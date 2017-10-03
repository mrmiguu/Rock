package rock

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init() {
	if IsClient {
		Addr = DefaultClientAddr
	} else {
		Addr = DefaultServerAddr
	}
}

func getAndOrPostIfServer() {
	if IsClient {
		return
	}

	// consider commenting out? idk
	http.Handle("/", http.FileServer(http.Dir("www")))

	http.HandleFunc("/"+POST, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			delayedError(w, http.StatusBadRequest)
			return
		}
		parts := bytes.Split(b, v)
		t, name, body := parts[0][0], string(parts[1]), parts[2]

		switch t {
		case Terror:
			errorDict.RLock()
			E, found := errorDict.m[name]
			errorDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			E.p.r.Do(E.makeR)
			E.p.r.c <- body
		case Tbool:
			boolDict.RLock()
			B, found := boolDict.m[name]
			boolDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.p.r.Do(B.makeR)
			B.p.r.c <- body
		case Tint:
			intDict.RLock()
			I, found := intDict.m[name]
			intDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.p.r.Do(I.makeR)
			I.p.r.c <- body
		case Tstring:
			stringDict.RLock()
			S, found := stringDict.m[name]
			stringDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			S.p.r.Do(S.makeR)
			S.p.r.c <- body
		case Tbytes:
			bytesDict.RLock()
			B, found := bytesDict.m[name]
			bytesDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.p.r.Do(B.makeR)
			B.p.r.c <- body
		default:
			delayedError(w, http.StatusBadRequest)
		}
	})

	http.HandleFunc("/"+GET, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			delayedError(w, http.StatusBadRequest)
			return
		}
		parts := bytes.Split(b, v)
		t, name := parts[0][0], string(parts[1])

		switch t {
		case Terror:
			errorDict.RLock()
			E, found := errorDict.m[name]
			errorDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			E.p.n.Do(E.makeN)
			E.p.n.c <- 1
			E.p.w.Do(E.makeW)
			b = <-E.p.w.c
		case Tbool:
			boolDict.RLock()
			B, found := boolDict.m[name]
			boolDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.p.n.Do(B.makeN)
			B.p.n.c <- 1
			B.p.w.Do(B.makeW)
			b = <-B.p.w.c
		case Tint:
			intDict.RLock()
			I, found := intDict.m[name]
			intDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			I.p.n.Do(I.makeN)
			I.p.n.c <- 1
			I.p.w.Do(I.makeW)
			b = <-I.p.w.c
		case Tstring:
			stringDict.RLock()
			S, found := stringDict.m[name]
			stringDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			S.p.n.Do(S.makeN)
			S.p.n.c <- 1
			S.p.w.Do(S.makeW)
			b = <-S.p.w.c
		case Tbytes:
			bytesDict.RLock()
			B, found := bytesDict.m[name]
			bytesDict.RUnlock()
			if !found {
				delayedError(w, http.StatusNotFound)
				return
			}
			B.p.n.Do(B.makeN)
			B.p.n.c <- 1
			B.p.w.Do(B.makeW)
			b = <-B.p.w.c
		default:
			delayedError(w, http.StatusBadRequest)
			return
		}

		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(Addr, nil))
}

func delayedError(w http.ResponseWriter, code int) {
	time.Sleep(ErrorDelay)
	http.Error(w, "", code)
}

func postIfClient(w chan []byte, t byte, name string) {
	if !IsClient {
		return
	}
	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}
	for {
		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name), <-w}, v)
		for {
			resp, err := http.Post(Addr+POST, "text/plain", bytes.NewReader(pkt))
			if err == nil && resp.StatusCode < 300 {
				break
			}
		}
	}
}

func getIfClient(r chan []byte, t byte, name string) {
	if !IsClient {
		return
	}
	if len(Addr) == 0 || Addr[len(Addr)-1] != '/' {
		Addr += "/"
	}
	for {
		pkt := bytes.Join([][]byte{[]byte{t}, []byte(name)}, v)
		for {
			resp, err := http.Post(Addr+GET, "text/plain", bytes.NewReader(pkt))
			if err != nil || resp.StatusCode > 299 {
				continue
			}
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err == nil {
				r <- b
				break
			}
		}
	}
}
