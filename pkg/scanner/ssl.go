package scanner

import (
	"crypto/tls"
	"net"
	"time"
)

// CheckSSL realiza el handshake TLS y extrae información del certificado.
func CheckSSL(target string) SSLResult {
	dialer := &net.Dialer{Timeout: 5 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", net.JoinHostPort(target, "443"), nil)
	if err != nil {
		return SSLResult{Valid: false, Error: err.Error()}
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return SSLResult{Valid: false, Error: "no se encontraron certificados"}
	}

	leafCert := certs[0]
	now := time.Now()
	daysRemaining := int(leafCert.NotAfter.Sub(now).Hours() / 24)

	return SSLResult{
		Valid:          now.Before(leafCert.NotAfter) && now.After(leafCert.NotBefore),
		Issuer:         leafCert.Issuer.CommonName,
		ExpirationDate: leafCert.NotAfter,
		DaysRemaining:  daysRemaining,
	}
}