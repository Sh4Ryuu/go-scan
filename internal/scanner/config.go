package scanner

import (
	"fmt"
	"strings"
	"time"
)

// Config holds all scanner configuration
type Config struct {
	// Basic settings
	Host           string
	StartPort      int
	EndPort        int
	MaxWorkers     int
	TimeoutSeconds int
	RateLimitMs    int

	// Feature flags
	BannerGrabbing    bool
	EnableSSL         bool
	EnableUDP         bool
	EnableGeolocation bool

	// Output settings
	Verbose    bool
	Quiet      bool
	JSONOutput bool

	// Profile and nmap
	Profile     string
	NmapScripts string

	// Internal - computed values
	Timeout       time.Duration
	WorkerTimeout time.Duration
}

// ProfileSettings define preset configurations
var ProfileSettings = map[string]map[string]interface{}{
	"aggressive": {
		"workers":   500,
		"timeout":   500,
		"rateLimit": 0,
	},
	"default": {
		"workers":   100,
		"timeout":   1000,
		"rateLimit": 10,
	},
	"conservative": {
		"workers":   50,
		"timeout":   3000,
		"rateLimit": 50,
	},
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host cannot be empty")
	}

	if c.StartPort < 1 || c.StartPort > 65535 {
		return fmt.Errorf("start port must be between 1 and 65535")
	}

	if c.EndPort < 1 || c.EndPort > 65535 {
		return fmt.Errorf("end port must be between 1 and 65535")
	}

	if c.EndPort < c.StartPort {
		return fmt.Errorf("end port must be greater than or equal to start port")
	}

	if c.MaxWorkers < 1 {
		return fmt.Errorf("workers must be at least 1")
	}

	if c.TimeoutSeconds < 1 {
		c.TimeoutSeconds = 1
	}

	// Apply profile settings if valid
	if profile, exists := ProfileSettings[c.Profile]; exists {
		if workers, ok := profile["workers"].(int); ok {
			c.MaxWorkers = workers
		}
		if timeout, ok := profile["timeout"].(int); ok {
			c.TimeoutSeconds = timeout / 1000
		}
		if rateLimit, ok := profile["rateLimit"].(int); ok {
			c.RateLimitMs = rateLimit
		}
	}

	// Set computed values
	c.Timeout = time.Duration(c.TimeoutSeconds) * time.Second
	c.WorkerTimeout = c.Timeout

	return nil
}

// GetNmapScriptsList returns parsed nmap scripts
func (c *Config) GetNmapScriptsList() []string {
	if c.NmapScripts == "" {
		return []string{}
	}

	scripts := strings.Split(c.NmapScripts, ",")
	for i := range scripts {
		scripts[i] = strings.TrimSpace(scripts[i])
	}

	return scripts
}

// IsFullScan returns true if scanning all ports
func (c *Config) IsFullScan() bool {
	return c.StartPort == 1 && c.EndPort == 65535
}

// GetPortCount returns total number of ports to scan
func (c *Config) GetPortCount() int {
	return c.EndPort - c.StartPort + 1
}
