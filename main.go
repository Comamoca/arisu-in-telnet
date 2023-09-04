package main

import (
	"github.com/reiver/go-telnet"
	"time"
)

type handler struct{}

func (h handler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	println("-- connect --")

	w.Write([]byte("\x1b[2J"))

	title := AA()
	info := Info()

	render(w, title)
	render(w, info)

	time.Sleep(180 * time.Second)
	w.Write([]byte("タイムアウト"))

	println("-- disconnect --")
}

func render(w telnet.Writer, text string) {
	for _, c := range text {
		w.Write([]byte(string(c)))
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	addr := ":5555"
	if err := telnet.ListenAndServe(addr, handler{}); nil != err {
		panic(err)
	}
}
