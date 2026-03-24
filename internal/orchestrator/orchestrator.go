package orchestrator

import (
	"context"
	"github.com/davidaparicio/secmux/internal/scanner"
	"golang.org/x/sync/errgroup"
)

// ScannerResult holds findings and any error for a single scanner run.
type ScannerResult struct {
	Scanner  string
	Findings []scanner.Finding
	Err      error
}

// Result is the aggregated output of all scanner runs.
type Result struct {
	Findings       []scanner.Finding
	ScannerResults []ScannerResult
	Errors         map[string]string
}

// Orchestrator runs multiple scanners in parallel.
type Orchestrator struct {
	scanners []scanner.Scanner
}

// New creates an Orchestrator for the given scanners.
func New(scanners []scanner.Scanner) *Orchestrator {
	return &Orchestrator{scanners: scanners}
}

// Run executes all scanners against path concurrently and aggregates results.
func (o *Orchestrator) Run(ctx context.Context, path string) (Result, error) {
	type item struct {
		name     string
		findings []scanner.Finding
		err      error
	}

	results := make([]item, len(o.scanners))

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(len(o.scanners))

	for i, s := range o.scanners {
		i, s := i, s
		g.Go(func() error {
			f, err := s.Scan(ctx, path)
			results[i] = item{name: s.Name(), findings: f, err: err}
			return nil // never propagate scanner errors; collect them
		})
	}
	_ = g.Wait()

	var result Result
	result.Errors = make(map[string]string)

	for _, r := range results {
		sr := ScannerResult{Scanner: r.name, Findings: r.findings, Err: r.err}
		result.ScannerResults = append(result.ScannerResults, sr)
		if r.err != nil {
			result.Errors[r.name] = r.err.Error()
		}
		result.Findings = append(result.Findings, r.findings...)
	}
	return result, nil
}
