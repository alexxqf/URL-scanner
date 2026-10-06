package scanner

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// TargetHeaders define las cabeceras a inspeccionar y su nivel de severidad si faltan.
var securityHeaders = map[string]struct {
	description string
	severity    string
}{
	"Content-Security-Policy":   {"Previene ataques XSS e inyecciones de datos", "HIGH"},
	"Strict-Transport-Security": {"Fuerza el uso exclusivo de conexiones HTTPS (HSTS)", "HIGH"},
	"X-Frame-Options":           {"Evita que el sitio sea embebido en un iframe (Clickjacking)", "MEDIUM"},
	"X-Content-Type-Options":    {"Evita la interpretación incorrecta de tipos MIME (nosniff)", "MEDIUM"},
	"Referrer-Policy":           {"Protege la privacidad controlando la cabecera Referer", "LOW"},
	"Permissions-Policy":        {"Restringe el acceso a APIs del navegador (cámara, micro)", "LOW"},
}

// CheckHeaders realiza la petición HTTP y analiza las cabeceras recibidas.
func CheckHeaders(target string) ([]HeaderCheck, string, error) {
	client := &http.Client{
		Timeout: 8 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Para capturar cabeceras aun si el SSL falla
		},
	}

	url := fmt.Sprintf("https://%s", target)
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, "", err
	}

	// User-Agent personalizado para evitar bloqueos básicos
	req.Header.Set("User-Agent", "HeaderShield-Scanner/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("error al conectar con %s: %w", url, err)
	}
	defer resp.Body.Close()

	var results []HeaderCheck

	// Evaluar cabeceras de seguridad
	for header, info := range securityHeaders {
		val := resp.Header.Get(header)
		results = append(results, HeaderCheck{
			Name:        header,
			Present:     val != "",
			Value:       val,
			Description: info.description,
			Severity:    info.severity,
		})
	}

	serverHeader := resp.Header.Get("Server")

	return results, serverHeader, nil
}