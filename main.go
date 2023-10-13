package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amsem/url-shortner/controler"
	"github.com/amsem/url-shortner/helper"
	"github.com/gorilla/mux"
)



func main()  {
    db, err := helper.InitDB()
    
    if err != nil {
        panic(err)
    }
    dbCLIENT := &controler.DBclient{DB: db}
    if err != nil {
        panic(err)
    }
    defer db.Close()
    r := mux.NewRouter()
    r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbCLIENT.GetOriginalURL).Methods("GET")
    r.HandleFunc("/v1/short", dbCLIENT.GenerateShortURL).Methods("POST")
    srv := &http.Server{
    Handler: r,
    Addr: "127.0.0.1:8000",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
    }
    log.Fatal(srv.ListenAndServe())
    

}
