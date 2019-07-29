package users

import (
	"html/template"
	"log"
	"net/http"
)

func (m *Module) HandlerLoginRender(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("files/www/html/login/index.html")
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
