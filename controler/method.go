package controler

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/amsem/url-shortner/utils"
	"github.com/gorilla/mux"
)


type DBclient struct {
    DB *sql.DB
}

type Record struct {
    ID int `json:"id"`
    URL string `json:"url"`
}


func (driver *DBclient) GetOriginalURL(w http.ResponseWriter, r *http.Request)  {
    var url string
    vars := mux.Vars(r)
    id := utils.ToDeCode(vars["encoded_string"])
    err := driver.DB.QueryRow("SELECT url FROM  web_url WHERE id = $1", id).Scan(&url)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
    }else {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        responseMap := map[string]interface{}{"url": url}
        response, _ := json.Marshal(responseMap)
        w.Write(response)
    }
}

func (driver *DBclient) GenerateShortURL(w http.ResponseWriter, r *http.Request)  {
    var id int
    var record Record
    postBody, _ := io.ReadAll(r.Body)
    err := json.Unmarshal(postBody, &record)
    if err != nil {
        panic(err)
    }
    err = driver.DB.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.URL).Scan(&id)
    responseMap := map[string]string{"encoded_string": utils.ToEncode(id)}
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
    }else {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        response, _ := json.Marshal(responseMap)
        w.Write(response)
    }
}

func (driver *DBclient) DeleteShortURL(w http.ResponseWriter, r *http.Request)  {
    var record Record
    postBody, _ := io.ReadAll(r.Body)
    err := json.Unmarshal(postBody, &record)
    
    if err != nil {
        panic(err)
    }
    log.Println("Im here 3333")
    err = driver.DB.QueryRow("DELETE FROM web_url * WHERE id = ($1)", record.ID).Err()
  if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
    }else {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("deleted"))
    }
    log.Println("Im here 1")
}
  
