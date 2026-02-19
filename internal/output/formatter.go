package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sh4Ryuu/go-scan/pkg/models"
)

const (
	ColorReset   = "\033[0m"
	ColorBold    = "\033[1m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorGray    = "\033[90m"

	SymCheck   = "[+]"
	SymCross   = "[-]"
	SymArrow   = "-->"
	SymBolt    = "[*]"
	SymInfo    = "[i]"
	SymWarning = "[!]"
	SymGeo     = "[G]"
	SymCert    = "[S]"
	SymNetwork = "[N]"
)

// FormatterConfig holds the minimal config needed for formatting
type FormatterConfig struct {
	Quiet             bool
	Verbose           bool
	JSONOutput        bool
	StartPort         int
	EndPort           int
	Host              string
	MaxWorkers        int
	TimeoutSeconds    int
	Profile           string
	BannerGrabbing    bool
	EnableSSL         bool
	EnableUDP         bool
	EnableGeolocation bool
	NmapScripts       string
}

type Formatter struct {
	config    *FormatterConfig
	startTime time.Time
}

// NewFormatter creates a new output formatter
func NewFormatter(config *FormatterConfig) *Formatter {
	return &Formatter{
		config:    config,
		startTime: time.Now(),
	}
}

// PrintBanner prints the application banner
func (f *Formatter) PrintBanner() {
	if f.config.Quiet || f.config.JSONOutput {
		return
	}

	fmt.Printf("%s%s", ColorBold, ColorCyan)
	fmt.Println("===================================================================")
	fmt.Println("                  GoScan - Advanced Port Scanner")
	fmt.Println("===================================================================")
	fmt.Println("   TCP & UDP Scanning         Banner Grabbing")
	fmt.Println("   SSL/TLS Certificates       Geolocation Lookup")
	fmt.Println("   Nmap Script Integration    Multi-threaded")
	fmt.Println("===================================================================")
	fmt.Print(ColorReset)
	fmt.Println()
}

// PrintConfigInfo prints configuration information
func (f *Formatter) PrintConfigInfo() {
	if f.config.Quiet || f.config.JSONOutput {
		return
	}

	portCount := f.config.EndPort - f.config.StartPort + 1
	fmt.Printf("%s%s SCAN CONFIGURATION %s\n", ColorBold, ColorCyan, ColorReset)
	fmt.Printf("  %s Target Host       : %s%s%s\n", SymInfo, ColorBold, f.config.Host, ColorReset)
	fmt.Printf("  %s Port Range        : %s%d-%d (%d ports)%s\n", SymInfo, ColorBold, f.config.StartPort, f.config.EndPort, portCount, ColorReset)
	fmt.Printf("  %s Workers           : %s%d%s\n", SymBolt, ColorBold, f.config.MaxWorkers, ColorReset)
	fmt.Printf("  %s Timeout           : %s%ds%s\n", SymInfo, ColorBold, f.config.TimeoutSeconds, ColorReset)
	fmt.Printf("  %s Profile           : %s%s%s\n", SymInfo, ColorBold, f.config.Profile, ColorReset)

	features := []string{}
	if f.config.BannerGrabbing {
		features = append(features, "Banner Grabbing")
	}
	if f.config.EnableSSL {
		features = append(features, "SSL/TLS Certs")
	}
	if f.config.EnableUDP {
		features = append(features, "UDP Scan")
	}
	if f.config.EnableGeolocation {
		features = append(features, "Geolocation")
	}
	if f.config.NmapScripts != "" {
		features = append(features, fmt.Sprintf("Nmap Scripts (%s)", f.config.NmapScripts))
	}

	if len(features) > 0 {
		fmt.Printf("  %s Features         : %s%s%s\n", SymInfo, ColorBold, strings.Join(features, ", "), ColorReset)
	}

	fmt.Println()
}

// PrintProgress prints scanning progress
func (f *Formatter) PrintProgress(current, total int) {
	if f.config.Quiet || f.config.JSONOutput {
		return
	}

	if total == 0 {
		return
	}

	percent := (current * 100) / total
	barLen := 40
	filled := (percent * barLen) / 100

	bar := strings.Repeat("", filled) + strings.Repeat("", barLen-filled)
	fmt.Printf("\rProgress: [%s] %d%% (%d/%d)", bar, percent, current, total)
}

// PrintResults prints scan results
func (f *Formatter) PrintResults(result *models.ScanResult) {
	if f.config.JSONOutput {
		f.printJSON(result)
	} else {
		f.printTextResults(result)
	}
}

// printTextResults prints results in text format
func (f *Formatter) printTextResults(result *models.ScanResult) {
	if f.config.Quiet {
		if result.Status == "open" {
			fmt.Printf("%s:%d\n", result.Host, result.Port)
		}
		return
	}

	statusColor := ColorGray
	statusSymbol := SymWarning

	switch result.Status {
	case "open":
		statusColor = ColorGreen
		statusSymbol = SymCheck
	case "filtered":
		statusColor = ColorYellow
		statusSymbol = SymWarning
	case "closed":
		statusColor = ColorRed
		statusSymbol = SymCross
	}

	fmt.Printf("%s%s%s ", statusColor, statusSymbol, ColorReset)
	fmt.Printf("%s:%d", result.Host, result.Port)

	if result.Service != "" {
		fmt.Printf(" (%s)", result.Service)
	}

	if result.Banner != "" {
		fmt.Printf(" - %s%s%s", ColorCyan, result.Banner, ColorReset)
	}

	if result.IsSSL && result.SSLInfo != nil {
		fmt.Printf(" %s[HTTPS]%s", ColorMagenta, ColorReset)
	}

	fmt.Println()

	if result.SSLInfo != nil && f.config.Verbose {
		f.printSSLInfo(result.SSLInfo)
	}

	if result.Geolocation != nil && f.config.Verbose {
		f.printGeolocation(result.Geolocation)
	}
}

// printSSLInfo prints SSL certificate information
func (f *Formatter) printSSLInfo(info *models.SSLCertInfo) {
	fmt.Printf("    %s SSL Certificate:\n", SymCert)
	fmt.Printf("      Subject: %s\n", info.Subject)
	fmt.Printf("      Issuer: %s\n", info.Issuer)
	fmt.Printf("      Valid: %s to %s\n", info.ValidFrom.Format("2006-01-02"), info.ValidTo.Format("2006-01-02"))
	if len(info.DNSNames) > 0 {
		fmt.Printf("      DNS Names: %s\n", strings.Join(info.DNSNames, ", "))
	}
	fmt.Printf("      Fingerprint (SHA-256): %s\n", info.Fingerprint)
}

// printGeolocation prints geolocation information
func (f *Formatter) printGeolocation(geo *models.GeoLocation) {
	fmt.Printf("    %s Geolocation:\n", SymGeo)
	if geo.Country != "" {
		fmt.Printf("      Country: %s (%s)\n", geo.Country, geo.CountryCode)
	}
	if geo.City != "" {
		fmt.Printf("      City: %s\n", geo.City)
	}
	if geo.ISP != "" {
		fmt.Printf("      ISP: %s\n", geo.ISP)
	}
	if geo.Latitude != 0 && geo.Longitude != 0 {
		fmt.Printf("      Coordinates: %.4f, %.4f\n", geo.Latitude, geo.Longitude)
	}
}

// PrintStatistics prints scan statistics
func (f *Formatter) PrintStatistics(stats *models.ScanStats) {
	if f.config.Quiet || f.config.JSONOutput {
		return
	}

	duration := time.Since(stats.StartTime)
	fmt.Printf("\n%s%s SCAN STATISTICS %s\n", ColorBold, ColorCyan, ColorReset)
	fmt.Printf("  %s Total Ports Scanned : %d\n", SymInfo, stats.TotalPorts)
	fmt.Printf("  %s Open Ports          : %s%d%s\n", SymCheck, ColorGreen, stats.OpenPorts, ColorReset)
	fmt.Printf("  %s Closed Ports        : %s%d%s\n", SymCross, ColorRed, stats.ClosedPorts, ColorReset)
	fmt.Printf("  %s Filtered Ports      : %s%d%s\n", SymWarning, ColorYellow, stats.FilteredPorts, ColorReset)
	fmt.Printf("  %s Scan Duration       : %.2fs\n", SymBolt, duration.Seconds())
	fmt.Println()
}

// printJSON prints results in JSON format
func (f *Formatter) printJSON(result *models.ScanResult) {
	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

// PrintError prints an error message
func (f *Formatter) PrintError(msg string) {
	fmt.Printf("%s%s Error: %s%s\n", ColorRed, ColorBold, msg, ColorReset)
}
