# GoScan Nmap Scripts Guide

## Overview

GoScan integrates with Nmap's NSE (Nmap Scripting Engine) to provide advanced reconnaissance capabilities. This guide shows you which scripts are available, how to use them, and what to expect from each.

## Available Scripts

### SERVICE DETECTION & INFORMATION

#### ssh-hostkey
- **Purpose**: Retrieves and displays SSH host keys
- **Use Case**: Server fingerprinting and identification
- **Port**: 22 (SSH)
- **Example**:
  ```
  go-scan -host target.com -start 22 -end 22 -nmap ssh-hostkey
  ```
- **What You'll Get**: RSA/ECDSA/ED25519 key fingerprints and algorithms
- **Useful For**: Identifying if a server is the same across restarts, security auditing

#### banner
- **Purpose**: Generic service banner grabbing
- **Use Case**: General service identification
- **Port**: Any open port
- **Example**:
  ```
  go-scan -host target.com -end 1000 -nmap banner
  ```
- **What You'll Get**: Service name, version, and sometimes OS information
- **Useful For**: Quick service identification

#### mysql-info
- **Purpose**: Gathers MySQL server information
- **Use Case**: Database reconnaissance
- **Port**: 3306 (MySQL)
- **Example**:
  ```
  go-scan -host target.com -start 3306 -end 3306 -nmap mysql-info
  ```
- **What You'll Get**: MySQL version, authentication info, available databases
- **Useful For**: Database server assessment and vulnerability checking

#### mongodb-info
- **Purpose**: Extracts MongoDB server information
- **Use Case**: NoSQL database reconnaissance
- **Port**: 27017 (MongoDB default)
- **Example**:
  ```
  go-scan -host target.com -start 27017 -end 27017 -nmap mongodb-info
  ```
- **What You'll Get**: MongoDB version, replication status, collections info
- **Useful For**: MongoDB server assessment and configuration review

#### redis-info
- **Purpose**: Gets Redis server configuration and statistics
- **Use Case**: In-memory database reconnaissance
- **Port**: 6379 (Redis default)
- **Example**:
  ```
  go-scan -host target.com -start 6379 -end 6379 -nmap redis-info
  ```
- **What You'll Get**: Redis version, memory usage, connected clients, replication info
- **Useful For**: Redis server assessment and performance monitoring

---

### SSL/TLS ANALYSIS

#### ssl-cert
- **Purpose**: Extracts SSL/TLS certificate details
- **Use Case**: Certificate validation and HTTPS reconnaissance
- **Port**: 443, 8443, or any HTTPS port
- **Example**:
  ```
  go-scan -host target.com -start 443 -end 443 -nmap ssl-cert
  ```
- **What You'll Get**: Subject, issuer, validity dates, CN/SAN information
- **Note**: GoScan's `-ssl` flag provides similar functionality
- **Useful For**: Certificate expiration tracking, domain verification

#### ssl-enum-ciphers
- **Purpose**: Enumerates supported SSL/TLS ciphers and detects vulnerabilities
- **Use Case**: SSL/TLS security assessment
- **Port**: 443, 8443, or any HTTPS port
- **Example**:
  ```
  go-scan -host target.com -start 443 -end 443 -nmap ssl-enum-ciphers
  ```
- **What You'll Get**: Supported protocols (SSL 3.0, TLS 1.0-1.3), cipher suites, weak ciphers detected
- **Useful For**: Security auditing, PCI-DSS compliance checking, vulnerability assessment
- **Example Output Shows**:
  - Supported TLS versions
  - Strong vs weak ciphers
  - Common vulnerabilities (POODLE, BEAST, etc.)

---

### HTTP RECONNAISSANCE

#### http-title
- **Purpose**: Grabs HTTP page title and server information
- **Use Case**: Web application identification and reconnaissance
- **Port**: 80, 8080, 8000, or any HTTP port
- **Example**:
  ```
  go-scan -host target.com -start 80 -end 80 -nmap http-title
  ```
- **What You'll Get**: Page title, Server header, HTTP status code
- **Useful For**: Identifying web applications, detecting content management systems (CMS)

#### http-methods
- **Purpose**: Checks which HTTP methods are supported (GET, POST, PUT, DELETE, etc.)
- **Use Case**: Web server configuration assessment
- **Port**: 80, 8080, 8000, or any HTTP port
- **Example**:
  ```
  go-scan -host target.com -start 80 -end 80 -nmap http-methods
  ```
