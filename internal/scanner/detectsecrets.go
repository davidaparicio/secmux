package scanner

import (
	"context"
	"encoding/json"
	"os/exec"
	"path/filepath"
)

// DetectSecretsScanner runs: detect-secrets scan --all-files .
// It runs with Dir=path so relative file paths in the output are anchored to
// the scan root. --all-files ensures untracked git files are included.
type DetectSecretsScanner struct{}

func (d *DetectSecretsScanner) Name() string { return "detect-secrets" }

func (d *DetectSecretsScanner) IsAvailable() bool {
	_, err := exec.LookPath("detect-secrets")
	return err == nil
}

// detectSecretsBaseline is the JSON structure returned by detect-secrets scan.
type detectSecretsBaseline struct {
	Results map[string][]struct {
		LineNumber int    `json:"line_number"`
		Type       string `json:"type"`
	} `json:"results"`
}

func (d *DetectSecretsScanner) Scan(ctx context.Context, path string) ([]Finding, error) {
	cmd := exec.CommandContext(ctx, "detect-secrets", "scan", "--all-files", ".")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return parseDetectSecretsJSON(out, path)
}

func parseDetectSecretsJSON(data []byte, basePath string) ([]Finding, error) {
	var baseline detectSecretsBaseline
	if err := json.Unmarshal(data, &baseline); err != nil {
		return nil, err
	}
	var findings []Finding
	for file, secrets := range baseline.Results {
		abs := filepath.Join(basePath, file)
		for _, s := range secrets {
			findings = append(findings, Finding{
				Scanner:     "detect-secrets",
				File:        abs,
				Line:        s.LineNumber,
				Rule:        s.Type,
				Description: s.Type,
				Severity:    SeverityHigh,
			})
		}
	}
	return findings, nil
}
