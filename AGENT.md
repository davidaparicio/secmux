# secmux — Agent Instructions

## What is secmux?

A Go CLI tool that runs multiple secret-scanning tools (gitleaks, trufflehog, detect-secrets, git-secrets, ggshield) in parallel, merges their findings into a unified schema, and outputs table, JSON, or SARIF 2.1.0 reports.

## Repository Layout

```
cmd/secmux/main.go          CLI entry point; version injected via ldflags
internal/cli/               cobra root + scan commands
internal/scanner/           Scanner interface + one file per tool
internal/orchestrator/      parallel runner (errgroup)
internal/formatter/         output renderers (table, json, sarif)
testdata/                   JSON fixtures for unit tests
Dockerfile                  multi-arch image (amd64/arm64) with all scanners pre-installed
.goreleaser.yaml            cross-compile and release config
```

## Core Abstractions

| Package | Purpose |
|---|---|
| `scanner.Scanner` | Interface every scanner implements: `Name() string`, `Available() bool`, `Scan(ctx, path) ([]Finding, error)` |
| `scanner.Finding` | Unified finding struct (RuleID, Description, File, Line, Secret) |
| `orchestrator` | Runs all available scanners concurrently, collects errors without aborting |
| `formatter` | Converts `[]Finding` to table/JSON/SARIF output |

## Key Constraints

- **No scanner failure is fatal**: a scanner error is reported but the run continues
- **Auto-detection**: `Available()` checks if the binary exists in PATH
- **ggshield**: needs `GITGUARDIAN_API_KEY`; skip if env var is absent
- **Integration tests**: gated by `//go:build integration`; require real scanner binaries

## Development Workflow

```bash
go build ./cmd/secmux
go test ./...
go test -tags integration ./...
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
goreleaser build --snapshot --clean
```

## Adding a New Scanner

1. Create `internal/scanner/<name>.go` implementing `Scanner`
2. Create `internal/scanner/<name>_test.go` with unit tests using fixtures in `testdata/`
3. Register the scanner in `internal/cli/scan.go`
4. Add the binary to the `Dockerfile`
5. Document it in `README.md`

## Output Formats

| Flag | Description |
|---|---|
| `--format table` | Human-readable table (default) |
| `--format json` | Unified JSON array of findings |
| `--format sarif` | SARIF 2.1.0 for GitHub Code Scanning |

## Exit Codes

| Code | Meaning |
|---|---|
| `0` | No secrets found |
| `1` | Secrets found, or unrecoverable error |
