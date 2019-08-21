package static

import "net/http"

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/style.css", style)
}

func home(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./static/files/converter.html")
}

func style(w http.ResponseWriter,
	req *http.Request) {
	http.ServeFile(w, req, "./static/files/style.css")
}
