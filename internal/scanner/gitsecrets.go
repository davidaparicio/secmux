package scanner

import (
	"bufio"
	"bytes"
	"context"
	"os/exec"
	"strconv"
	"strings"
)

// GitSecretsScanner runs: git secrets --scan -r <path>
type GitSecretsScanner struct{}

func (g *GitSecretsScanner) Name() string { return "git-secrets" }

func (g *GitSecretsScanner) IsAvailable() bool {
	_, gitErr := exec.LookPath("git")
	_, gsErr := exec.LookPath("git-secrets")
	// git-secrets can be invoked as "git secrets" (via git alias) or standalone
	if gitErr != nil {
		return false
	}
	return gsErr == nil
}

// git secrets prints matches to stderr in the form: <file>:<line>:<match>
func (g *GitSecretsScanner) Scan(ctx context.Context, path string) ([]Finding, error) {
	cmd := exec.CommandContext(ctx, "git", "secrets", "--scan", "-r", path)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_ = cmd.Run() // non-zero exit when findings exist

	var findings []Finding
	sc := bufio.NewScanner(&stderr)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 3 {
			continue
		}
		lineNum, _ := strconv.Atoi(parts[1])
		findings = append(findings, Finding{
			Scanner:     g.Name(),
			File:        parts[0],
			Line:        lineNum,
			Rule:        "aws-secret",
			Description: strings.TrimSpace(parts[2]),
			Severity:    SeverityCritical,
		})
	}
	return findings, sc.Err()
}
