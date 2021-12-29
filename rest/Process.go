package rest

import (
	"fmt"
	"net/http"
)

func Process(rw http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) //max 10MB
	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
}
