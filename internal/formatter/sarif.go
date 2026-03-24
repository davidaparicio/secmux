package formatter

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/davidaparicio/secmux/internal/orchestrator"
	"github.com/davidaparicio/secmux/internal/scanner"
)

// sarifOutput is a minimal SARIF 2.1.0 envelope.
// We hand-roll the struct to avoid a heavy dependency for a simple subset.
type sarifOutput struct {
	Version string     `json:"version"`
	Schema  string     `json:"$schema"`
	Runs    []sarifRun `json:"runs"`
}

type sarifRun struct {
	Tool    sarifTool     `json:"tool"`
	Results []sarifResult `json:"results"`
}

type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

type sarifDriver struct {
	Name    string      `json:"name"`
	Version string      `json:"version"`
	Rules   []sarifRule `json:"rules"`
}

type sarifRule struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	ShortDescription sarifMessage      `json:"shortDescription"`
	Properties       map[string]string `json:"properties,omitempty"`
}

type sarifResult struct {
	RuleID              string            `json:"ruleId"`
	Level               string            `json:"level"`
	Message             sarifMessage      `json:"message"`
	Locations           []sarifLocation   `json:"locations"`
	PartialFingerprints map[string]string `json:"partialFingerprints,omitempty"`
}

type sarifMessage struct {
	Text string `json:"text"`
}

type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
	Region           sarifRegion           `json:"region"`
}

type sarifArtifactLocation struct {
	URI string `json:"uri"`
}

type sarifRegion struct {
	StartLine int `json:"startLine"`
}

// SARIFFormatter formats results as SARIF 2.1.0.
type SARIFFormatter struct{}

func NewSARIF() Formatter { return &SARIFFormatter{} }

func (s *SARIFFormatter) Format(result orchestrator.Result) ([]byte, error) {
	rulesSeen := map[string]bool{}
	var rules []sarifRule
	var results []sarifResult

	for _, f := range result.Findings {
		ruleID := fmt.Sprintf("%s/%s", f.Scanner, f.Rule)
		if !rulesSeen[ruleID] {
			rulesSeen[ruleID] = true
			rules = append(rules, sarifRule{
				ID:               ruleID,
				Name:             f.Rule,
				ShortDescription: sarifMessage{Text: f.Description},
				Properties:       map[string]string{"scanner": f.Scanner},
			})
		}

		results = append(results, sarifResult{
			RuleID:  ruleID,
			Level:   sarifLevel(f.Severity),
			Message: sarifMessage{Text: f.Description},
			Locations: []sarifLocation{{
				PhysicalLocation: sarifPhysicalLocation{
					ArtifactLocation: sarifArtifactLocation{URI: f.File},
					Region:           sarifRegion{StartLine: f.Line},
				},
			}},
			PartialFingerprints: map[string]string{
				"primaryLocationLineHash": lineHash(f.File, f.Line, f.Rule),
			},
		})
	}

	out := sarifOutput{
		Version: "2.1.0",
		Schema:  "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
		Runs: []sarifRun{{
			Tool: sarifTool{
				Driver: sarifDriver{
					Name:    "secmux",
					Version: "dev",
					Rules:   rules,
				},
			},
			Results: results,
		}},
	}
	return json.MarshalIndent(out, "", "  ")
}

func sarifLevel(sev scanner.Severity) string {
	switch sev {
	case scanner.SeverityCritical, scanner.SeverityHigh:
		return "error"
	case scanner.SeverityMedium:
		return "warning"
	default:
		return "note"
	}
}

func lineHash(file string, line int, rule string) string {
	h := sha256.Sum256([]byte(fmt.Sprintf("%s:%d:%s", file, line, rule)))
	return fmt.Sprintf("%x", h[:8])
}
