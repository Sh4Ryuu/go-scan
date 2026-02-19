package ssl

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Sh4Ryuu/go-scan/pkg/models"
)

// GrabCertificate retrieves SSL/TLS certificate information
func GrabCertificate(address string, timeout time.Duration) *models.SSLCertInfo {
	dialer := &tls.Dialer{
		Config: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	ctx, cancel := timeoutContext(timeout)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil
	}
	defer conn.Close()

	tlsConn, ok := conn.(*tls.Conn)
	if !ok {
		return nil
	}

	certs := tlsConn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return nil
	}

	cert := certs[0]
	now := time.Now()

	certInfo := &models.SSLCertInfo{
		Subject:            cert.Subject.String(),
		Issuer:             cert.Issuer.String(),
		ValidFrom:          cert.NotBefore,
		ValidTo:            cert.NotAfter,
		DNSNames:           cert.DNSNames,
		IsExpired:          now.After(cert.NotAfter),
		SignatureAlgorithm: cert.SignatureAlgorithm.String(),
		Fingerprint:        calculateFingerprint(cert.Raw),
		PublicKeyBits:      getKeySize(cert.PublicKey),
	}

	return certInfo
}

// calculateFingerprint calculates SHA-256 fingerprint of certificate
func calculateFingerprint(certDER []byte) string {
	hash := sha256.Sum256(certDER)
	return hex.EncodeToString(hash[:])
}

// getKeySize returns the size of the public key
func getKeySize(pubKey interface{}) int {
	switch key := pubKey.(type) {
	case *rsa.PublicKey:
		return key.N.BitLen()
	case *ecdsa.PublicKey:
		return key.Curve.Params().BitSize
	default:
		return 0
	}
}

// VerifyCertificateChain verifies the certificate chain
func VerifyCertificateChain(address string, timeout time.Duration) (bool, error) {
	dialer := &tls.Dialer{
		Config: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	ctx, cancel := timeoutContext(timeout)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	return true, nil
}

// GetCertificateInfo returns detailed certificate information
func GetCertificateInfo(address string, timeout time.Duration) (map[string]interface{}, error) {
	dialer := &tls.Dialer{
		Config: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	ctx, cancel := timeoutContext(timeout)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	tlsConn, ok := conn.(*tls.Conn)
	if !ok {
		return nil, fmt.Errorf("not a TLS connection")
	}

	state := tlsConn.ConnectionState()
	certs := state.PeerCertificates

	info := make(map[string]interface{})
	info["protocol_version"] = state.Version
	info["cipher_suite"] = state.CipherSuite
	info["certificate_count"] = len(certs)

	if len(certs) > 0 {
		cert := certs[0]
		info["subject"] = cert.Subject.String()
		info["issuer"] = cert.Issuer.String()
		info["valid_from"] = cert.NotBefore
		info["valid_until"] = cert.NotAfter
		info["dns_names"] = cert.DNSNames
	}

	return info, nil
}

// timeoutContext creates a context with timeout
func timeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
