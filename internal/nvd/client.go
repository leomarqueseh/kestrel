// Package nvd provides a minimal client for the NVD (National Vulnerability
// Database) public CVE search API — no API key required, subject to a
// public rate limit of ~5 requests per rolling 30-second window.
package nvd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://services.nvd.nist.gov/rest/json/cves/2.0"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// CVE is a simplified view of what the NVD returns — just enough to
// build a finding.
type CVE struct {
	ID          string
	Description string
	CVSS        float64
	Severity    string // NVD's own label: CRITICAL/HIGH/MEDIUM/LOW
}

type cveResponse struct {
	Vulnerabilities []struct {
		CVE struct {
			ID           string `json:"id"`
			Descriptions []struct {
				Lang  string `json:"lang"`
				Value string `json:"value"`
			} `json:"descriptions"`
			Metrics struct {
				CvssMetricV31 []struct {
					CvssData struct {
						BaseScore    float64 `json:"baseScore"`
						BaseSeverity string  `json:"baseSeverity"`
					} `json:"cvssData"`
				} `json:"cvssMetricV31"`
			} `json:"metrics"`
		} `json:"cve"`
	} `json:"vulnerabilities"`
}

// SearchByKeyword queries the NVD for CVEs matching free-text keyword
// (e.g. "OpenSSH 7.4") and returns up to limit results.
func SearchByKeyword(ctx context.Context, keyword string, limit int) ([]CVE, error) {
	q := url.Values{}
	q.Set("keywordSearch", keyword)
	q.Set("resultsPerPage", fmt.Sprintf("%d", limit))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("nvd: rate limited, try again later")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nvd: unexpected status %d", resp.StatusCode)
	}

	var parsed cveResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	var results []CVE
	for _, v := range parsed.Vulnerabilities {
		desc := ""
		for _, d := range v.CVE.Descriptions {
			if d.Lang == "en" {
				desc = d.Value
				break
			}
		}

		var score float64
		var severity string
		if len(v.CVE.Metrics.CvssMetricV31) > 0 {
			score = v.CVE.Metrics.CvssMetricV31[0].CvssData.BaseScore
			severity = v.CVE.Metrics.CvssMetricV31[0].CvssData.BaseSeverity
		}

		results = append(results, CVE{
			ID:          v.CVE.ID,
			Description: desc,
			CVSS:        score,
			Severity:    severity,
		})
	}
	return results, nil
}
