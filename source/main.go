package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", ServeIndex)
	http.HandleFunc("/Fase1", ServeFase1)
	http.HandleFunc("/Fase2", ServeFase1)
	http.HandleFunc("/Fase3", ServeFase1)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Servidor en: http://127.0.0.1:8080")
	}

}
