package formatter

import (
	"encoding/json"

	"github.com/davidaparicio/secmux/internal/orchestrator"
)

// Formatter converts a scan result to bytes.
type Formatter interface {
	Format(result orchestrator.Result) ([]byte, error)
}

// jsonOutput is the envelope for JSON format.
type jsonOutput struct {
	Total    int               `json:"total"`
	Findings []jsonFinding     `json:"findings"`
	Errors   map[string]string `json:"errors,omitempty"`
}

type jsonFinding struct {
	Scanner     string `json:"scanner"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	Secret      string `json:"secret,omitempty"`
}

// JSONFormatter formats results as JSON.
type JSONFormatter struct{}

func NewJSON() Formatter { return &JSONFormatter{} }

func (j *JSONFormatter) Format(result orchestrator.Result) ([]byte, error) {
	out := jsonOutput{
		Total:  len(result.Findings),
		Errors: result.Errors,
	}
	for _, f := range result.Findings {
		out.Findings = append(out.Findings, jsonFinding{
			Scanner:     f.Scanner,
			File:        f.File,
			Line:        f.Line,
			Rule:        f.Rule,
			Description: f.Description,
			Severity:    string(f.Severity),
			Secret:      f.Secret,
		})
	}
	return json.MarshalIndent(out, "", "  ")
}
