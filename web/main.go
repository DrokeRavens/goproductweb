package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/productweb/service"

	_ "gorm.io/driver/sqlserver"
	_ "gorm.io/gorm"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produto := service.ProductService.Buscar(1)
	fmt.Println(produto)
	temp.ExecuteTemplate(w, "index", *produto)
}

func cadastrar(w http.ResponseWriter, r *http.Request) {
	temp.Execute(w, "cadastrar")
}
