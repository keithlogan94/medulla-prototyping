package main

import (
	"fmt"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
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

	http.Redirect(w, r, url, 301)
}

func Serve() {
	http.HandleFunc("/", redirect)

	fmt.Printf("Starting server on port :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting up server")
		log.Fatal(err)
	}
}
