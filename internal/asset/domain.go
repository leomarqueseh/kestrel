// Package asset stores what reconnaissance and enumeration discover about
// a target: hosts, ports, protocols, and detected technologies.
package asset

import "time"

type Asset struct {
	ID         string    `json:"id"`
	ScanID     string    `json:"scan_id"`
	Host       string    `json:"host"`
	Port       *int      `json:"port,omitempty"`
	Protocol   string    `json:"protocol,omitempty"`
	Service    string    `json:"service,omitempty"`
	Version    string    `json:"version,omitempty"`
	Technology string    `json:"technology,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
