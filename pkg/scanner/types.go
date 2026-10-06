package scanner

import "time"

// HeaderCheck representa el estado de una cabecera individual.
type HeaderCheck struct {
	Name        string `json:"name"`
	Present     bool   `json:"present"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description"`
	Severity    string `json:"severity"` // "HIGH", "MEDIUM", "LOW", "INFO"
}

// SSLResult almacena la información del certificado digital.
type SSLResult struct {
	Valid          bool      `json:"valid"`
	Issuer         string    `json:"issuer,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	DaysRemaining  int       `json:"days_remaining,omitempty"`
	Error          string    `json:"error,omitempty"`
}

// ScanResult guarda el análisis completo de un objetivo.
type ScanResult struct {
	Target       string        `json:"target"`
	ScanTime     time.Time     `json:"scan_time"`
	Headers      []HeaderCheck `json:"headers"`
	ServerHeader string        `json:"server_header,omitempty"`
	SSL          SSLResult     `json:"ssl"`
	Score        int           `json:"score"` // Puntuación de 0 a 100
}