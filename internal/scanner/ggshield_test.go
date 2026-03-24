package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGGShieldScanner_Name(t *testing.T) {
	s := &GGShieldScanner{}
	if s.Name() != "ggshield" {
		t.Errorf("got %q, want ggshield", s.Name())
	}
}

func TestGGShieldScanner_IsAvailable_NoAPIKey(t *testing.T) {
	t.Setenv("GITGUARDIAN_API_KEY", "")
	s := &GGShieldScanner{}
	if s.IsAvailable() {
		t.Error("expected IsAvailable=false when GITGUARDIAN_API_KEY is unset")
	}
}

func TestGGShieldScanner_IsAvailable_WithAPIKey(t *testing.T) {
	t.Setenv("GITGUARDIAN_API_KEY", "test-key")
	s := &GGShieldScanner{}
	// IsAvailable also checks binary; can be true or false — just must not panic.
	_ = s.IsAvailable()
}

func TestParseGGShieldJSON_WithFindings(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "ggshield", "findings.json"))
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	findings, err := parseGGShieldJSON(data)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("got %d findings, want 1", len(findings))
	}
	f := findings[0]
	if f.Scanner != "ggshield" {
		t.Errorf("got scanner %q, want ggshield", f.Scanner)
	}
	if f.File != "src/config.go" {
		t.Errorf("got file %q, want src/config.go", f.File)
	}
	if f.Line != 8 {
		t.Errorf("got line %d, want 8", f.Line)
	}
	if f.Severity != SeverityCritical {
		t.Errorf("got severity %q, want critical", f.Severity)
	}
}

func TestParseGGShieldJSON_InvalidJSON(t *testing.T) {
	_, err := parseGGShieldJSON([]byte("not json"))
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestParseGGShieldJSON_NoFindings(t *testing.T) {
	findings, err := parseGGShieldJSON([]byte(`{"entities_with_incidents":[]}`))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings, want 0", len(findings))
	}
}

var _ Scanner = (*GGShieldScanner)(nil)
