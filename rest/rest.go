package rest

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gopherjs/gopherjs/js"
)

func (S *String) Post(s string) {
	getAndOrPostIfServer()
	if _, found := stringDict[S.Name]; !found {
		stringDict[S.Name] = S
	}
	if S.w == nil {
		S.w = make(chan []byte, S.Len)
		go postIfClient(S.w, Tstring, S.Name)
	}
	S.w <- []byte(s)
}

func (S *String) Get() string {
	getAndOrPostIfServer()
	if _, found := stringDict[S.Name]; !found {
		stringDict[S.Name] = S
	}
	if S.r == nil {
		S.r = make(chan []byte, S.Len)
		go getIfClient(S.r, Tstring, S.Name)
	}
	return string(<-S.r)
}

const (
	DefaultClientAddr = "/"
	DefaultServerAddr = ":80"
	V                 = "▼"
	POST              = "9b466094ec991a03cb95c489c19c4d75635f0ae5"
	GET               = "783923e57ba5e8f1044632c31fd806ee24814bb5"

	Tstring byte = iota
)

var (
	Addr     string
	IsClient = js.Global != nil

	v       = []byte(V)
	started bool

	stringDict = map[string]*String{}
)

func init() {
	if IsClient {
		Addr = DefaultClientAddr
	} else {
		Addr = DefaultServerAddr
	}
}

func getAndOrPostIfServer() {
	if started || IsClient {
		return
	}

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
			stringDict[name].r <- body
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
			b = <-stringDict[name].w
		default:
			panic("bad message type")
		}
		w.Write(b)
	})

	go func() { log.Fatal(http.ListenAndServe(Addr, nil)) }()
	started = true
}

func postIfClient(w chan []byte, t byte, name string) {
	if !IsClient {
		return
	}
	if Addr[len(Addr)-1] != '/' {
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
	if Addr[len(Addr)-1] != '/' {
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

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func must(e error) {
	if e == nil {
		return
	}
	panic(e)
}
