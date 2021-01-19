package main

import (
	"fmt"
	"net/http"
	"text/template"

	_ "gorm.io/driver/sqlserver"
	_ "gorm.io/gorm"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println(produtos)
	temp.ExecuteTemplate(w, "index", produtos)
}
