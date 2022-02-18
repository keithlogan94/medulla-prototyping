package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	var isCss bool = r.URL.Path[len(r.URL.Path)-4:] == ".css"
	var isJs bool = r.URL.Path[len(r.URL.Path)-3:] == ".js"
	if basename == "index.html" || basename == "blazor.boot.json" || isCss || isJs {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(Get(url)))
		return
	} else {
		// redirect every other url
		http.Redirect(w, r, url, 301)
	}

}

func Serve() {
	fmt.Println("Running server version 1.03")
	http.HandleFunc("/", proxy)

	fmt.Printf("Starting server on port :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting up server")
		log.Fatal(err)
	}
}
