package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sh4Ryuu/go-scan/internal/nmap"
	"github.com/Sh4Ryuu/go-scan/internal/output"
	"github.com/Sh4Ryuu/go-scan/internal/scanner"
)

func main() {
	config := &scanner.Config{}

	// Define command-line flags
	flag.StringVar(&config.Host, "host", "scanme.nmap.org", "Target host to scan")
	flag.IntVar(&config.StartPort, "start", 1, "Starting port number")
	flag.IntVar(&config.EndPort, "end", 1024, "Ending port number")
	flag.IntVar(&config.MaxWorkers, "workers", 100, "Number of concurrent workers")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.Quiet, "quiet", false, "Quiet mode - only show open ports")
	flag.BoolVar(&config.JSONOutput, "json", false, "Output results as JSON")
	flag.BoolVar(&config.BannerGrabbing, "banners", true, "Enable banner grabbing")
	flag.BoolVar(&config.EnableSSL, "ssl", true, "Enable SSL/TLS certificate grabbing")
	flag.BoolVar(&config.EnableUDP, "udp", false, "Enable UDP scanning")
	flag.BoolVar(&config.EnableGeolocation, "geo", true, "Enable geolocation lookup")
	flag.StringVar(&config.Profile, "profile", "default", "Scanning profile (aggressive, default, conservative)")
	flag.StringVar(&config.NmapScripts, "nmap", "", "Nmap scripts to run (comma-separated, e.g., 'ssh-hostkey,ssl-cert')")
	flag.IntVar(&config.RateLimitMs, "rate-limit", 10, "Rate limit in milliseconds")

	timeout := flag.Int("timeout", 1, "Connection timeout in seconds")
	help := flag.Bool("help", false, "Show help message")
	showProfiles := flag.Bool("profiles", false, "Show available profiles")
	showNmapScripts := flag.Bool("nmap-help", false, "Show available Nmap scripts")

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *showProfiles {
		printProfiles()
		return
	}

	if *showNmapScripts {
		printNmapScripts()
		return
	}

	// Apply timeout
	config.TimeoutSeconds = *timeout

	// Validate configuration
	if err := config.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	// Create formatter config from scanner config
	formatterConfig := &output.FormatterConfig{
		Quiet:             config.Quiet,
		Verbose:           config.Verbose,
		JSONOutput:        config.JSONOutput,
		StartPort:         config.StartPort,
		EndPort:           config.EndPort,
		Host:              config.Host,
		MaxWorkers:        config.MaxWorkers,
		TimeoutSeconds:    config.TimeoutSeconds,
		Profile:           config.Profile,
		BannerGrabbing:    config.BannerGrabbing,
		EnableSSL:         config.EnableSSL,
		EnableUDP:         config.EnableUDP,
		EnableGeolocation: config.EnableGeolocation,
		NmapScripts:       config.NmapScripts,
	}

	// Initialize output formatter
	formatter := output.NewFormatter(formatterConfig)

	// Print banner
	formatter.PrintBanner()

	// Print configuration info
	formatter.PrintConfigInfo()

	// Create and run scanner
	portScanner := scanner.NewPortScanner(config, formatter)

	// Run the scan
	results, stats, err := portScanner.Scan()
	if err != nil {
		formatter.PrintError(fmt.Sprintf("Scan error: %v", err))
		os.Exit(1)
	}

	// Print results
	if results != nil {
		for _, result := range results {
			formatter.PrintResults(&result)
		}
	}

	// Print statistics
	formatter.PrintStatistics(stats)
}

func printHelp() {
	fmt.Println(`
GoScan - Advanced Port Scanner

USAGE:
  go-scan [OPTIONS]

OPTIONS:
  -host string              Target host to scan (default: scanme.nmap.org)
  -start int                Starting port number (default: 1)
  -end int                  Ending port number (default: 1024)
  -workers int              Number of concurrent workers (default: 100)
  -timeout int              Connection timeout in seconds (default: 1)
  -profile string           Scanning profile: aggressive, default, conservative
  -banners bool             Enable banner grabbing (default: true)
  -ssl bool                 Enable SSL/TLS certificate grabbing (default: true)
  -udp bool                 Enable UDP scanning (default: false)
  -geo bool                 Enable geolocation lookup (default: true)
  -nmap string              Nmap scripts to run (comma-separated)
  -rate-limit int           Rate limit in milliseconds (default: 10)
  -verbose bool             Enable verbose output (default: false)
  -quiet bool               Quiet mode - only show open ports (default: false)
  -json bool                Output results as JSON (default: false)
  -help                     Show this help message
  -profiles                 Show available scanning profiles
  -nmap-help                Show available Nmap scripts

EXAMPLES:
  go-scan -host example.com
  go-scan -host example.com -nmap-help
  go-scan -host scanme.nmap.org -start 1 -end 100 -nmap ssh-hostkey,ssl-cert
`)
}

