package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"

	"header-shield/pkg/scanner"
)

// StartServer inicia el servidor HTTP local recibiendo el sistema de archivos embebido
func StartServer(port int, webFS embed.FS) error {
	// API de escaneo
	http.HandleFunc("/api/scan", handleScanAPI)

	// Extraer la subcarpeta "web" para servir los archivos estáticos (HTML, CSS, JS, imágenes)
	webSubFS, err := fs.Sub(webFS, "web")
	if err != nil {
		return fmt.Errorf("error al obtener subcarpeta web: %w", err)
	}

	// Servidor de archivos estáticos automático (sirve index.html, styles.css, etc.)
	fileServer := http.FileServer(http.FS(webSubFS))
	http.Handle("/", fileServer)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("🌐 Servidor Web iniciado en: http://localhost%s\n", addr)
	fmt.Println("Presiona Ctrl+C para detener el servidor")

	return http.ListenAndServe(addr, nil)
}

func handleScanAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, `{"error":"El parámetro 'domain' es obligatorio"}`, http.StatusBadRequest)
		return
	}

	headers, serverHeader, err := scanner.CheckHeaders(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	sslResult := scanner.CheckSSL(domain)

	result := scanner.ScanResult{
		Target:       domain,
		Headers:      headers,
		ServerHeader: serverHeader,
		SSL:          sslResult,
	}

	json.NewEncoder(w).Encode(result)
}