package orchestrator_test

import (
	"context"
	"errors"
	"testing"

	"github.com/davidaparicio/secmux/internal/orchestrator"
	"github.com/davidaparicio/secmux/internal/scanner"
)

func TestOrchestrator_AggregatesFindings(t *testing.T) {
	s1 := &stubScanner{
		name: "s1",
		findings: []scanner.Finding{
			{Scanner: "s1", File: "a.go", Line: 1, Rule: "rule1", Severity: scanner.SeverityHigh},
		},
	}
	s2 := &stubScanner{
		name: "s2",
		findings: []scanner.Finding{
			{Scanner: "s2", File: "b.go", Line: 2, Rule: "rule2", Severity: scanner.SeverityCritical},
		},
	}

	orch := orchestrator.New([]scanner.Scanner{s1, s2})
	result, err := orch.Run(context.Background(), "/tmp/test")
	if err != nil {
		t.Fatalf("Run() error: %v", err)
	}

	if len(result.Findings) != 2 {
		t.Errorf("got %d findings, want 2", len(result.Findings))
	}
}

func TestOrchestrator_CollectsScannerErrors(t *testing.T) {
	failing := &stubScanner{name: "failing", err: errors.New("binary not found")}
	ok := &stubScanner{
		name:     "ok",
		findings: []scanner.Finding{{Scanner: "ok", Rule: "r", Severity: scanner.SeverityLow}},
	}

	orch := orchestrator.New([]scanner.Scanner{failing, ok})
	result, err := orch.Run(context.Background(), "/tmp/test")
	if err != nil {
		t.Fatalf("Run() returned non-nil error; scanner errors should be collected, not propagated")
	}

	if result.Errors["failing"] == "" {
		t.Error("expected error for 'failing' scanner to be recorded")
	}
	if len(result.Findings) != 1 {
		t.Errorf("got %d findings, want 1 (from ok scanner)", len(result.Findings))
	}
}

func TestOrchestrator_EmptyScanners(t *testing.T) {
	orch := orchestrator.New(nil)
	result, err := orch.Run(context.Background(), "/tmp/test")
	if err != nil {
		t.Fatalf("Run() error: %v", err)
	}
	if len(result.Findings) != 0 {
		t.Errorf("got %d findings, want 0", len(result.Findings))
	}
}

func TestOrchestrator_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	slow := &stubScanner{name: "slow"}
	orch := orchestrator.New([]scanner.Scanner{slow})
	// Should not hang.
	_, err := orch.Run(ctx, "/tmp/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// stubScanner satisfies scanner.Scanner for test use.
type stubScanner struct {
	name     string
	findings []scanner.Finding
	err      error
}

func (s *stubScanner) Name() string      { return s.name }
func (s *stubScanner) IsAvailable() bool { return true }
func (s *stubScanner) Scan(_ context.Context, _ string) ([]scanner.Finding, error) {
	return s.findings, s.err
}
