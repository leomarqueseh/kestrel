package report

import (
	"bytes"
	"html/template"
)

// htmlTemplate renders a self-contained, printable report page — no
// external assets, so it works offline and prints cleanly to PDF from
// any browser (the stand-in for real PDF generation in this MVP).
const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Kestrel Security Assessment Report — {{.Project.Name}}</title>
<style>
  body { font-family: -apple-system, Arial, sans-serif; max-width: 900px; margin: 40px auto; padding: 0 20px; color: #1a1a1a; }
  h1, h2 { border-bottom: 2px solid #1a1a1a; padding-bottom: 8px; }
  .severity-critical { color: #b91c1c; font-weight: bold; }
  .severity-high { color: #c2410c; font-weight: bold; }
  .severity-medium { color: #a16207; font-weight: bold; }
  .severity-low { color: #4d7c0f; }
  .severity-informational { color: #64748b; }
  table { border-collapse: collapse; width: 100%; margin: 16px 0; }
  th, td { border: 1px solid #ddd; padding: 8px; text-align: left; font-size: 14px; }
  th { background: #f3f4f6; }
  .evidence { background: #f9fafb; border-left: 3px solid #1a1a1a; padding: 12px; margin: 8px 0; font-family: monospace; font-size: 13px; white-space: pre-wrap; }
  .badge { display: inline-block; padding: 2px 10px; border-radius: 12px; font-size: 12px; margin-right: 6px; background: #e5e7eb; }
  .finding { margin-bottom: 24px; padding-bottom: 16px; border-bottom: 1px solid #eee; }
</style>
</head>
<body>

<h1>Security Assessment Report</h1>
<p><strong>Project:</strong> {{.Project.Name}}<br>
<strong>Generated:</strong> {{.GeneratedAt.Format "2006-01-02 15:04 UTC"}}</p>

<h2>Executive Summary</h2>
<p>{{.Project.Description}}</p>
<p>{{range $severity, $count := .SeverityCounts}}<span class="badge">{{$severity}}: {{$count}}</span>{{end}}</p>

<h2>Scope</h2>
<table>
<tr><th>Target</th><th>Type</th><th>Authorized</th></tr>
{{range .Targets}}<tr><td>{{.Value}}</td><td>{{.Type}}</td><td>{{if .Authorized}}Yes{{else}}No{{end}}</td></tr>{{end}}
</table>

<h2>Methodology</h2>
<p>Assessment performed via Kestrel's standard workflow: reconnaissance, port enumeration, NVD-correlated vulnerability assessment, and manual validation. Only <strong>confirmed</strong> findings below carry validated evidence; <strong>detected</strong> findings are automated candidates pending analyst review.</p>

<h2>Findings</h2>
{{range .Findings}}
<div class="finding">
  <h3 class="severity-{{.Severity}}">{{.Title}} — [{{.Severity}}] [{{.Status}}]</h3>
  <p>{{.Description}}</p>
  {{if .CVSS}}<p><strong>CVSS:</strong> {{.CVSS}}</p>{{end}}
  <p><strong>Recommendation:</strong> {{.Recommendation}}</p>
</div>
{{end}}

<h2>Evidence</h2>
{{range $findingID, $items := .EvidenceByID}}{{range $items}}
<div class="evidence"><strong>Finding {{$findingID}}</strong>
{{if .Request}}Request: {{.Request}}{{end}}
{{if .Response}}Response: {{.Response}}{{end}}
{{if .Notes}}Notes: {{.Notes}}{{end}}
</div>
{{end}}{{end}}

<h2>Conclusion</h2>
<p>This report reflects the assessment state at generation time. Findings with status "detected" require validation before being treated as confirmed risk.</p>

</body>
</html>`

func RenderHTML(data *Data) ([]byte, error) {
	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
