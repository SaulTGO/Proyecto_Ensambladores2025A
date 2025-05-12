package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:embed templates/* assets/* validations/* references/*
var embeddedFiles embed.FS

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFS(embeddedFiles, "templates/*.html")
	if err != nil {
		log.Printf("Error al cargar templates: %v", err)
	}
}

func main() {
	log.Println("Iniciando servidor en http://127.0.0.1:9999")

	assetFS, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		log.Fatalf("Error accediendo a assets: %v", err)
	}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assetFS))))

	// Configurar rutas
	http.HandleFunc("/", ServeIndex)
	http.HandleFunc("/Fase1", ServeFase1)
	http.HandleFunc("/Fase2", ServeFase2)
	http.HandleFunc("/Fase3", ServeFase3)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr: "127.0.0.1:9999",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar servidor: %v", err)
		}
	}()
	log.Println("Servidor ejecutándose. Presiona Ctrl+C para detener.")

	<-done
	log.Println("Cerrando servidor...")
}
