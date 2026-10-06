package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	// importa tus otros paquetes (net, crypto/tls, etc.)
)

//go:embed web/*
var webFS embed.FS

func main() {
	// Extraer el subdirectorio "web" para servir el index.html
	contentStatic, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatal(err)
	}

	// Rutas HTTP
	http.Handle("/", http.FileServer(http.FS(contentStatic)))
	http.HandleFunc("/api/scan", handleScan) // Tu función que realiza el escaneo

	fmt.Println("🚀 Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	// Aquí va la lógica de tu escáner que devuelve el JSON
}