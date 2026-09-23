// Package enum performs TCP port enumeration against an authorized target:
// a concurrent connect scan over a curated list of common ports, with a
// passive banner read on each open port.
//
// Wiring nmap via os/exec is a valid future upgrade for full-range or
// UDP scanning — not required for this module to already be useful.
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

// commonPorts maps port number to the service conventionally running there.
var commonPorts = map[int]string{
	21: "ftp", 22: "ssh", 23: "telnet", 25: "smtp", 53: "dns",
	80: "http", 110: "pop3", 111: "rpcbind", 135: "msrpc", 139: "netbios-ssn",
	143: "imap", 443: "https", 445: "microsoft-ds", 465: "smtps", 587: "submission",
	993: "imaps", 995: "pop3s", 1433: "mssql", 1723: "pptp", 3306: "mysql",
	3389: "rdp", 5432: "postgresql", 5900: "vnc", 6379: "redis",
	8080: "http-alt", 8443: "https-alt", 9200: "elasticsearch",
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
