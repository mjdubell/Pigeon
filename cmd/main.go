package main

import (
	"log"
	"net/http"

	"github.com/mjdubell/pigeon/pkg/onetimesecret"
)

const (
	staticDir = "web/static"
	port      = "9999"
)

func main() {

	db, _ := onetimesecret.NewDB()

	fs := http.FileServer(http.Dir(staticDir))

	router := onetimesecret.Router(db)

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", router)

	log.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

	defer db.Close()
}