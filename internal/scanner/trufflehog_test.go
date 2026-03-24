package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTrufflehogScanner_Name(t *testing.T) {
	s := &TrufflehogScanner{}
	if s.Name() != "trufflehog" {
		t.Errorf("got %q, want trufflehog", s.Name())
	}
}

func TestTrufflehogScanner_IsAvailable(t *testing.T) {
	s := &TrufflehogScanner{}
	_ = s.IsAvailable()
}

func TestParseTrufflehogJSONLines_WithFindings(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "trufflehog", "findings.jsonl"))
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	findings, err := parseTrufflehogJSONLines(data)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 2 {
		t.Fatalf("got %d findings, want 2", len(findings))
	}
	f := findings[0]
	if f.Scanner != "trufflehog" {
		t.Errorf("got scanner %q, want trufflehog", f.Scanner)
	}
	if f.Rule != "AWS" {
		t.Errorf("got rule %q, want AWS", f.Rule)
	}
	if f.Line != 7 {
		t.Errorf("got line %d, want 7", f.Line)
	}
	if f.Severity != SeverityCritical {
		t.Errorf("got severity %q, want critical", f.Severity)
	}
}

func TestParseTrufflehogJSONLines_Empty(t *testing.T) {
	findings, err := parseTrufflehogJSONLines([]byte(""))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings, want 0", len(findings))
	}
}

func TestParseTrufflehogJSONLines_SkipsInvalidLines(t *testing.T) {
	data := []byte("{\"DetectorName\":\"AWS\",\"SourceMetadata\":{\"Data\":{\"Filesystem\":{\"file\":\"x.go\",\"line\":1}}},\"Raw\":\"secret\"}\nnot-json\n")
	findings, err := parseTrufflehogJSONLines(data)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 1 {
		t.Errorf("got %d findings, want 1 (bad line should be skipped)", len(findings))
	}
}

var _ Scanner = (*TrufflehogScanner)(nil)
