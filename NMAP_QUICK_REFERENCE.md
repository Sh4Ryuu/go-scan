# GoScan Nmap Scripts - Quick Reference

## View Available Scripts
```bash
go-scan -nmap-help
```

## Basic Syntax
```bash
go-scan -host <target> -start <port> -end <port> -nmap <script1>,<script2>,...
```

## Scripts by Category

###  SSH (Port 22)
```bash
go-scan -host target.com -start 22 -end 22 -nmap ssh-hostkey
```
**Output**: RSA/ECDSA/ED25519 key fingerprints and algorithms

---

###  HTTP (Port 80)
```bash
go-scan -host target.com -start 80 -end 80 -nmap http-title,http-methods
```
**Scripts**:
- `http-title` - Page title and server info
- `http-methods` - Supported HTTP methods (GET, POST, PUT, DELETE, etc.)

---

### ’ HTTPS/SSL (Port 443)
```bash
go-scan -host target.com -start 443 -end 443 -nmap ssl-cert,ssl-enum-ciphers
```
**Scripts**:
- `ssl-cert` - Certificate details (subject, issuer, validity)
- `ssl-enum-ciphers` - Supported TLS versions and cipher strength

---

### Ÿ SMB/Windows (Port 445)
```bash
go-scan -host target.com -start 445 -end 445 -nmap smb-enum-shares,smb-os-discovery
```
**Scripts**:
- `smb-enum-shares` - List accessible shares
- `smb-os-discovery` - OS version and SMB info

---

### „ Databases

**MySQL (Port 3306)**
```bash
go-scan -host target.com -start 3306 -end 3306 -nmap mysql-info
```

**MongoDB (Port 27017)**
```bash
go-scan -host target.com -start 27017 -end 27017 -nmap mongodb-info
```

**Redis (Port 6379)**
```bash
go-scan -host target.com -start 6379 -end 6379 -nmap redis-info
```

---

###  FTP (Port 21)
```bash
go-scan -host target.com -start 21 -end 21 -nmap ftp-anon
```
**Output**: Whether anonymous login is allowed

---

### ¯ Generic Service Banner
```bash
go-scan -host target.com -end 1000 -nmap banner
```
**Output**: Service name and version on any open port

---

## All 12 Available Scripts

| Script | Port(s) | Purpose |
|--------|---------|---------|
| ssh-hostkey | 22 | SSH host key fingerprinting |
| http-title | 80, 8080 | HTTP page title and server |
| http-methods | 80, 8080 | HTTP methods enumeration |
| ssl-cert | 443, 8443 | SSL certificate extraction |
| ssl-enum-ciphers | 443, 8443 | SSL/TLS cipher enumeration |
| smb-enum-shares | 445, 139 | SMB shares listing |
| smb-os-discovery | 445, 139 | OS detection from SMB |
| ftp-anon | 21 | Anonymous FTP testing |
| mysql-info | 3306 | MySQL server info |
| mongodb-info | 27017 | MongoDB server info |
| redis-info | 6379 | Redis server info |
| banner | Any | Generic service banner |

---

## Common Combinations

### Web Server Full Assessment
```bash
go-scan -host target.com -start 80 -end 80 -nmap http-title,http-methods
go-scan -host target.com -start 443 -end 443 -nmap ssl-cert,ssl-enum-ciphers
```

### Quick Full Scan with Scripts
```bash
go-scan -host target.com -end 1000 \
  -nmap ssh-hostkey,http-title,ssl-cert,smb-os-discovery
```

### Database Enumeration
```bash
go-scan -host target.com -start 3306 -end 3306 -nmap mysql-info
go-scan -host target.com -start 27017 -end 27017 -nmap mongodb-info
go-scan -host target.com -start 6379 -end 6379 -nmap redis-info
```

### Aggressive Scan (Faster)
```bash
go-scan -host target.com -profile aggressive -end 1000 \
  -nmap ssh-hostkey,http-title,ssl-cert
```

### Conservative Scan (Slower, Safer)
```bash
go-scan -host target.com -profile conservative -end 1000 \
  -nmap ssh-hostkey,http-title,ssl-cert
```

---

## Quick Tips

 **Comma-separated** for multiple scripts
```bash
-nmap ssh-hostkey,http-title,ssl-cert
```

 **Specific port range** for focused testing
```bash
-start 20 -end 30
```

 **JSON output** for automation
```bash
-json > results.json
```

 **Quiet mode** for open ports only
```bash
-quiet
```

 **Verbose** for detailed output
```bash
-verbose
```

---

## Important Notes

 **Requires Nmap**: Install from https://nmap.org/download.html

 **Authorization Required**: Only scan systems you own or have explicit permission to test

 **Script Prerequisites**:
- Service must be running on the port
- Script must be compatible with the service
- Some services may be firewalled

 **No Scripts Launched Yet**: These are safe to view/plan - nothing runs until you execute!

---

## Testing on scanme.nmap.org

Safe target for learning:
```bash
# SSH
go-scan -host scanme.nmap.org -start 22 -end 22 -nmap ssh-hostkey

# HTTP
go-scan -host scanme.nmap.org -start 80 -end 80 -nmap http-title

# Multiple
go-scan -host scanme.nmap.org -start 1 -end 100 -nmap ssh-hostkey,http-title
```

---

## Troubleshooting

**"nmap not installed"**
 Download and install Nmap: https://nmap.org/download.html

**No script output**
 Service may not be running or is firewalled. Verify port is open first.

**Scripts running slowly**
 Use `-profile aggressive` or reduce port range with `-start` and `-end`

---

## More Information

- Full guide: `NMAP_SCRIPTS.md`
- GoScan help: `go-scan -help`
- List profiles: `go-scan -profiles`
- View scripts: `go-scan -nmap-help`

---

**Remember: Scan responsibly with proper authorization!** ’
