package scanner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDetectSecretsScanner_Name(t *testing.T) {
	s := &DetectSecretsScanner{}
	if s.Name() != "detect-secrets" {
		t.Errorf("got %q, want detect-secrets", s.Name())
	}
}

func TestDetectSecretsScanner_IsAvailable(t *testing.T) {
	s := &DetectSecretsScanner{}
	_ = s.IsAvailable()
}

func TestParseDetectSecretsJSON_WithFindings(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "detect-secrets", "baseline.json"))
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	findings, err := parseDetectSecretsJSON(data, "/repo")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 2 {
		t.Fatalf("got %d findings, want 2", len(findings))
	}
	for _, f := range findings {
		if f.Scanner != "detect-secrets" {
			t.Errorf("got scanner %q, want detect-secrets", f.Scanner)
		}
		if f.Severity != SeverityHigh {
			t.Errorf("got severity %q, want high", f.Severity)
		}
		if !strings.HasPrefix(f.File, "/repo/") {
			t.Errorf("file %q should be prefixed with /repo/", f.File)
		}
	}
}

func TestParseDetectSecretsJSON_InvalidJSON(t *testing.T) {
	_, err := parseDetectSecretsJSON([]byte("not json"), "/repo")
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestParseDetectSecretsJSON_Empty(t *testing.T) {
	findings, err := parseDetectSecretsJSON([]byte(`{"results":{}}`), "/repo")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings, want 0", len(findings))
	}
}

var _ Scanner = (*DetectSecretsScanner)(nil)
