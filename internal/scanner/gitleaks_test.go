package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGitleaksScanner_Name(t *testing.T) {
	g := &GitleaksScanner{}
	if g.Name() != "gitleaks" {
		t.Errorf("got %q, want gitleaks", g.Name())
	}
}

func TestGitleaksScanner_IsAvailable(t *testing.T) {
	g := &GitleaksScanner{}
	_ = g.IsAvailable() // must not panic
}

func TestParseGitleaksJSON_WithFindings(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "gitleaks", "findings.json"))
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	findings, err := parseGitleaksJSON(data)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("got %d findings, want 1", len(findings))
	}
	f := findings[0]
	if f.Scanner != "gitleaks" {
		t.Errorf("got scanner %q, want gitleaks", f.Scanner)
	}
	if f.Rule != "aws-access-key" {
		t.Errorf("got rule %q, want aws-access-key", f.Rule)
	}
	if f.Line != 12 {
		t.Errorf("got line %d, want 12", f.Line)
	}
	if f.Severity != SeverityHigh {
		t.Errorf("got severity %q, want high", f.Severity)
	}
}

func TestParseGitleaksJSON_Empty(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "gitleaks", "empty.json"))
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	findings, err := parseGitleaksJSON(data)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings, want 0", len(findings))
	}
}

func TestParseGitleaksJSON_InvalidJSON(t *testing.T) {
	_, err := parseGitleaksJSON([]byte("not json"))
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

// Compile-time interface check (internal package).
var _ Scanner = (*GitleaksScanner)(nil)
