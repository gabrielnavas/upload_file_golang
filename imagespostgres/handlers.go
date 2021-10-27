package imagespostgres

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func UploadImage(db *sql.DB) http.HandlerFunc {
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

		// create model
		product := NewProduct(header.Filename)
		if err := product.SetFile(fileIn); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Fatalln(err)
			return
		}

		// insert into database
		repo := &RepositoryPostgres{db}
		id, err := repo.Store(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("server error"))
			log.Fatalln(err)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(id)))
	}
}

func DownloadImage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// just like GET Method
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// get ID from query param
		// http://localhost:8000/postgres/download?id=8
		query := r.URL.Query()
		id := query.Get("id")
		if id == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("need query param id"))
			return
		}
		productID, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("only number on query param id"))
			return
		}

		// get file from repository
		repo := &RepositoryPostgres{db}
		product, err := repo.Get(productID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("server error"))
			log.Fatalln(err)
			return
		}

		// download file
		w.Write(product.File.Bytes())
	}
}
