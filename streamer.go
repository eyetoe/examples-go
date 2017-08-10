package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}
	fmt.Fprintln(w, "Run the script:")
	//	cmd := exec.Command("ls")
	cmd := exec.Command("/Users/john/bin/streamer.sh")
	cmd.Stdout = &fw
	cmd.Stderr = &fw
	cmd.Run()
}

func handle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "sending first line of data\n")
	if f, ok := res.(http.Flusher); ok {
		f.Flush()
	} else {
		log.Println("Damn, no flush\n")
	}
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(res, "sending another line of data\n")
	}
}

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
