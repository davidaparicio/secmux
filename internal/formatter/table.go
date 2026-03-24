package formatter

import (
	"bytes"
	"fmt"

	"github.com/davidaparicio/secmux/internal/orchestrator"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// TableFormatter formats results as a human-readable table.
type TableFormatter struct{}

func NewTable() Formatter { return &TableFormatter{} }

func (t *TableFormatter) Format(result orchestrator.Result) ([]byte, error) {
	var buf bytes.Buffer

	tw := table.NewWriter()
	tw.SetOutputMirror(&buf)
	tw.SetStyle(table.StyleLight)
	tw.AppendHeader(table.Row{"Scanner", "Severity", "File", "Line", "Rule", "Description"})

	for _, f := range result.Findings {
		tw.AppendRow(table.Row{
			f.Scanner,
			colorSeverity(string(f.Severity)),
			truncate(f.File, 50),
			f.Line,
			truncate(f.Rule, 30),
			truncate(f.Description, 60),
		})
	}

	tw.Render()

	fmt.Fprintf(&buf, "\n%d finding(s) total\n", len(result.Findings))
	if len(result.Errors) > 0 {
		fmt.Fprintf(&buf, "%d scanner error(s):\n", len(result.Errors))
		for name, msg := range result.Errors {
			fmt.Fprintf(&buf, "  %s: %s\n", name, msg)
		}
	}
	return buf.Bytes(), nil
}

func colorSeverity(sev string) string {
	switch sev {
	case "critical":
		return text.FgHiRed.Sprint(sev)
	case "high":
		return text.FgRed.Sprint(sev)
	case "medium":
		return text.FgYellow.Sprint(sev)
	case "low":
		return text.FgCyan.Sprint(sev)
	default:
		return sev
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return "…" + s[len(s)-max+1:]
}
