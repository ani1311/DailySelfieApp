package main

import (
	"io"
	"net/http"

	"fortio.org/log"
)

func uploadhander(w http.ResponseWriter, r *http.Request) {
	log.Infof("Received upload request")

	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // limit your maxMultipartMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the file from form data.
	file, _, err := r.FormFile("image") // image is the key of the form data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	// For this example, we're just dumping the file into the server's local storage.
	// You might want to do something more interesting here, like send the file to a
	// remote server, save it to a database, or process it in some way.
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Infof("Received file with length: %d", len(fileBytes))

	err = UploadFile(fileBytes)
	if err != nil {
		log.Errf("Error uploading file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
