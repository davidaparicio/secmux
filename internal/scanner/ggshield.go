package scanner

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
)

// GGShieldScanner runs: ggshield secret scan path -r <path> --json
// Requires GITGUARDIAN_API_KEY env var.
type GGShieldScanner struct{}

func (g *GGShieldScanner) Name() string { return "ggshield" }

func (g *GGShieldScanner) IsAvailable() bool {
	_, err := exec.LookPath("ggshield")
	return err == nil && os.Getenv("GITGUARDIAN_API_KEY") != ""
}

func (g *GGShieldScanner) SkipReason() string {
	if _, err := exec.LookPath("ggshield"); err != nil {
		return ""
	}
	if os.Getenv("GITGUARDIAN_API_KEY") == "" {
		return "ggshield is installed but GITGUARDIAN_API_KEY is not set"
	}
	return ""
}

type ggshieldOutput struct {
	Entities []struct {
		Filename  string `json:"filename"`
		Incidents []struct {
			PolicyBreakCount int    `json:"policy_break_count"`
			Type             string `json:"type"`
			PolicyBreaks     []struct {
				Type  string `json:"type"`
				Match struct {
					Match     string `json:"match"`
					LineStart int    `json:"line_start"`
				} `json:"match"`
			} `json:"policy_breaks"`
		} `json:"incidents"`
	} `json:"entities_with_incidents"`
}

func (g *GGShieldScanner) Scan(ctx context.Context, path string) ([]Finding, error) {
	// -y skips the interactive "N files will be scanned. Continue? [y/N]" prompt.
	cmd := exec.CommandContext(ctx, "ggshield", "secret", "scan", "path", "-r", "-y", path, "--json")
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != 0 && len(out) == 0 {
			return nil, err
		}
		// exit 1 with JSON output = findings present
	}

	if len(out) == 0 {
		return nil, nil
	}
	return parseGGShieldJSON(out)
}

func parseGGShieldJSON(data []byte) ([]Finding, error) {
	var result ggshieldOutput
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	var findings []Finding
	for _, entity := range result.Entities {
		for _, incident := range entity.Incidents {
			for _, pb := range incident.PolicyBreaks {
				findings = append(findings, Finding{
					Scanner:     "ggshield",
					File:        entity.Filename,
					Line:        pb.Match.LineStart,
					Rule:        pb.Type,
					Description: incident.Type,
					Secret:      pb.Match.Match,
					Severity:    SeverityCritical,
				})
			}
		}
	}
	return findings, nil
}
