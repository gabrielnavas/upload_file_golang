package main

import (
	"log"
	"net/http"
	"server/database"
	"server/imagesfile"
	"server/imagespostgres"
)

func main() {
	db := database.NewDatabase()
	defer db.Close()

	router := http.NewServeMux()
	router.HandleFunc("/file/upload", imagesfile.UploadImage())
	router.HandleFunc("/file/download", imagesfile.DownloadImage())
	router.HandleFunc("/postgres/upload", imagespostgres.UploadImage(db))
	router.HandleFunc("/postgres/download", imagespostgres.DownloadImage(db))

	log.Print("Server started!")
	log.Print(http.ListenAndServe(":8000", router))
}
