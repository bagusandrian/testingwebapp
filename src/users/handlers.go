package users

import (
	session "github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var store = session.NewCookieStore([]byte("SESSION_KEY"))

func (m *Module) HandlerLoginRender(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("files/www/html/login/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func (m *Module) HandlerValidateLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if session.Values["username"] != "" {
		log.Println("sudah login")
	}
	username := rgxSQLInjectionChar.ReplaceAllString(r.FormValue("username"), "")
	password := rgxSQLInjectionChar.ReplaceAllString(r.FormValue("password"), "")
	users := Users{}
	result := ""
	err := m.queries.ValidateUser.QueryRow(username, password).Scan(&users.UserID, &users.Username)
	if err != nil {
		result = "Maaf password/username salah"
	} else {
		// Set some session values.
		session.Values["username"] = users.Username
		session.Values["user_id"] = users.UserID
		// Save it before we write to the response/return from the handler.
		session.Save(r, w)
	}
	data := struct {
		Title  string
		Result string
		Data   Users
	}{
		Title:  "My page",
		Result: result,
		Data:   users,
	}
	const tpl = `
        <!DOCTYPE html>
        <html>
            <body>
               <p> {{.Result}} {{.Data.Username}}</p>
            </body>
        </html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("index").Parse(tpl)
	check(err)

	err = t.Execute(w, data)
	check(err)
}
