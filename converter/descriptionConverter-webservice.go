package converter

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
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

	output, err := dc.ConvertDescriptions(file)
	if err != nil {
		log.Println(reqInfo+" Response: ", http.StatusInternalServerError,
			" "+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// update the file extension
	filename := strings.Split(handler.Filename, ".")[0] + ".asc"

	fmt.Println(reqInfo+" Response: ", http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	content := bytes.NewReader([]byte(output))

	http.ServeContent(w, req, filename, time.Now(), content)
}
