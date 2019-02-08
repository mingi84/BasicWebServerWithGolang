package main

import (
	"fmt"
	"net/http"

	//"github.com/qor/qor"
	//_ "CMS/config/db/migrations"

	//"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/qor/admin"
	"CMS/app/API"
	"CMS/models/content"

)

func main() {
	
	// Set up the database
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", "root", "_media_", "127.0.0.1", "3306", "mteg_cms"))
	defer db.Close()
	if err == nil {
		} else {
			panic(err)
		}
	
	//db Migrate
	db.AutoMigrate(&content.CMS_Content{}, &content.CMS_Video{}, &content.CMS_Audio{}, &content.CMS_Image{}, &content.Map_Video_Extension{})
	db.Close()
	// Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	
	// set Handle
	mux.HandleFunc("/API/GetKey", func(res http.ResponseWriter, req *http.Request){
		API.GetKey(res, req)
	})
	mux.HandleFunc("/API/ContentCreate", func(res http.ResponseWriter, req *http.Request){
		API.ContentCreate(res, req)
	})
	mux.HandleFunc("/API/ContentRead", func(res http.ResponseWriter, req *http.Request){
		API.ContentRead(res, req)
	})
	mux.HandleFunc("/API/ContentUpdate", func(res http.ResponseWriter, req *http.Request){
		API.ContentUpdate(res, req)
	})
	mux.HandleFunc("/API/ContentDelete", func(res http.ResponseWriter, req *http.Request){
		API.ContentDelete(res, req)
	})


	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
