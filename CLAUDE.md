# secmux — Claude Code Project Instructions

**secmux** orchestrates multiple secret-scanning tools in parallel, normalizes their output into a unified schema, and exports SARIF for GitHub Code Scanning.

## Project Overview

```
cmd/secmux/        entry point (version via ldflags)
internal/
  cli/             cobra commands (root, scan)
  scanner/         Scanner interface + 5 implementations
  orchestrator/    parallel execution via errgroup
  formatter/       JSON, SARIF 2.1.0, table output
testdata/          fixture files for unit tests
```

**Supported scanners:** gitleaks, trufflehog, detect-secrets, git-secrets, ggshield

## Development Commands

```bash
go build ./cmd/secmux               # Build binary
go test ./...                       # Unit tests
go test -tags integration ./...     # Integration tests (requires scanner binaries)
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
goreleaser build --snapshot --clean # Cross-compile
```

## Key Conventions

- **Scanner interface**: all scanners implement the `Scanner` interface in `internal/scanner/scanner.go`
- **Auto-detection**: scanners are skipped if the binary is not installed
- **Parallel execution**: `internal/orchestrator/` uses `golang.org/x/sync/errgroup`
- **Output formats**: table (default), json, sarif — handled in `internal/formatter/`
- **Exit codes**: 0 = no secrets, 1 = secrets found or error
- **ggshield**: requires `GITGUARDIAN_API_KEY` env var

## Testing

- Unit tests live alongside source files (`*_test.go`)
- Integration tests use `//go:build integration` tag
- Fixture files for unit tests live in `testdata/`
- Always run `go test ./...` before committing

## Architecture Decisions

- Scanners run concurrently via errgroup; individual scanner failures are collected and reported without aborting the run
- SARIF output follows the 2.1.0 spec for GitHub Code Scanning compatibility
- Version, commit, and date are injected via ldflags at build time (see `.goreleaser.yaml`)