func printProfiles() {
	fmt.Println(`
AVAILABLE SCANNING PROFILES:

1. AGGRESSIVE
   Workers: 500, Timeout: 500ms, Rate Limit: 0ms
   Use Case: Fast scanning of trusted networks

2. DEFAULT
   Workers: 100, Timeout: 1s, Rate Limit: 10ms
   Use Case: Balanced speed and reliability

3. CONSERVATIVE
   Workers: 50, Timeout: 3s, Rate Limit: 50ms
   Use Case: Slower but safer scanning

USE: go-scan -host example.com -profile aggressive
`)
}

func printNmapScripts() {
	fmt.Println(`

                      AVAILABLE NMAP SCRIPTS                              


The following Nmap NSE scripts are available for advanced reconnaissance:
`)

	scripts := nmap.ListAvailableScripts()

	fmt.Println()
	fmt.Println("SERVICE DETECTION & INFORMATION:")
	fmt.Println("")
	serviceScripts := []string{"ssh-hostkey", "mysql-info", "mongodb-info", "redis-info", "banner"}
	for _, script := range serviceScripts {
		if desc, ok := scripts[script]; ok {
			fmt.Printf("  %-20s | %s\n", script, desc)
		}
	}

	fmt.Println()
	fmt.Println("SSL/TLS ANALYSIS:")
	fmt.Println("")
	sslScripts := []string{"ssl-cert", "ssl-enum-ciphers"}
	for _, script := range sslScripts {
		if desc, ok := scripts[script]; ok {
			fmt.Printf("  %-20s | %s\n", script, desc)
		}
	}

	fmt.Println()
	fmt.Println("HTTP RECONNAISSANCE:")
	fmt.Println("")
	httpScripts := []string{"http-title", "http-methods"}
	for _, script := range httpScripts {
		if desc, ok := scripts[script]; ok {
			fmt.Printf("  %-20s | %s\n", script, desc)
		}
	}

	fmt.Println()
	fmt.Println("SMB/WINDOWS ENUMERATION:")
	fmt.Println("")
	smbScripts := []string{"smb-enum-shares", "smb-os-discovery"}
	for _, script := range smbScripts {
		if desc, ok := scripts[script]; ok {
			fmt.Printf("  %-20s | %s\n", script, desc)
		}
	}

	fmt.Println()
	fmt.Println("FTP RECONNAISSANCE:")
	fmt.Println("")
	ftpScripts := []string{"ftp-anon"}
	for _, script := range ftpScripts {
		if desc, ok := scripts[script]; ok {
			fmt.Printf("  %-20s | %s\n", script, desc)
		}
	}

	fmt.Println(`



USAGE EXAMPLES:

Single script:
  go-scan -host example.com -end 1000 -nmap ssh-hostkey

Multiple scripts (comma-separated):
  go-scan -host example.com -end 1000 -nmap ssh-hostkey,ssl-cert,http-title

Combine with profiles:
  go-scan -host example.com -profile aggressive -nmap ssl-enum-ciphers

With specific port range:
  go-scan -host scanme.nmap.org -start 20 -end 100 -nmap ssl-cert,http-title



SCRIPT DESCRIPTIONS:

  ssh-hostkey        - SSH host key fingerprinting (good for server identification)
  ssl-cert           - Extract SSL/TLS certificate info (use -ssl flag instead)
  ssl-enum-ciphers   - List supported SSL/TLS ciphers and check vulnerabilities
  http-title         - Grab HTTP page title and server version
  http-methods       - Check supported HTTP methods (GET, POST, DELETE, etc.)
  smb-enum-shares    - List available SMB network shares on Windows
  smb-os-discovery   - Detect operating system from SMB service
  mysql-info         - Get MySQL server version and configuration info
  mongodb-info       - Extract MongoDB server information
  redis-info         - Get Redis server configuration and stats
  ftp-anon           - Test for anonymous FTP access
  banner             - Generic service banner grabbing



RECOMMENDED SCRIPTS TO TEST:

For SSH Servers:
  go-scan -host target.com -start 22 -end 22 -nmap ssh-hostkey

For HTTPS/SSL Services:
  go-scan -host target.com -start 443 -end 443 -nmap ssl-cert,ssl-enum-ciphers

For HTTP Services:
  go-scan -host target.com -start 80 -end 80 -nmap http-title,http-methods

For Windows Servers (SMB):
  go-scan -host target.com -start 445 -end 445 -nmap smb-enum-shares,smb-os-discovery

For Databases:
  go-scan -host target.com -start 3306 -end 3306 -nmap mysql-info
  go-scan -host target.com -start 27017 -end 27017 -nmap mongodb-info
  go-scan -host target.com -start 6379 -end 6379 -nmap redis-info



TIPS:
 Use on known-good targets to learn about services
 Requires nmap to be installed on your system
 Combine with aggressive profile for faster scanning
 Results depend on target configuration and firewall rules
 Always obtain authorization before testing!

Note: Use 'go-scan -nmap-help' anytime to see this information.
`)
}
