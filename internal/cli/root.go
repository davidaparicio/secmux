package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

func newRootCmd(version, commit string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "secmux",
		Short:   "Orchestrate multiple secret-scanning tools in parallel",
		Version: fmt.Sprintf("%s (commit %s)", version, commit),
	}
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	cmd.AddCommand(newScanCmd())
	return cmd
}

func Execute(version, commit string) {
	if err := newRootCmd(version, commit).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
