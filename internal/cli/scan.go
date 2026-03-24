package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/davidaparicio/secmux/internal/formatter"
	"github.com/davidaparicio/secmux/internal/orchestrator"
	"github.com/davidaparicio/secmux/internal/scanner"
)

func newScanCmd() *cobra.Command {
	var (
		scanners []string
		format   string
		timeout  time.Duration
	)

	cmd := &cobra.Command{
		Use:   "scan <path>",
		Short: "Run secret scanners against a path",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			registry := scanner.DefaultRegistry()

			var active []scanner.Scanner
			if len(scanners) > 0 {
				for _, name := range scanners {
					s, ok := registry.Get(name)
					if !ok {
						return fmt.Errorf("unknown scanner: %s", name)
					}
					active = append(active, s)
				}
			} else {
				active = registry.Available()
			}

			if len(active) == 0 {
				fmt.Fprintln(os.Stderr, "no scanners available (install gitleaks, trufflehog, detect-secrets, or ggshield)")
				return nil
			}

if verbose {
				names := make([]string, len(active))
				for i, s := range active {
					names[i] = s.Name()
				}
				fmt.Fprintf(os.Stderr, "starting scanners: %s\n", strings.Join(names, ", "))
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			orch := orchestrator.New(active)
			result, err := orch.Run(ctx, path)
			if err != nil {
				return err
			}

			if verbose {
				for _, sr := range result.ScannerResults {
					if sr.Err != nil {
						fmt.Fprintf(os.Stderr, "  %-20s error: %s\n", sr.Scanner, sr.Err)
					} else {
						fmt.Fprintf(os.Stderr, "  %-20s %d finding(s)\n", sr.Scanner, len(sr.Findings))
					}
				}
			}

			var f formatter.Formatter
			switch strings.ToLower(format) {
			case "json":
				f = formatter.NewJSON()
			case "sarif":
				f = formatter.NewSARIF()
			case "table":
				f = formatter.NewTable()
			default:
				return fmt.Errorf("unknown format: %s (use json, sarif, or table)", format)
			}

			out, err := f.Format(result)
			if err != nil {
				return err
			}
			fmt.Print(string(out))

			if verbose && len(result.Errors) > 0 {
				enc := json.NewEncoder(os.Stderr)
				enc.SetIndent("", "  ")
				_ = enc.Encode(result.Errors)
			}

			if len(result.Findings) > 0 {
				os.Exit(1)
			}
			return nil
		},
	}

	cmd.Flags().StringSliceVarP(&scanners, "scanners", "s", nil, "comma-separated list of scanners (default: all available)")
	cmd.Flags().StringVarP(&format, "format", "f", "table", "output format: json, sarif, table")
	cmd.Flags().DurationVarP(&timeout, "timeout", "T", 5*time.Minute, "scan timeout")
	return cmd
}
