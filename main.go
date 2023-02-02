package main

import (
	"CMS/app/API"
	"CMS/app/DB"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {

	//db.go init()에서 DBOpen 진행

	// Set up the database
	DB.Init()
	defer DB.CloseDB()
	API.Init()

	// Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Initalize an HTTP request multiplexer

	// set Handle
	//현재는 Main페이지에서 ADDSERVER, ALIVE, JOB등의 파라메터로 페이지를 나눔.
	//따라서 Main에서 해당 파라메터가 있는지 없는지, if로 구현이 되고, 이걸 redirect시키는걸로 변경해야함.
	mux.HandleFunc("/Community/CheckStart", func(res http.ResponseWriter, req *http.Request) {
		API.AddContent(res, req)
	})

	handler := cors.Default().Handler(mux)
	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:9070", "*:9070"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           0,
		Debug:            false,
	})

	handler = corHandler.Handler(handler)
	//logRequest(handler)

	http.ListenAndServe(":9070", handler)
}
