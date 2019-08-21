package main

import (
  "log "
	"fmt"
	"net/http"

	"github.com/woz5999/CueDescriptionsToASCII/converter"
	"github.com/woz5999/CueDescriptionsToASCII/static"
	"google.golang.org/appengine"
)

func main() {
	appengine.Main()

// app engine wrapper
func init() {
    main()
}

// func main() {
// 	dc := converter.DescriptionConverter{}
// 	// port := ":8080"

// 	http.HandleFunc(dc.GetPath(), dc.WebPost)
// 	http.HandleFunc("/", home)
// 	http.HandleFunc("/style.css", style)

// 	// fmt.Println("Server started on localhost" + port)
// 	// log.Fatal(http.ListenAndServe(port, nil))
// }

func home(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./converter.html")
}

func style(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./style.css")
}
