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
