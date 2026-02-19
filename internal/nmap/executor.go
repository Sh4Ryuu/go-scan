package nmap

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/Sh4Ryuu/go-scan/pkg/models"
)

// AvailableScripts contains commonly used nmap scripts
var AvailableScripts = map[string]string{
	"ssh-hostkey":      "Grabs SSH host keys",
	"ssl-cert":         "Retrieves SSL certificate info",
	"ssl-enum-ciphers": "Enumerates SSL ciphers",
	"http-title":       "Grabs HTTP page title",
	"http-methods":     "Finds supported HTTP methods",
	"smb-enum-shares":  "Enumerates SMB shares",
	"smb-os-discovery": "Detects SMB OS",
	"mysql-info":       "Gets MySQL server info",
	"mongodb-info":     "Gets MongoDB info",
	"redis-info":       "Gets Redis server info",
	"ftp-anon":         "Checks for anonymous FTP",
	"banner":           "Grabs service banner",
}

// RunScript executes an nmap script on a specific port
func RunScript(host string, port int, protocol string, script string, timeout time.Duration) models.NmapScriptResult {
	result := models.NmapScriptResult{
		Script:   script,
		Port:     port,
		Protocol: protocol,
		Status:   "error",
	}

	startTime := time.Now()
	defer func() {
		result.Duration = time.Since(startTime)
	}()

	// Check if nmap is available
	if !isNmapInstalled() {
		result.Error = "nmap not installed"
		return result
	}

	// Build nmap command
	portStr := fmt.Sprintf("%d/%s", port, protocol)
	cmdStr := fmt.Sprintf("nmap -p %s -sV --script=%s %s", portStr, script, host)

	// Execute command with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout+5*time.Second)
	defer cancel()

	output, err := executeCommand(ctx, cmdStr)
	if err != nil {
		result.Error = err.Error()
		return result
	}

	result.Output = output
	result.Status = "success"
	return result
}

// RunScriptMultiple runs multiple scripts on a port
func RunScriptMultiple(host string, port int, protocol string, scripts []string, timeout time.Duration) []models.NmapScriptResult {
	results := make([]models.NmapScriptResult, 0, len(scripts))

	for _, script := range scripts {
		result := RunScript(host, port, protocol, script, timeout)
		results = append(results, result)
	}

	return results
}

// isNmapInstalled checks if nmap is available
func isNmapInstalled() bool {
	cmd := exec.Command("nmap", "-V")
	err := cmd.Run()
	return err == nil
}

// executeCommand executes a shell command and returns output
func executeCommand(ctx context.Context, cmdStr string) (string, error) {
	cmd := exec.CommandContext(ctx, "nmap", parseNmapArgs(cmdStr)...)

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command failed: %v", err)
	}

	return out.String(), nil
}

// parseNmapArgs parses nmap command string to arguments
func parseNmapArgs(cmdStr string) []string {
	parts := strings.Fields(cmdStr)
	if len(parts) > 0 && parts[0] == "nmap" {
		return parts[1:]
	}
	return parts
}

// ListAvailableScripts returns list of available scripts
func ListAvailableScripts() map[string]string {
	return AvailableScripts
}

// ValidateScript checks if a script name is valid
func ValidateScript(script string) bool {
	_, exists := AvailableScripts[script]
	return exists
}
