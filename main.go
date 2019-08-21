package main

import (
	"net/http"

	_ "github.com/woz5999/CueDescriptionsToASCII/converter"
	_ "github.com/woz5999/CueDescriptionsToASCII/static"
	"google.golang.org/appengine"
)

func main() {
	appengine.Main()
}

func home(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./converter.html")
}

func style(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./style.css")
}
