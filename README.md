# 🛡️ HeaderShield

**HeaderShield** es una herramienta CLI rápida e independiente desarrollada en Go para auditar las cabeceras de seguridad HTTP y el estado del certificado SSL/TLS de cualquier dominio.

---

## 🚀 Características

- 🔍 **Cabeceras HTTP de Seguridad:** Evaluación de `CSP`, `HSTS`, `X-Frame-Options`, `X-Content-Type-Options`, `Referrer-Policy` y `Permissions-Policy`.
- 🔐 **Inspección SSL/TLS:** Comprobación de validez, fecha de expiración, días restantes y emisor de la CA.
- ⚡ **Sin Dependencias Externas:** Construido utilizando la librería estándar de Go.
- 📦 **Binario Nativo:** Se compila en un único archivo ejecutable listo para desplegar.

---

## 🛠️ Instalación y Uso

### Prerrequisitos
Tener instalado [Go](https://go.dev/) (v1.20 o superior).

### Ejecutar directamente
```bash
go run main.go github.com