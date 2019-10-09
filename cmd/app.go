package main

import (
	"encoding/json"
	"fmt"
	"github.com/claudioontheweb/url-shortener/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) Initialize(user,password,dbname,mysql_port string) {

	fmt.Println("initializing!!!")

	var err error
	a.DB, err = gorm.Open("mysql", user + ":" + password + "@tcp(mysql:"+ mysql_port +")/" + dbname + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	a.DB.AutoMigrate(&models.UrlShorten{})

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}
func (a *App) Run(server_port string) {
	fmt.Println("running...")
	log.Fatal(http.ListenAndServe(":" + server_port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(a.Router)))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/{code}", a.redirectUrlHandler).Methods("GET")
	a.Router.HandleFunc("/", a.createShortUrlHandler).Methods("POST")
}

func (a *App) redirectUrlHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	url, err := models.GetOriginalUrl(a.DB, params["code"])
	if err != nil {
		respondWithError(w, http.StatusNotFound, "URL not found")
	}
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func (a *App) createShortUrlHandler(w http.ResponseWriter, r *http.Request) {

	var url models.UrlShorten

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		panic(err)
	}

	shortUrl, err := models.CreateShortUrl(a.DB, url)
	if err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, shortUrl)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}