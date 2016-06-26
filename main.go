package main

import (
	"converter"
    // "log "
	// "fmt"
	"net/http"
)

// app engine wrapper
func init() {
    main()
}

// documentation for csv is at http://golang.org/pkg/encoding/csv/
func main() {
	dc := converter.DescriptionConverter{}
	// port := ":8080"

	http.HandleFunc(dc.GetPath(), dc.WebPost)
	http.HandleFunc("/", home)
	http.HandleFunc("/style.css", style)

	// fmt.Println("Server started on localhost" + port)
	// log.Fatal(http.ListenAndServe(port, nil))
}

func home(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./converter.html")
}

func style(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./style.css")
}
