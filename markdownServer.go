// markdown_server.go
package main

/*
 * (c) 2012 David Rook  ( hotei1352@gmail.com )
 *
 * to serve markdown docs, browse to "http://127.0.0.1:8080/www/docname.md"
 *
 * status - working
 * 
 */

import (
	"fmt"
	"github.com/blackfriday"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const portNum = 8080
const virtualURL = "/markdown/"
const serverRoot = "/www/"

var portNumString = fmt.Sprintf(":%d", portNum)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<!-- %s %v -->", r.Method, r.URL) debug request input
	urlOffset := len(virtualURL)
	output := htmlFromMd("/www/" + r.URL.Path[urlOffset:])
	w.Write(output)
}

func htmlFromMd(fname string) []byte {
	var output []byte
	input, err := ioutil.ReadFile(fname)
	if err != nil {
		tmp := fmt.Sprintf("Problem reading input, can't open %s", fname)
		output = []byte(tmp)
	} else {
		output = blackfriday.MarkdownCommon(input)
	}
	if false {
		os.Stdout.Write(input)
		os.Stdout.Write(output)
	}
	return output
}

func main() {
	fmt.Printf("Starting markdown server at localhost:%s\n", portNumString)
	http.HandleFunc(virtualURL, HelloHandler)
	http.Handle(serverRoot, http.StripPrefix(serverRoot, http.FileServer(http.Dir(serverRoot))))
	err := http.ListenAndServe(portNumString, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
