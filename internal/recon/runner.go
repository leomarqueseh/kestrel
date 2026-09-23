// Package recon performs passive-to-light-active reconnaissance against an
// authorized target: DNS resolution, subdomain discovery via certificate
// transparency logs, and basic HTTP service detection.
//
// This is a self-contained Go implementation. Wiring in dedicated tools
// (subfinder, httpx) via exec.Command, or via the Phase 12 Python
// automation module, is a natural future upgrade — not required for
// this module to be useful today.
package recon

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/leomarqueseh/kestrel/internal/asset"
)

var httpClient = &http.Client{Timeout: 5 * time.Second}

// Run performs reconnaissance against value (a domain, IP, or URL) and
// returns the discovered assets, ready to be persisted.
func Run(ctx context.Context, value string) ([]asset.Asset, error) {
	hosts := []string{value}

	// crt.sh failures are not fatal — DNS/HTTP checks still run on the
	// base host even if subdomain discovery is unavailable.
	if subdomains, err := discoverSubdomains(ctx, value); err == nil {
		hosts = append(hosts, subdomains...)
	}

	var results []asset.Asset
	for _, host := range dedupe(hosts) {
		results = append(results, resolveDNS(host)...)
		results = append(results, probeHTTP(host)...)
	}
	return results, nil
}

func resolveDNS(host string) []asset.Asset {
	ips, err := net.LookupHost(host)
	if err != nil {
		return nil
	}

	var results []asset.Asset
	for _, ip := range ips {
		results = append(results, asset.Asset{
			Host:     host,
			Protocol: "dns",
			Service:  "a-record",
			Version:  ip,
		})
	}
	return results
}

func probeHTTP(host string) []asset.Asset {
	var results []asset.Asset

	for _, scheme := range []string{"https", "http"} {
		url := fmt.Sprintf("%s://%s", scheme, host)
		resp, err := httpClient.Get(url)
		if err != nil {
			continue
		}
		resp.Body.Close()

		port := 80
		if scheme == "https" {
			port = 443
		}

		results = append(results, asset.Asset{
			Host:       host,
			Port:       &port,
			Protocol:   scheme,
			Service:    "http",
			Version:    fmt.Sprintf("HTTP %d", resp.StatusCode),
			Technology: resp.Header.Get("Server"),
		})
	}
	return results
}

// discoverSubdomains queries crt.sh (public Certificate Transparency log
// search) for certificates issued to *.domain — a passive technique that
// never sends traffic to the target itself.
func discoverSubdomains(ctx context.Context, domain string) ([]string, error) {
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entries []struct {
		NameValue string `json:"name_value"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, err
	}

	seen := make(map[string]bool)
	var names []string
	for _, e := range entries {
		for _, name := range strings.Split(e.NameValue, "\n") {
			name = strings.TrimSpace(strings.ToLower(name))
			if name == "" || strings.Contains(name, "*") || seen[name] {
				continue
			}
			seen[name] = true
			names = append(names, name)
		}
	}
	return names, nil
}

func dedupe(items []string) []string {
	seen := make(map[string]bool)
	var out []string
	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			out = append(out, item)
		}
	}
	return out
}
