package rock

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	if IsClient {
		Addr = DefaultClientAddr
	} else {
		Addr = DefaultServerAddr
	}
}

func getAndOrPostIfServer() {
	started.Lock()
	defer started.Lock()
	if started.b || IsClient {
		return
	}

	// consider commenting out? idk
	http.Handle("/", http.FileServer(http.Dir("www")))

	http.HandleFunc("/"+POST, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		parts := bytes.Split(b, v)
		t, name, body := parts[0][0], string(parts[1]), parts[2]
		switch t {
		case Tstring:
			stringDict.RLock()
			S := stringDict.m[name]
			stringDict.RUnlock()
			S.w.RLock()
			C := S.r.c
			S.w.RUnlock()
			C <- body
		default:
			panic("bad message type")
		}
	})

	http.HandleFunc("/"+GET, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		parts := bytes.Split(b, v)
		t, name := parts[0][0], string(parts[1])
		switch t {
		case Tstring:
			stringDict.RLock()
			S := stringDict.m[name]
			stringDict.RUnlock()
			S.w.RLock()
			C := S.w.c
			S.w.RUnlock()
			S.n.RLock()
			N := S.n.c
			S.n.RUnlock()
			N <- 1
			b = <-C
		default:
			panic("bad message type")
		}
		w.Write(b)
	})

	go func() { log.Fatal(http.ListenAndServe(Addr, nil)) }()
	started.b = true
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
