package main

import (
	// "html/template"
	"log"
	"net/http"

	"github.com/bagusandrian/testingwebapp/src/config"
	"github.com/bagusandrian/testingwebapp/src/users"
)

var conf *config.Config

func init() {
	conf = config.ReadConfig()

}

func main() {
	log.Println("running on http://localhost", conf.Server.Port)
	mdlUsers := users.NewModule(conf)
	users.RegisterRouters(mdlUsers)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("files/www/html"))))
	err := http.ListenAndServe(conf.Server.Port, nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
