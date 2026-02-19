# GoScan Quick Start Guide

## ¦ Setup (One-time)

### 1. Initialize the Go Module
```bash
cd go-scan
go mod init github.com/Sh4Ryuu/go-scan
```

### 2. Download Dependencies
```bash
go mod tidy
```

### 3. Build the Binary
```bash
go build -o go-scan ./cmd/main.go
```

On Windows:
```bash
go build -o go-scan.exe ./cmd/main.go
```

##  Test It Immediately

### Test 1: Show Help
```bash
./go-scan -help
```

### Test 2: Scan Local Ports (Quick)
```bash
./go-scan -host localhost -start 1 -end 100 -timeout 1 -quiet
```

### Test 3: Scan Public Target (scanme.nmap.org)
```bash
./go-scan -host scanme.nmap.org -start 1 -end 100
```

Expected output: Should find port 22 (SSH) and port 80 (HTTP) open with banners!

### Test 4: Show Available Profiles
```bash
./go-scan -profiles
```

### Test 5: Aggressive Scan
```bash
./go-scan -host scanme.nmap.org -profile aggressive -end 500 -timeout 2
```

### Test 6: JSON Output
```bash
./go-scan -host scanme.nmap.org -start 20 -end 30 -json
```

## ¤ Push to GitHub

### 1. Create Repository on GitHub
- Go to https://github.com/new
- Create repo named `go-scan`
- **Do NOT initialize with README** (you already have one)

### 2. Push Your Code
```bash
cd C:\Users\MohamedAliZehri\go-work\go-scan

# Initialize git
git init

# Add all files
git add .

# Commit
git commit -m "Initial commit: Advanced port scanner in Go"

# Add remote
git remote add origin https://github.com/Sh4Ryuu/go-scan.git

# Push to main branch
git branch -M main
git push -u origin main
```

### 3. Verify on GitHub
Visit: https://github.com/Sh4Ryuu/go-scan

## ¯ Common Commands

### Basic Scan
```bash
./go-scan -host example.com
```

### Fast Scan
```bash
./go-scan -host example.com -profile aggressive -end 1000
```

### Slow/Safe Scan
```bash
./go-scan -host example.com -profile conservative -end 1000
```

### Quiet Mode (Only Open Ports)
```bash
./go-scan -host example.com -quiet
```

### Save to JSON
```bash
./go-scan -host example.com -json > results.json
```

### Full Port Scan
```bash
./go-scan -host example.com -end 65535 -profile aggressive
```

### With Banner Grabbing
```bash
./go-scan -host example.com -banners -end 1000
```

### With SSL Certificate Grabbing
```bash
./go-scan -host example.com -ssl -end 1000
```

### With Geolocation
```bash
./go-scan -host example.com -geo -end 1000
```

## ª What to Test After Building

1.  Help works: `./go-scan -help`
2.  Quick scan works: `./go-scan -host localhost -start 1 -end 50 -quiet`
3.  Public target scan: `./go-scan -host scanme.nmap.org -start 1 -end 100`
4.  Progress bar displays
5.  Banners are grabbed (SSH on port 22)
6.  Statistics show at end
7.  JSON output works: `./go-scan -host scanme.nmap.org -start 20 -end 30 -json`
8.  Different profiles work: `-profile aggressive`, `-profile conservative`

## ‹ Module Information

- **Module Name**: `github.com/Sh4Ryuu/go-scan`
- **Go Version**: 1.21+
- **Main Entry Point**: `cmd/main.go`
- **Packages**:
  - `internal/scanner` - Core scanning logic
  - `internal/output` - Output formatting
  - `internal/ssl` - SSL/TLS certificate grabbing
  - `internal/geolocation` - IP geolocation lookup
  - `pkg/models` - Data structures

##  Important Notes

1. **Authorization Required**: Only scan networks you own or have explicit permission to scan
2. **Rate Limiting**: Use `-rate-limit` to avoid network flooding
3. **Timeout Settings**: Increase `-timeout` for slow networks
4. **Worker Count**: Increase `-workers` for faster scanning on powerful machines
5. **Profile Selection**: Use profiles that match your network reliability

## ˜ Troubleshooting

### Build Fails
```bash
go mod tidy
go clean
go build -o go-scan ./cmd/main.go
```

### Timeout Issues
```bash
./go-scan -host example.com -timeout 5
```

### Too Many Errors
```bash
./go-scan -host example.com -profile conservative -rate-limit 100
```

### Permission Denied (Linux/Mac)
```bash
chmod +x go-scan
./go-scan -help
```

## š Next Steps

1. Create GitHub repository
2. Push code with `git push`
3. Add `.gitignore` (optional):
   ```
   go-scan
   go-scan.exe
   .DS_Store
   *.test
   ```
4. Add GitHub Actions CI/CD (optional)
5. Create releases for different platforms (optional)

## § Development

To run without building:
```bash
go run ./cmd/main.go -host scanme.nmap.org -start 1 -end 100
```

To format code:
```bash
go fmt ./...
```

To check for issues:
```bash
go vet ./...
```

---

**You're all set!** Start scanning with GoScan! €