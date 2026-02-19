package models

import (
	"time"
)

// ScanResult represents a single port scan result
type ScanResult struct {
	Host        string       `json:"host"`
	Port        int          `json:"port"`
	Protocol    string       `json:"protocol"` // "tcp" or "udp"
	Status      string       `json:"status"`   // "open", "closed", "filtered"
	Service     string       `json:"service,omitempty"`
	Banner      string       `json:"banner,omitempty"`
	IsSSL       bool         `json:"is_ssl"`
	SSLInfo     *SSLCertInfo `json:"ssl_info,omitempty"`
	Geolocation *GeoLocation `json:"geolocation,omitempty"`
	Severity    string       `json:"severity,omitempty"`
}

// SSLCertInfo contains SSL/TLS certificate information
type SSLCertInfo struct {
	Subject            string    `json:"subject"`
	Issuer             string    `json:"issuer"`
	ValidFrom          time.Time `json:"valid_from"`
	ValidTo            time.Time `json:"valid_to"`
	DNSNames           []string  `json:"dns_names,omitempty"`
	IsExpired          bool      `json:"is_expired"`
	Fingerprint        string    `json:"fingerprint"`
	PublicKeyBits      int       `json:"public_key_bits"`
	SignatureAlgorithm string    `json:"signature_algorithm"`
}

// ScanStats contains scan statistics
type ScanStats struct {
	TotalPorts        int          `json:"total_ports"`
	OpenPorts         int          `json:"open_ports"`
	ClosedPorts       int          `json:"closed_ports"`
	FilteredPorts     int          `json:"filtered_ports"`
	ErrorCount        int          `json:"error_count"`
	DurationSeconds   float64      `json:"duration_seconds"`
	PortsPerSec       float64      `json:"ports_per_second"`
	StartTime         time.Time    `json:"start_time"`
	EndTime           time.Time    `json:"end_time"`
	TargetHost        string       `json:"target_host"`
	TargetGeolocation *GeoLocation `json:"target_geolocation,omitempty"`
}

// GeoLocation contains geolocation information
type GeoLocation struct {
	IP          string  `json:"ip"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Region      string  `json:"region,omitempty"`
	City        string  `json:"city,omitempty"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ISP         string  `json:"isp,omitempty"`
	Error       string  `json:"error,omitempty"`
}

// NmapScriptResult contains results from nmap script execution
type NmapScriptResult struct {
	Script   string        `json:"script"`
	Port     int           `json:"port"`
	Protocol string        `json:"protocol"`
	Output   string        `json:"output"`
	Status   string        `json:"status"` // "success" or "error"
	Error    string        `json:"error,omitempty"`
	Duration time.Duration `json:"duration"`
}
