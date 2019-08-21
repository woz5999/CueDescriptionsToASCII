package converter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/woz5999/CueDescriptionsToASCII/cues"
	"golang.org/x/net/context"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

func init() {
	dc := DescriptionConverter{}
	http.HandleFunc(dc.GetPath(), dc.WebPost)
}

//DescriptionConverter ... detault constructor
type DescriptionConverter struct{}

//ConvertDescriptions ... Main method
func (dc DescriptionConverter) ConvertDescriptions(
	file multipart.File,
	filename string,
	ctx context.Context) (string, error) {

	var err error
	fmt.Println("Getting cues")
	log.Infof(ctx, "Getting cues")
	cues, err := getCues(file, ctx)
	if err != nil {
		return "", err
	}
	fmt.Println("Converting cues")
	log.Infof(ctx, "Converting cues")
	ascii, err := cues.ConvertCues()
	if err != nil {
		return "", err
	}
	fmt.Println("Writing cues")
	log.Infof(ctx, "Writing cues")
	filename, err = writeCues(ascii, filename, ctx)
	if err != nil {
		return "", err
	}

	return filename, err
}

func getCues(file multipart.File, ctx context.Context) (cues.CueList, error) {
	var err error
	cueList := cues.CueList{}

	if file != nil {
		reader := csv.NewReader(file)
		reader.Comma = ','

		//iterate through csv file
		//detect the header row and create template for mapping cue fields
		//for normal records, create & save a cue object
		tmpl := cues.CueTemplate{}
		lineCount := 0
		bTmplSet := false

		for {
			record, err := reader.Read()

			if err == nil {
				if bTmplSet != true {
					//load the fields into a map in order to do column detection
					set := make(map[string]bool)
					for _, v := range record {
						key := strings.ToLower(v)
						set[key] = true
					}

					//attempt to detect column headers
					_, cueHeader := set["cue"]
					if cueHeader {
						tmpl, err = tmpl.Create(record)

						if err != nil {
							fmt.Println("Error creating cue template: " + err.Error())
							log.Infof(ctx, "Error creating cue template: "+err.Error())
							break
						} else {
							fmt.Println("Cue template created")
							log.Infof(ctx, "Cue template created")
							bTmplSet = true
						}
					}
				} else {
					//ignore items encountered before template is created
					if &tmpl != nil {
						cue := cues.Cue{Record: record, Template: tmpl}
						cueList.Cues = append(cueList.Cues, cue)
					}
				}
			} else {
				if err == io.EOF {
					break
				} else {
					break
				}
			}
			lineCount++
		}
	} else {
		err = errors.New("No file specified")
	}

	return cueList, err
}

func writeCues(
	output string,
	filename string,
	ctx context.Context) (string, error) {
	var err error

	fileURL := ""

	// get default bucket name
	bucket, err := file.DefaultBucketName(ctx)
	if err == nil {
		//convert filename to output filename
		filenameSplit := strings.Split(filename, ".")
		filenameOut := filenameSplit[0] + ".asc"
		fileURL = "http://" + bucket + ".storage.googleapis.com/" + filenameOut

		fmt.Println("Writing ascii file " + filenameOut)
		log.Infof(ctx, "Writing ascii file "+bucket+": "+filenameOut)

		// create storage client
		client, err := storage.NewClient(ctx)
		if err == nil {
			defer client.Close()

			// get bucket
			b := client.Bucket(bucket)

			wc := b.Object(filenameOut).NewWriter(ctx)
			wc.ContentType = "text/plain"

			_, err := wc.Write([]byte(output))
			if err == nil {
				err := wc.Close()
				if err != nil {
					log.Errorf(ctx, "Unable to close bucket "+bucket+
						" "+err.Error())
				}
			} else {
				log.Errorf(ctx, "Unable to write data to bucket "+bucket+
					" "+err.Error())
			}
		} else {
			log.Errorf(ctx, "failed to create client: "+err.Error())
		}
	} else {
		log.Errorf(ctx, "Failed to get default GCS bucket name: "+
			err.Error())
	}

	return fileURL, err
}