- **What You'll Get**: Supported HTTP methods, potential security issues
- **Useful For**: Detecting misconfigured servers, identifying unusual HTTP methods (WebDAV)

---

### SMB/WINDOWS ENUMERATION

#### smb-enum-shares
- **Purpose**: Lists available SMB network shares
- **Use Case**: Windows network reconnaissance
- **Port**: 445 (SMB), 139 (NetBIOS)
- **Example**:
  ```
  go-scan -host target.com -start 445 -end 445 -nmap smb-enum-shares
  ```
- **What You'll Get**: List of accessible shares, share types, comments
- **Useful For**: File sharing assessment, identifying data exposure risks

#### smb-os-discovery
- **Purpose**: Detects operating system and SMB server information
- **Use Case**: Windows/Samba server identification
- **Port**: 445 (SMB), 139 (NetBIOS)
- **Example**:
  ```
  go-scan -host target.com -start 445 -end 445 -nmap smb-os-discovery
  ```
- **What You'll Get**: OS name and version, SMB version, domain info, workgroup
- **Useful For**: Operating system fingerprinting, security baseline assessment

---

### FTP RECONNAISSANCE

#### ftp-anon
- **Purpose**: Tests for anonymous FTP access
- **Use Case**: FTP server security assessment
- **Port**: 21 (FTP)
- **Example**:
  ```
  go-scan -host target.com -start 21 -end 21 -nmap ftp-anon
  ```
- **What You'll Get**: Whether anonymous login is allowed, what files are accessible
- **Useful For**: Security vulnerability detection, misconfiguration identification

---

## Quick Testing Reference

### Common Service Ports

| Service | Port | Recommended Scripts |
|---------|------|-------------------|
| SSH | 22 | ssh-hostkey |
| FTP | 21 | ftp-anon |
| HTTP | 80 | http-title, http-methods |
| HTTPS | 443 | ssl-cert, ssl-enum-ciphers |
| SMB | 445 | smb-enum-shares, smb-os-discovery |
| MySQL | 3306 | mysql-info |
| MongoDB | 27017 | mongodb-info |
| Redis | 6379 | redis-info |

---

## Usage Examples

### Single Script on Single Port
```bash
go-scan -host scanme.nmap.org -start 22 -end 22 -nmap ssh-hostkey
```

### Multiple Scripts on Single Port
```bash
go-scan -host scanme.nmap.org -start 443 -end 443 -nmap ssl-cert,ssl-enum-ciphers
```

### Multiple Scripts on Multiple Ports
```bash
go-scan -host scanme.nmap.org -start 1 -end 1000 -nmap ssh-hostkey,http-title,ssl-cert
```

### With Aggressive Profile (Faster)
```bash
go-scan -host scanme.nmap.org -profile aggressive -nmap ssh-hostkey,http-title
```

### With Specific Port Range
```bash
go-scan -host scanme.nmap.org -start 20 -end 100 -nmap http-title,http-methods
```

### Combined with Banner Grabbing
```bash
go-scan -host scanme.nmap.org -banners -nmap ssh-hostkey,http-title
```

### JSON Output for Automation
```bash
go-scan -host scanme.nmap.org -nmap ssh-hostkey -json > results.json
```

---

## Testing Scenarios

### Scenario 1: Web Server Assessment
**Goal**: Identify web server and check HTTP methods

```bash
go-scan -host target.com -start 80 -end 80 -nmap http-title,http-methods
```

**What to Look For**:
- Server version from http-title
- Dangerous HTTP methods (PUT, DELETE, TRACE)
- CMS identification

---

### Scenario 2: HTTPS/TLS Security Check
**Goal**: Verify SSL/TLS configuration and cipher strength

```bash
go-scan -host target.com -start 443 -end 443 -nmap ssl-cert,ssl-enum-ciphers
```

**What to Look For**:
- Certificate validity and expiration
- Deprecated TLS versions (SSL 3.0, TLS 1.0)
- Weak ciphers
- Certificate chain validity

---

### Scenario 3: SSH Server Fingerprinting
**Goal**: Identify SSH server and get host key information

```bash
go-scan -host target.com -start 22 -end 22 -nmap ssh-hostkey
```

**What to Look For**:
- SSH version
- Host key algorithms (RSA, ECDSA, ED25519)
- Key fingerprints (useful for tracking server changes)

---

