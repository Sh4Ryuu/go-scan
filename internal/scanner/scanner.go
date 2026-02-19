package scanner

import (
	"bufio"
	"fmt"
	"net"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Sh4Ryuu/go-scan/internal/geolocation"
	"github.com/Sh4Ryuu/go-scan/internal/output"
	"github.com/Sh4Ryuu/go-scan/internal/ssl"
	"github.com/Sh4Ryuu/go-scan/pkg/models"
)

// PortScanner is the main scanner struct
type PortScanner struct {
	config    *Config
	formatter *output.Formatter
	stats     *models.ScanStats
}

// NewPortScanner creates a new port scanner
func NewPortScanner(config *Config, formatter *output.Formatter) *PortScanner {
	return &PortScanner{
		config:    config,
		formatter: formatter,
		stats: &models.ScanStats{
			TargetHost: config.Host,
			StartTime:  time.Now(),
		},
	}
}

// Scan performs the port scan
func (ps *PortScanner) Scan() ([]models.ScanResult, *models.ScanStats, error) {
	startTime := time.Now()
	ps.stats.StartTime = startTime

	totalPorts := ps.config.GetPortCount()

	// Resolve target IP if needed for geolocation
	var targetIP string
	if ps.config.EnableGeolocation {
		ips, err := net.LookupIP(ps.config.Host)
		if err == nil && len(ips) > 0 {
			targetIP = ips[0].String()
		}
	}

	// TCP Scanning
	results := ps.scanTCP(totalPorts)

	// UDP Scanning if enabled
	if ps.config.EnableUDP {
		udpResults := ps.scanUDP(totalPorts)
		results = append(results, udpResults...)
	}

	// Sort results
	sort.Slice(results, func(i, j int) bool {
		if results[i].Port == results[j].Port {
			return results[i].Protocol < results[j].Protocol
		}
		return results[i].Port < results[j].Port
	})

	// Geolocation lookup if enabled
	if ps.config.EnableGeolocation && targetIP != "" {
		ps.stats.TargetGeolocation = geolocation.LookupIP(targetIP)
	}

	// Update stats
	ps.stats.EndTime = time.Now()
	ps.stats.TotalPorts = totalPorts
	ps.stats.OpenPorts = countOpen(results, "tcp")
	ps.stats.ClosedPorts = totalPorts - ps.stats.OpenPorts
	ps.stats.DurationSeconds = ps.stats.EndTime.Sub(startTime).Seconds()
	ps.stats.PortsPerSec = float64(totalPorts) / ps.stats.DurationSeconds

	return results, ps.stats, nil
}

// scanTCP performs TCP port scanning
func (ps *PortScanner) scanTCP(totalPorts int) []models.ScanResult {
	ports := make(chan int, ps.config.MaxWorkers)
	results := make(chan models.ScanResult, totalPorts)
	var wg sync.WaitGroup
	var scanned int64

	// Start workers
	for i := 0; i < ps.config.MaxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range ports {
				result := ps.probeTCP(port)
				if result.Status != "" {
					results <- result
				}
				atomic.AddInt64(&scanned, 1)

				// Show progress
				current := int(atomic.LoadInt64(&scanned))
				if current%10 == 0 {
					ps.formatter.PrintProgress(current, totalPorts)
				}

				// Rate limiting
				if ps.config.RateLimitMs > 0 {
					time.Sleep(time.Duration(ps.config.RateLimitMs) * time.Millisecond)
				}
			}
		}()
	}

	// Send ports to scan
	go func() {
		for port := ps.config.StartPort; port <= ps.config.EndPort; port++ {
			ports <- port
		}
		close(ports)
	}()

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	var scanResults []models.ScanResult
	for result := range results {
		scanResults = append(scanResults, result)
	}

	return scanResults
}

// scanUDP performs UDP port scanning
func (ps *PortScanner) scanUDP(totalPorts int) []models.ScanResult {
	var results []models.ScanResult

	for port := ps.config.StartPort; port <= ps.config.EndPort; port++ {
		result := ps.probeUDP(port)
		if result.Status != "" {
			results = append(results, result)
		}

		if ps.config.RateLimitMs > 0 {
			time.Sleep(time.Duration(ps.config.RateLimitMs) * time.Millisecond)
		}
	}

	return results
}

// probeTCP probes a single TCP port
func (ps *PortScanner) probeTCP(port int) models.ScanResult {
	result := models.ScanResult{
		Host:     ps.config.Host,
		Port:     port,
		Protocol: "tcp",
		Status:   "closed",
	}

	address := fmt.Sprintf("%s:%d", ps.config.Host, port)
	conn, err := net.DialTimeout("tcp", address, ps.config.Timeout)
	if err != nil {
		result.Status = "closed"
		return result
	}
	defer conn.Close()

	result.Status = "open"

	// Banner grabbing
	if ps.config.BannerGrabbing {
		banner := grabBanner(conn)
		if banner != "" {
			result.Banner = banner
		}
	}

	// SSL/TLS certificate grabbing
	if ps.config.EnableSSL && (port == 443 || port == 8443) {
		certInfo := ssl.GrabCertificate(address, ps.config.Timeout)
		if certInfo != nil {
			result.IsSSL = true
			result.SSLInfo = certInfo
		}
	}

	return result
}

// probeUDP probes a single UDP port
func (ps *PortScanner) probeUDP(port int) models.ScanResult {
	result := models.ScanResult{
		Host:     ps.config.Host,
		Port:     port,
		Protocol: "udp",
		Status:   "closed",
	}

	address := fmt.Sprintf("%s:%d", ps.config.Host, port)
	conn, err := net.DialTimeout("udp", address, ps.config.Timeout)
	if err != nil {
		result.Status = "closed"
		return result
	}
	defer conn.Close()

	// Try to write to UDP port
	_, err = conn.Write([]byte("test"))
	if err != nil {
		result.Status = "closed"
		return result
	}

	result.Status = "open"
	return result
}

// grabBanner attempts to grab banner from a connection
func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	defer conn.SetReadDeadline(time.Time{})

	reader := bufio.NewReader(conn)
	banner, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return banner[:len(banner)-1]
}

// countOpen counts open ports of a given protocol
func countOpen(results []models.ScanResult, protocol string) int {
	count := 0
	for _, r := range results {
		if r.Protocol == protocol && r.Status == "open" {
			count++
		}
	}
	return count
}
