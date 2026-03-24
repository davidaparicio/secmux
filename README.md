# secmux

**secmux** orchestrates multiple secret-scanning tools in parallel, normalizes their output into a unified schema, and exports SARIF for GitHub Code Scanning — all from a single binary.

## Why

Each secret scanner has different strengths, output formats, and blind spots. Running them one by one is slow and produces incompatible outputs. secmux runs them all concurrently, merges the results, and gives you a single report.

## Supported scanners

| Scanner                                                         | Install                    | Notes                        |
| --------------------------------------------------------------- | -------------------------- | ---------------------------- |
| [gitleaks](https://github.com/gitleaks/gitleaks)               | `brew install gitleaks`    |                              |
| [trufflehog](https://github.com/trufflesecurity/trufflehog)    | `brew install trufflehog`  |                              |
| [detect-secrets](https://github.com/Yelp/detect-secrets)       | `pip install detect-secrets` |                            |
| [git-secrets](https://github.com/awslabs/git-secrets)          | `brew install git-secrets` |                              |
| [ggshield](https://github.com/GitGuardian/ggshield)            | `pip install ggshield`     | Requires `GITGUARDIAN_API_KEY` |

secmux automatically detects which scanners are installed and skips unavailable ones. You can also select scanners explicitly with `--scanners`.

## Installation

```bash
go install github.com/davidaparicio/secmux/cmd/secmux@latest
```

Or download a pre-built binary from the [releases page](https://github.com/davidaparicio/secmux/releases).

## Usage

```bash
# Scan current directory (runs all available scanners)
secmux scan .

# Scan a specific path
secmux scan /path/to/repo

# Select specific scanners
secmux scan . --scanners gitleaks,trufflehog

# Output as JSON
secmux scan . --format json

# Output as SARIF (for GitHub Code Scanning)
secmux scan . --format sarif > results.sarif

# Custom timeout
secmux scan . --timeout 2m

# Show scanner errors (verbose mode)
secmux scan . --verbose
```

### Exit codes

| Code | Meaning                    |
| ---- | -------------------------- |
| `0`  | No secrets found           |
| `1`  | Secrets found (or error)   |

### Output formats

| Format  | Description                            |
| ------- | -------------------------------------- |
| `table` | Human-readable table (default)         |
| `json`  | Unified JSON with all findings         |
| `sarif` | SARIF 2.1.0 for GitHub Code Scanning   |

### Upload SARIF to GitHub

```bash
secmux scan . --format sarif > results.sarif
gh api repos/{owner}/{repo}/code-scanning/sarifs \
  --method POST \
  -f commit_sha=$(git rev-parse HEAD) \
  -f ref=$(git symbolic-ref HEAD) \
  -f sarif=$(gzip -c results.sarif | base64)
```

## Docker

secmux ships as a multi-arch Docker image (`linux/amd64`, `linux/arm64`) with
gitleaks, trufflehog, detect-secrets, ggshield, and git-secrets pre-installed —
no setup required.

```bash
# Scan current directory
docker run --rm -v "$(pwd):/scan" ghcr.io/davidaparicio/secmux scan /scan

# JSON output
docker run --rm -v "$(pwd):/scan" ghcr.io/davidaparicio/secmux scan /scan --format json

# SARIF for GitHub Code Scanning
docker run --rm -v "$(pwd):/scan" ghcr.io/davidaparicio/secmux scan /scan --format sarif \
  > results.sarif

# With ggshield (needs API key)
docker run --rm -v "$(pwd):/scan" \
  -e GITGUARDIAN_API_KEY="$GITGUARDIAN_API_KEY" \
  ghcr.io/davidaparicio/secmux scan /scan
```

### Build locally

```bash
docker build \
  --build-arg VERSION=$(git describe --tags --always) \
  --build-arg COMMIT=$(git rev-parse --short HEAD) \
  -t secmux .
```

### Bundled tool versions

| Tool | Version |
|---|---|
| gitleaks | 8.21.2 |
| trufflehog | 3.88.1 |
| detect-secrets | latest (pip) |
| ggshield | latest (pip) |
| git-secrets | latest (git) |

Override at build time:

```bash
docker build \
  --build-arg GITLEAKS_VERSION=8.22.0 \
  --build-arg TRUFFLEHOG_VERSION=3.89.0 \
  -t secmux .
```

## GitHub Actions

```yaml
- name: Run secmux
  run: |
    docker run --rm -v "${{ github.workspace }}:/scan" \
      ghcr.io/davidaparicio/secmux scan /scan --format sarif > results.sarif

- name: Upload SARIF
  uses: github/codeql-action/upload-sarif@v3
  with:
    sarif_file: results.sarif
```

Or without Docker:

```yaml
- name: Run secmux
  run: |
    go install github.com/davidaparicio/secmux/cmd/secmux@latest
    secmux scan . --format sarif > results.sarif

- name: Upload SARIF
  uses: github/codeql-action/upload-sarif@v3
  with:
    sarif_file: results.sarif
```

## Development

```bash
# Build
go build ./cmd/secmux

# Run tests
go test ./...

# Run integration tests (requires scanner binaries)
go test -tags integration ./...

# Check test coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Cross-compile (requires goreleaser)
goreleaser build --snapshot --clean
```

## Architecture

```text
cmd/secmux/          entry point (version via ldflags)
internal/
  cli/               cobra commands (root, scan)
  scanner/           Scanner interface + 5 implementations
  orchestrator/      parallel execution via errgroup
  formatter/         JSON, SARIF 2.1.0, table output
testdata/            fixture files for unit tests
```

## License

MIT
