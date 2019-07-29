package users

import (
	"log"

	"database/sql"
	"github.com/bagusandrian/testingwebapp/src/config"
	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	"html/template"
	"net/http"
)

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

type Module struct {
	Conf     *config.Config
	queries  *Queries
	DBMaster *sql.DB
	DBSlave  *sql.DB
}

func NewModule(c *config.Config) *Module {

	m := &Module{
		Conf: c,
	}

	dbMaster, err := sql.Open("mysql", c.Database.DBMaster)
	if dbMaster == nil {
		log.Println("[ERROR] connecting to DBMaster err:", err)
		return nil
	}
	m.DBMaster = dbMaster

	m.DBSlave, err = sql.Open("mysql", c.Database.DBSlave)
	if m.DBSlave == nil {
		log.Panic("[ERROR] connecting to DBSlave err:", err)
		return nil
	}

	m.queries = NewQueries(m.DBMaster, m.DBSlave)

	return m
}

func RegisterRouters(mdle *Module) {
	http.HandleFunc("/ping", mdle.ping)
	http.HandleFunc("/login", mdle.HandlerLoginRender)
}

func (m *Module) ping(w http.ResponseWriter, r *http.Request) {
	daniel := struct {
		Name   string
		Age    int
		Alamat string
	}{"Daniel Sudibyo", 23, "BSD"}
	tmplt, err := template.New("index").Parse("Nama saya {{.Name}} dan umur saya {{.Age}} tahun tinggal di {{.Alamat}}")
	if err != nil {
		log.Fatal(err)
	}

	err = tmplt.Execute(w, daniel) // send data to client side
	if err != nil {
		log.Fatal(err)
	}
}
