// markdownServer.go

//   (c) 2012-2013 David Rook  ( hotei1352@gmail.com )
//		serves markdown docs @ "http://127.0.0.1:8080/md/"
//		md is a virtual URL mapped onto /www
//		status - working
package main

import (
	// Alien imports
	"github.com/blackfriday" // Russ Ross 2012-11-22 version
	// below refers only to go 1.0.3 standard lib pkg
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const portNum = 8080
const virtualURL = "/md/"
const serverRoot = "/www/"

var mdDir = []byte(`
<h3>Cant find index.html</h3>
<a href="./MarkDownTest.md">My BlackFriday Test</a><br>
<a href="./markdown-syntax.md">Markdown Syntax</a><br>
<!-- <a href="./walker.md">Walker</a><br>
-->
`)

var portNumString = fmt.Sprintf(":%d", portNum)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	var output []byte
	var err error
	fmt.Fprintf(w, "<!-- %s %v -->", r.Method, r.URL) // debug request input
	if len(r.URL.Path) == len(virtualURL) {
		// browse directory via index.html, don't allow 'raw' directory
		ndxBytes, err := ioutil.ReadFile(serverRoot + "index.html")
		if err != nil {
			w.Write(mdDir)
			return
		}
		w.Write(ndxBytes)
		return
	}
	// if tail of URL.Path == ".md" then use htmlFromMd, else just output it
	urlOffset := len(virtualURL)
	fileName := r.URL.Path[urlOffset:]
	ext := filepath.Ext(fileName)
	if ext == ".md" || ext == ".markdown" || ext == ".mdown" {
		output = htmlFromMd(serverRoot + fileName)
	} else {
		output, err = ioutil.ReadFile(serverRoot + fileName)
		if err != nil {
			fmt.Printf("error %v\n", err)
			return
		}
	}
	w.Write(output)
	//fmt.Fprintf(w, "<-- %q %q-->",fileName, ext)
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
	if false { // debug use only
		os.Stdout.Write(input)
		os.Stdout.Write(output)
	}
	return output
}

func main() {
	fmt.Printf("Starting markdown server at localhost:%s\n", portNumString)
	http.HandleFunc(virtualURL, WebHandler)
	http.Handle(serverRoot, http.StripPrefix(serverRoot, http.FileServer(http.Dir(serverRoot))))
	err := http.ListenAndServe(portNumString, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
