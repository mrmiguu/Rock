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
