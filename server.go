package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

func Get(url string) string {
	fmt.Println("Sending request to " + url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func threadConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling request")
	proxy(w, r)
}

func DetermineMimeType(filePath string) string {
	fileExtension := filepath.Ext(filePath)
	return mime.TypeByExtension(fileExtension)
}

func proxy(w http.ResponseWriter, r *http.Request) {
	// check
	switch r.Method {
	case "GET":
		break
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("Only GET methods are supported"))
		return
	}

	var s3ObjectPath string = r.URL.Path
	url, err := CreatePresignedUrl(s3ObjectPath)
	if err != nil {
		log.Fatalln(err)
	}

	index := strings.LastIndex(r.URL.Path, "/")
	fmt.Println("index is " + string(index))
	var basename string = r.URL.Path[index+1:]
	fmt.Println("basename: " + basename)
	w.Header().Set("Content-Type", DetermineMimeType(basename))
	w.Write([]byte(Get(url)))
	return
}

func Serve() {
	fmt.Println("Running server version 1.03")
	http.HandleFunc("/", threadConnection)

	fmt.Printf("Starting server on port :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting up server")
		log.Fatal(err)
	}
}
