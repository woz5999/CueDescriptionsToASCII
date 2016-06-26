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
    ctx := appengine.NewContext(req)
	defer req.Body.Close()
	var err error

	reqInfo := req.RemoteAddr + " " + req.Method

	if req.Method != "POST" {
        log.Errorf(ctx, reqInfo+" Response: ", http.StatusMethodNotAllowed,
			" "+err.Error())
		fmt.Println(reqInfo+" Response: ", http.StatusMethodNotAllowed,
			" "+err.Error())
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	file, handler, err := req.FormFile("uploadfile")
	if err != nil {
        log.Errorf(ctx, reqInfo+" Response: ", http.StatusBadRequest,
            " "+err.Error())
		fmt.Println(reqInfo+" Response: ", http.StatusBadRequest,
		    " "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//do file conversion
	filename, err := dc.ConvertDescriptions(file, handler.Filename, ctx)
	if err != nil {
        log.Errorf(ctx, reqInfo+" Response: ", http.StatusInternalServerError,
            " "+err.Error())
		fmt.Println(reqInfo+" Response: ", http.StatusInternalServerError,
		          " "+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

    log.Infof(ctx, reqInfo+" Response: ", http.StatusMovedPermanently)
	fmt.Println(reqInfo+" Response: ", http.StatusMovedPermanently)

    http.Redirect(w, req, filename, 301)

	return
}
