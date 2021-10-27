package imagesfile

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

const nameTemp = "photofile.jpg"

func UploadImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// just like POST Method
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// get file from formFile
		fileIn, header, err := r.FormFile("photo") // name key from formFile
		defer fileIn.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Fatalln(err)
			return
		}

		// info files
		fmt.Println(header.Filename)      // filename
		fmt.Println(header.Header)        // meta data
		fmt.Println(float64(header.Size)) // len in bytes

		if err := validateFileName(header.Filename); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(err.Error()))
			log.Fatalln(err)
			return
		}

		// open and create file
		fileOut, err := os.Create(nameTemp)
		defer fileOut.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("server error"))
			log.Fatalln(err)
			return
		}

		// copy to file created
		_, err = io.Copy(fileOut, fileIn)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("server error"))
			log.Fatalln(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func DownloadImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// just like GET Method
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// read bytes from file
		bytesFromFile, err := os.ReadFile(nameTemp)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write(bytesFromFile)
	}
}

func validateFileName(filename string) error {
	match, _ := regexp.MatchString("^[a-z|A-Z]+$", filename)
	if !match {
		return errors.New("only letters (upper or lower case) are allowed")
	}
	return nil
}