### Scenario 4: Windows Server Enumeration
**Goal**: Gather information about Windows/Samba servers

```bash
go-scan -host target.com -start 445 -end 445 -nmap smb-enum-shares,smb-os-discovery
```

**What to Look For**:
- Operating system version
- Shared folders and their accessibility
- Domain information
- SMB version (check for outdated SMB v1)

---

### Scenario 5: Database Server Assessment
**Goal**: Check for exposed database services

```bash
# MySQL
go-scan -host target.com -start 3306 -end 3306 -nmap mysql-info

# MongoDB
go-scan -host target.com -start 27017 -end 27017 -nmap mongodb-info

# Redis
go-scan -host target.com -start 6379 -end 6379 -nmap redis-info
```

**What to Look For**:
- Service version
- Authentication status
- Database names and collections
- Unnecessary services running on these ports

---

### Scenario 6: Comprehensive Network Scan
**Goal**: Quick reconnaissance of all common services

```bash
go-scan -host target.com -end 1000 \
  -nmap ssh-hostkey,http-title,ssl-cert,smb-os-discovery,ftp-anon
```

---

## Important Notes

### Prerequisites
- **Nmap must be installed** on your system for these scripts to work
- Verify with: `nmap -V`
- Scripts are run via Nmap's NSE engine

### Authorization
- **Always get written authorization** before running these scripts on any target
- Unauthorized scanning may be illegal in your jurisdiction
- Use only on systems you own or have explicit permission to test

### Performance Considerations
- Scripts add overhead to scanning (slower than basic port scanning)
- Use `-profile aggressive` for faster execution on trusted networks
- Combine with specific port ranges to reduce scan time
- Example: `go-scan -host target.com -start 20 -end 30 -nmap http-title`

### Results Interpretation
- Not all scripts will return results on all ports
- Some scripts require the service to be responsive
- Lack of output doesn't mean the service isn't running
- Some services may have firewalls or rate limiting enabled

### Common Issues

**"nmap not installed"**
- Install Nmap from: https://nmap.org/download.html
- Verify installation: `nmap -V`

**Script returns no output**
- Service may be behind a firewall
- Service may not be running
- Try with longer timeout: `go-scan -timeout 5 -nmap script-name`

**Scripts running slowly**
- Use `-profile aggressive` for faster execution
- Reduce port range with `-start` and `-end`
- Increase `-rate-limit` or decrease it for slower execution

---

## Real-World Testing Tips

### Best Practices

1. **Start with Known Services**
   - Test against your own servers first
   - Understand script output before using in production

2. **Use Appropriate Profiles**
   - Aggressive: Known good networks, time-critical scans
   - Default: Most scenarios, balanced speed/reliability
   - Conservative: Production systems, unreliable networks

3. **Combine Scripts Strategically**
   - Don't run all scripts on all ports
   - Target specific services with relevant scripts
   - Reduces scan time and noise

4. **Document Findings**
   - Save results to JSON: `-json > findings.json`
   - Track baseline configurations
   - Compare across scan runs

5. **Respect Rate Limits**
   - Use reasonable `-rate-limit` values (default: 10ms)
   - Avoid overwhelming the target
   - Essential for production systems

---

## Examples by Use Case

### Security Audit
```bash
go-scan -host target.com -profile conservative -end 10000 \
  -nmap ssl-enum-ciphers,ssh-hostkey,smb-os-discovery,http-title,mysql-info
```

### Vulnerability Assessment
```bash
go-scan -host target.com -profile default -end 5000 \
  -nmap ssl-enum-ciphers,http-methods,ftp-anon,smb-enum-shares
```

### Quick Recon
```bash
go-scan -host target.com -profile aggressive -end 1000 \
  -nmap ssh-hostkey,http-title,ssl-cert
```

### Asset Inventory
```bash
go-scan -host target.com -end 65535 -quiet \
  -nmap ssh-hostkey,http-title,smb-os-discovery -json > inventory.json
```

---

## View Available Scripts Anytime

```bash
go-scan -nmap-help
```

This will display all available scripts with descriptions and usage examples.

---

## Support & Documentation

For more information:
- GoScan help: `go-scan -help`
- Available profiles: `go-scan -profiles`
- Nmap documentation: https://nmap.org/nsedoc/
- GoScan repository: https://github.com/Sh4Ryuu/go-scan

---

**Remember**: Always scan responsibly with proper authorization!