package converter

import (
	"fmt"
	"net/http"
    "google.golang.org/appengine"
    "google.golang.org/appengine/log"
)

//GetPath ... return the url path for this web service
func (dc DescriptionConverter) GetPath() string {
	return "/converter"
}

//WebPost ... handler for POST event
func (dc DescriptionConverter) WebPost(w http.ResponseWriter,
	req *http.Request) {
	defer req.Body.Close()
	var err error

	reqInfo := req.RemoteAddr + " " + req.Method

	if req.Method != "POST" {
		log.Println(reqInfo+" Response: ", http.StatusMethodNotAllowed,
			" "+err.Error())
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	file, handler, err := req.FormFile("uploadfile")
	if err != nil {
		log.Println(reqInfo+" Response: ", http.StatusBadRequest,
			" "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//do file conversion
	filename, err := dc.ConvertDescriptions(file, handler.Filename)
	if err != nil {
		log.Println(reqInfo+" Response: ", http.StatusInternalServerError,
			" "+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(reqInfo+" Response: ", http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	http.ServeFile(w, req, filename)

	os.Remove(filename)
	return
}
