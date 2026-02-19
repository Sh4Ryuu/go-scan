# GoScan - Advanced Port Scanner

A high-performance, multi-threaded port scanner written in Go with advanced security capabilities. Inspired by RustScan with features for TCP/UDP scanning, SSL/TLS certificate grabbing, banner grabbing, and geolocation lookup.

## Features

- Multi-threaded Scanning: Concurrent TCP and UDP port scanning with configurable workers
- TCP & UDP Support: Scan both TCP and UDP ports simultaneously
- Banner Grabbing: Automatically grab service banners from open ports
- SSL/TLS Certificate Grabbing: Extract and analyze SSL/TLS certificates with SHA-256 fingerprints
- Geolocation Lookup: Get geolocation information about target IP addresses using ip-api.com
- Rate Limiting: Control scanning speed to avoid network flooding
- Beautiful Output: Colored output with progress bar and detailed statistics
- Multiple Output Formats: JSON output for easy integration with other tools
- Scanning Profiles: Pre-configured profiles (aggressive, default, conservative)

## Installation

### Prerequisites
- Go 1.21 or later
- (Optional) Nmap for advanced script integration

### Quick Start

```bash
# Clone the repository
git clone https://github.com/Sh4Ryuu/go-scan.git
cd go-scan

# Build the project
go build -o go-scan ./cmd/main.go

# Or run directly
go run ./cmd/main.go -help
```

## Quick Start Examples

### Basic Scan
```bash
./go-scan -host example.com
```

### Aggressive Scan with All Features
```bash
./go-scan -host example.com -profile aggressive -ssl -geo -banners
```

### Full Port Scan
```bash
./go-scan -host example.com -end 65535 -profile aggressive
```

### Scan Specific Ports
```bash
./go-scan -host example.com -start 1 -end 1000
```

### Quiet Mode (Only Open Ports)
```bash
./go-scan -host example.com -quiet
```

### JSON Output
```bash
./go-scan -host example.com -json > results.json
```

## Command-Line Options

### Basic Options
```
-host string              Target host to scan (default: scanme.nmap.org)
-start int                Starting port number (default: 1)
-end int                  Ending port number (default: 1024)
-workers int              Number of concurrent workers (default: 100)
-timeout int              Connection timeout in seconds (default: 1)
```

### Scanning Profiles
```
-profile string           Scanning profile: aggressive, default, conservative (default: default)
```

### Features
```
-banners bool             Enable banner grabbing (default: true)
-ssl bool                 Enable SSL/TLS certificate grabbing (default: true)
-udp bool                 Enable UDP scanning (default: false)
-geo bool                 Enable geolocation lookup (default: true)
-nmap string              Nmap scripts to run (comma-separated)
```

### Output Options
```
-verbose bool             Enable verbose output (default: false)
-quiet bool               Quiet mode - only show open ports (default: false)
-json bool                Output results as JSON (default: false)
```

### Other Options
```
-rate-limit int           Rate limit in milliseconds (default: 10)
-help                     Show help message
-profiles                 Show available scanning profiles
-nmap-help                Show available Nmap scripts
```
## Advanced Features

### SSL/TLS Certificate Grabbing
The scanner automatically extracts:
- Certificate subject and issuer
- Valid from/until dates
- DNS names
- Public key size and algorithm
- Certificate fingerprint (SHA-256)
- Expiration status

### Geolocation Lookup
Uses ip-api.com service to return:
- Country and country code
- City and region
- ISP information
- GPS coordinates

### Banner Grabbing
Automatically reads initial responses from services to identify:
- Service name and version
- Protocol information
- Server details

## Nmap Scripts Integration

See NMAP_SCRIPTS.md for detailed information.

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues for bugs and feature requests.

## Support

For issues, questions, or suggestions, please open an issue on GitHub:
https://github.com/Sh4Ryuu/go-scan/issues
