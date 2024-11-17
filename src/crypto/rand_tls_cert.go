package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

// generateTLSCertificate erstellt ein temporäres Zertifikat und speichert es zusammen mit dem privaten Schlüssel
func GenerateTempTLSLocalhostCertificate() (*tls.Certificate, error) {
	// Erzeuge einen privaten ECC-Schlüssel mit secp384r1 (P-384)
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Generieren des privaten Schlüssels: %w", err)
	}

	// Erstelle ein selbstsigniertes Zertifikat
	template := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(time.Now().UnixNano()),
		Subject:      pkix.Name{Organization: []string{"Temp Org"}},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// Selbstsigniertes Zertifikat erzeugen
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Erzeugen des Zertifikats: %w", err)
	}

	// Zertifikat und privaten Schlüssel in PEM-Format kodieren
	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	privkeyBytes, _ := x509.MarshalECPrivateKey(privateKey)
	keyPem := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privkeyBytes})

	// Konvertiere die PEM-kodierten Daten in tls.Certificate
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Erzeugen des tls.Certificate-Objekts: %w", err)
	}

	return &cert, nil
}
