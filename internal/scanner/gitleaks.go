package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GitleaksScanner runs: gitleaks dir -f json -r <tmpfile> <path>
type GitleaksScanner struct{}

func (g *GitleaksScanner) Name() string { return "gitleaks" }

func (g *GitleaksScanner) IsAvailable() bool {
	_, err := exec.LookPath("gitleaks")
	return err == nil
}

type gitleaksFinding struct {
	Description string `json:"Description"`
	File        string `json:"File"`
	Line        int    `json:"StartLine"`
	Rule        string `json:"RuleID"`
	Secret      string `json:"Secret"`
}

func (g *GitleaksScanner) Scan(ctx context.Context, path string) ([]Finding, error) {
	tmp, err := os.CreateTemp("", "gitleaks-*.json")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())
	tmp.Close()

	// gitleaks dir exits 0 (no findings) or 1 (findings found).
	// Any other exit code is a real error — capture stderr for diagnostics.
	// Note: `dir` takes the path as a positional argument, not via -s.
	var stderr strings.Builder
	cmd := exec.CommandContext(ctx, "gitleaks", "dir", "-f", "json", "-r", tmp.Name(), path)
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			// exit 1 means findings were found; parse the report below
		} else {
			msg := strings.TrimSpace(stderr.String())
			if msg != "" {
				return nil, fmt.Errorf("gitleaks: %s", msg)
			}
			return nil, err
		}
	}

	data, err := os.ReadFile(tmp.Name())
	if err != nil || len(data) == 0 {
		return nil, nil
	}

	return parseGitleaksJSON(data)
}

func parseGitleaksJSON(data []byte) ([]Finding, error) {
	var raw []gitleaksFinding
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	findings := make([]Finding, 0, len(raw))
	for _, r := range raw {
		findings = append(findings, Finding{
			Scanner:     "gitleaks",
			File:        r.File,
			Line:        r.Line,
			Rule:        r.Rule,
			Description: r.Description,
			Secret:      r.Secret,
			Severity:    SeverityHigh,
		})
	}
	return findings, nil
}
