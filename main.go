package main

import (
	"fmt"
	"net/http"

	//"github.com/qor/qor"
	//_ "CMS/config/db/migrations"

	//"log"
	//"github.com/qor/admin"
	"CMS/app/API"
	"CMS/config/db"
	"CMS/models/content"
)

func main() {

	//db.go init()에서 DBOpen 진행

	// Set up the database
	db.DB.AutoMigrate(&content.CMS_Content{}, &content.CMS_Video{}, &content.CMS_Audio{}, &content.CMS_Image{}, &content.Map_Video_Extension{})

	// Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// set Handle
	mux.HandleFunc("/API/AddContent", func(res http.ResponseWriter, req *http.Request) {
		API.AddContent(res, req)
	})
	mux.HandleFunc("/API/CompleteTranscoding", func(res http.ResponseWriter, req *http.Request) {
		API.CompleteTranscoding(res, req)
	})
	mux.HandleFunc("/API/StartHawkeye", func(res http.ResponseWriter, req *http.Request) {
		API.StartHawkeye(res, req)
	})
	mux.HandleFunc("/API/CompleteHawkeye", func(res http.ResponseWriter, req *http.Request) {
		API.CompleteHawkeye(res, req)
	})
	mux.HandleFunc("/API/UpdateMediainfo", func(res http.ResponseWriter, req *http.Request) {
		API.UpdateMediainfo(res, req)
	})

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
