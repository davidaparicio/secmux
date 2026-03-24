package scanner

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// TrufflehogScanner runs: trufflehog filesystem <path> --json
type TrufflehogScanner struct{}

func (t *TrufflehogScanner) Name() string { return "trufflehog" }

func (t *TrufflehogScanner) IsAvailable() bool {
	_, err := exec.LookPath("trufflehog")
	return err == nil
}

type trufflehogResult struct {
	DetectorName   string `json:"DetectorName"`
	SourceMetadata struct {
		Data struct {
			Filesystem struct {
				File string `json:"file"`
				Line int64  `json:"line"`
			} `json:"Filesystem"`
		} `json:"Data"`
	} `json:"SourceMetadata"`
	Raw string `json:"Raw"`
}

func (t *TrufflehogScanner) Scan(ctx context.Context, path string) ([]Finding, error) {
	// trufflehog filesystem <path> --json
	// Without --fail: exits 0 regardless of findings; output is JSON lines on stdout.
	// Exit code 183 only applies when --fail is passed.
	// Any other non-zero exit is a real error — capture stderr for diagnostics.
	var stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "trufflehog", "filesystem", path, "--json", "--no-update")
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 183 {
			// findings found (--fail mode) — parse anyway
		} else {
			msg := strings.TrimSpace(stderr.String())
			if msg != "" {
				return nil, fmt.Errorf("trufflehog: %s", msg)
			}
			return nil, err
		}
	}

	return parseTrufflehogJSONLines(out)
}

func parseTrufflehogJSONLines(data []byte) ([]Finding, error) {
	var findings []Finding
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var r trufflehogResult
		if err := json.Unmarshal(line, &r); err != nil {
			continue
		}
		findings = append(findings, Finding{
			Scanner:     "trufflehog",
			File:        r.SourceMetadata.Data.Filesystem.File,
			Line:        int(r.SourceMetadata.Data.Filesystem.Line),
			Rule:        r.DetectorName,
			Description: r.DetectorName,
			Secret:      r.Raw,
			Severity:    SeverityCritical,
		})
	}
	return findings, sc.Err()
}
