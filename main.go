package main

import (
	//"converter"
	"fmt"
	"github.com/woz5999/descriptionConverter/converter"
	"log"
	"net/http"
)

// documentation for csv is at http://golang.org/pkg/encoding/csv/
func main() {
	dc := converter.DescriptionConverter{}
	port := ":3000"

	http.HandleFunc(dc.GetPath(), dc.WebPost)

	fmt.Println("Server started on localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
