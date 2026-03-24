package formatter_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/davidaparicio/secmux/internal/formatter"
	"github.com/davidaparicio/secmux/internal/orchestrator"
	"github.com/davidaparicio/secmux/internal/scanner"
)

var sampleResult = orchestrator.Result{
	Findings: []scanner.Finding{
		{
			Scanner:     "gitleaks",
			File:        "config/aws.go",
			Line:        12,
			Rule:        "aws-access-key",
			Description: "AWS Access Key",
			Secret:      "AKIAIOSFODNN7EXAMPLE",
			Severity:    scanner.SeverityHigh,
		},
		{
			Scanner:     "trufflehog",
			File:        "secrets/env.go",
			Line:        7,
			Rule:        "AWS",
			Description: "AWS",
			Secret:      "AKIAIOSFODNN7EXAMPLE",
			Severity:    scanner.SeverityCritical,
		},
	},
	Errors: map[string]string{},
}

func TestJSONFormatter_OutputIsValidJSON(t *testing.T) {
	f := formatter.NewJSON()
	out, err := f.Format(sampleResult)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(out, &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v\n%s", err, out)
	}
}

func TestJSONFormatter_CountMatchesFindings(t *testing.T) {
	f := formatter.NewJSON()
	out, _ := f.Format(sampleResult)

	var parsed struct {
		Total    int `json:"total"`
		Findings []struct {
			Scanner string `json:"scanner"`
		} `json:"findings"`
	}
	if err := json.Unmarshal(out, &parsed); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if parsed.Total != 2 {
		t.Errorf("got total=%d, want 2", parsed.Total)
	}
	if len(parsed.Findings) != 2 {
		t.Errorf("got %d findings, want 2", len(parsed.Findings))
	}
}

func TestSARIFFormatter_OutputIsValidJSON(t *testing.T) {
	f := formatter.NewSARIF()
	out, err := f.Format(sampleResult)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(out, &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v\n%s", err, out)
	}
}

func TestSARIFFormatter_HasExpectedVersion(t *testing.T) {
	f := formatter.NewSARIF()
	out, _ := f.Format(sampleResult)

	var parsed struct {
		Version string `json:"version"`
		Runs    []struct {
			Results []struct {
				RuleID string `json:"ruleId"`
			} `json:"results"`
		} `json:"runs"`
	}
	if err := json.Unmarshal(out, &parsed); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if parsed.Version != "2.1.0" {
		t.Errorf("got version %q, want 2.1.0", parsed.Version)
	}
	if len(parsed.Runs[0].Results) != 2 {
		t.Errorf("got %d SARIF results, want 2", len(parsed.Runs[0].Results))
	}
}

func TestTableFormatter_ContainsScannerNames(t *testing.T) {
	f := formatter.NewTable()
	out, err := f.Format(sampleResult)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}
	s := string(out)
	if !strings.Contains(s, "gitleaks") {
		t.Error("table output missing 'gitleaks'")
	}
	if !strings.Contains(s, "trufflehog") {
		t.Error("table output missing 'trufflehog'")
	}
	if !strings.Contains(s, "2 finding(s) total") {
		t.Errorf("table output missing findings summary\n%s", s)
	}
}

func TestTableFormatter_WithErrors(t *testing.T) {
	f := formatter.NewTable()
	result := orchestrator.Result{
		Findings: sampleResult.Findings,
		Errors:   map[string]string{"broken-scanner": "binary not found"},
	}
	out, err := f.Format(result)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}
	if !strings.Contains(string(out), "broken-scanner") {
		t.Error("table output should include scanner error")
	}
}

func TestTableFormatter_LongFilename(t *testing.T) {
	f := formatter.NewTable()
	longPath := strings.Repeat("a", 80)
	result := orchestrator.Result{
		Findings: []scanner.Finding{
			{Scanner: "test", File: longPath, Line: 1, Rule: "r", Severity: scanner.SeverityLow},
		},
		Errors: map[string]string{},
	}
	out, err := f.Format(result)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}
	if len(out) == 0 {
		t.Error("expected non-empty table output")
	}
}

func TestSARIFFormatter_AllSeverityLevels(t *testing.T) {
	f := formatter.NewSARIF()
	result := orchestrator.Result{
		Findings: []scanner.Finding{
			{Scanner: "s", File: "f", Line: 1, Rule: "r1", Severity: scanner.SeverityCritical},
			{Scanner: "s", File: "f", Line: 2, Rule: "r2", Severity: scanner.SeverityHigh},
			{Scanner: "s", File: "f", Line: 3, Rule: "r3", Severity: scanner.SeverityMedium},
			{Scanner: "s", File: "f", Line: 4, Rule: "r4", Severity: scanner.SeverityLow},
			{Scanner: "s", File: "f", Line: 5, Rule: "r5", Severity: scanner.SeverityInfo},
		},
		Errors: map[string]string{},
	}
	out, err := f.Format(result)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}
	s := string(out)
	if !strings.Contains(s, "error") {
		t.Error("expected 'error' level in SARIF output")
	}
	if !strings.Contains(s, "warning") {
		t.Error("expected 'warning' level in SARIF output")
	}
	if !strings.Contains(s, "note") {
		t.Error("expected 'note' level in SARIF output")
	}
}

func TestTableFormatter_AllSeverityColors(t *testing.T) {
	f := formatter.NewTable()
	result := orchestrator.Result{
		Findings: []scanner.Finding{
			{Scanner: "s", File: "f", Line: 1, Rule: "r", Severity: scanner.SeverityCritical},
			{Scanner: "s", File: "f", Line: 2, Rule: "r", Severity: scanner.SeverityHigh},
			{Scanner: "s", File: "f", Line: 3, Rule: "r", Severity: scanner.SeverityMedium},
			{Scanner: "s", File: "f", Line: 4, Rule: "r", Severity: scanner.SeverityLow},
			{Scanner: "s", File: "f", Line: 5, Rule: "r", Severity: scanner.SeverityInfo},
		},
		Errors: map[string]string{},
	}
	out, err := f.Format(result)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}
	if len(out) == 0 {
		t.Error("expected non-empty output")
	}
}

func TestJSONFormatter_EmptyResult(t *testing.T) {
	f := formatter.NewJSON()
	empty := orchestrator.Result{Errors: map[string]string{}}
	out, err := f.Format(empty)
	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}
	var parsed struct {
		Total int `json:"total"`
	}
	if err := json.Unmarshal(out, &parsed); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if parsed.Total != 0 {
		t.Errorf("got total=%d, want 0", parsed.Total)
	}
}
