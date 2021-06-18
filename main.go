package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading file\n")

	// 1. Parse input, type multi-part/form data e.g 10 MB data is being used here
	r.ParseMultipartForm(10 << 20)

	// 2. retrieve file from posted form data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %v\n", handler.Filename)
	fmt.Printf("File size: %v\n", handler.Size)
	fmt.Printf("MIME Header: %v\n", handler.Header)

	// 3. write temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	// 4. return whether or not this has been successful
	fmt.Fprintf(w, "Successfully uploaded file\n")
}

// func showFile(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Here is your piccy you uploaded\n")

// }

func setUpRoutes() {
	http.HandleFunc("/upload", uploadFile)
	// http.HandleFunc("/show", showFile)
	http.ListenAndServe(":9090", nil)
}

func main() {
	fmt.Println("Go file upload tutorial")
	setUpRoutes()
}
