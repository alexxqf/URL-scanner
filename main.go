package main

import (
	"fmt"
	"log"
	"os"

	"header-shield/pkg/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <dominio>")
		fmt.Println("Ejemplo: go run main.go github.com")
		os.Exit(1)
	}

	target := os.Args[1]
	fmt.Printf("🔍 Escaneando objetivo: %s...\n\n", target)

	// 1. Analizar Headers
	headers, server, err := scanner.CheckHeaders(target)
	if err != nil {
		log.Fatalf("❌ Error en el escaneo de cabeceras: %v", err)
	}

	// 2. Analizar SSL
	ssl := scanner.CheckSSL(target)

	// 3. Imprimir reporte por consola
	fmt.Println("======== CABECERAS DE SEGURIDAD ========")
	for _, h := range headers {
		if h.Present {
			fmt.Printf("  [✅ PRESENTE]  %-28s: %s\n", h.Name, h.Value)
		} else {
			fmt.Printf("  [❌ FALTANTE]  %-28s (Severidad: %s)\n", h.Name, h.Severity)
		}
	}

	if server != "" {
		fmt.Printf("\n⚠️  Divulgación de Servidor: Server = %s\n", server)
	}

	fmt.Println("\n======== ESTADO DEL CERTIFICADO SSL/TLS ========")
	if ssl.Valid {
		fmt.Printf("  ✅ Certificado Válido\n")
		fmt.Printf("  Emisor: %s\n", ssl.Issuer)
		fmt.Printf("  Expiración: %s (%d días restantes)\n", ssl.ExpirationDate.Format("2006-01-02"), ssl.DaysRemaining)
	} else {
		fmt.Printf("  ❌ Error SSL: %s\n", ssl.Error)
	}
}
