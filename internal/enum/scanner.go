// Package enum performs TCP port enumeration against an authorized target:
// a concurrent connect scan over a curated list of common ports, with a
// passive banner read on each open port.
//
// The port list is grouped below by category for readability. Richer
// per-port metadata (recommended enumeration technique, associated risk)
// is documented in docs/ports.md rather than encoded here — adding it to
// the runtime map would require a schema change (assets.category) that
// isn't justified until a concrete feature consumes it.
//
// UDP services (DNS, SNMP, VPN, VoIP) are intentionally out of scope: this
// scanner is TCP-connect only. UDP scanning needs a different technique
// (no handshake to confirm state) and is tracked as future work.
//
// Wiring nmap via os/exec remains a valid future upgrade for full-range
// scanning — not required for this module to already be useful.
package enum

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/leomarqueseh/kestrel/internal/asset"
)

// commonPorts maps port number to the service conventionally running
// there. Port number alone never confirms the service — banner grabbing
// below provides the closest thing to confirmation this module offers;
// true fingerprinting is a future enhancement.
var commonPorts = map[int]string{
	// Web
	80: "http", 443: "https", 3000: "http-dev", 5000: "http-dev",
	8000: "http", 8080: "http-proxy", 8443: "https-alt", 8888: "http-jupyter",

	// Remote access
	22: "ssh", 23: "telnet", 3389: "rdp", 5900: "vnc",
	5985: "winrm", 5986: "winrm-tls",

	// File sharing
	20: "ftp-data", 21: "ftp", 139: "netbios-ssn", 445: "smb", 873: "rsync", 2049: "nfs",

	// Directory / authentication
	88: "kerberos", 135: "msrpc", 137: "netbios-ns", 138: "netbios-dgm",
	389: "ldap", 636: "ldaps",

	// Databases
	1433: "mssql", 1521: "oracle", 3306: "mysql", 5432: "postgresql",
	6379: "redis", 9200: "elasticsearch", 9300: "elasticsearch-cluster",
	27017: "mongodb",

	// Infrastructure & platform services
	53: "dns", 111: "rpcbind", 161: "snmp", 162: "snmptrap", 514: "syslog",
	902: "vmware", 1080: "socks-proxy", 2181: "zookeeper", 2375: "docker-api",
	2376: "docker-api-tls", 5601: "kibana", 6443: "kubernetes-api",
	11211: "memcached", 50000: "sap", 5060: "sip", 5061: "sip-tls",

	// Mail
	25: "smtp", 110: "pop3", 143: "imap", 465: "smtps",
	587: "smtp-submission", 993: "imaps", 995: "pop3s",
}

const (
	dialTimeout = 1 * time.Second
	bannerWait  = 500 * time.Millisecond
	workerCount = 20 // concurrent goroutines scanning ports
)

// Run scans host across commonPorts concurrently and returns one Asset
// per open port found.
func Run(ctx context.Context, host string) ([]asset.Asset, error) {
	ports := make(chan int, len(commonPorts))
	for port := range commonPorts {
		ports <- port
	}
	close(ports)

	results := make(chan asset.Asset, len(commonPorts))
	var wg sync.WaitGroup

	// A fixed-size worker pool: workerCount goroutines pull from the same
	// ports channel until it's drained, instead of spawning one goroutine
	// per port (which would open too many sockets at once).
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range ports {
				if a, open := scanPort(host, port); open {
					results <- a
				}
			}
		}()
	}

	wg.Wait()
	close(results)

	var assets []asset.Asset
	for a := range results {
		assets = append(assets, a)
	}
	return assets, nil
}

func scanPort(host string, port int) (asset.Asset, bool) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, dialTimeout)
	if err != nil {
		return asset.Asset{}, false
	}
	defer conn.Close()

	banner := readBanner(conn)
	p := port

	return asset.Asset{
		Host:     host,
		Port:     &p,
		Protocol: "tcp",
		Service:  commonPorts[port],
		Version:  banner,
	}, true
}

// readBanner passively waits a short window for the service to speak
// first (SSH, FTP, SMTP conventionally do). It never sends probe bytes —
// this stays a connect scan, not an active fingerprinting probe.
func readBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(bannerWait))
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}
