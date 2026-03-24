package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootCmd_Version(t *testing.T) {
	cmd := newRootCmd("1.2.3", "abc123")
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--version"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "1.2.3") {
		t.Errorf("version output missing '1.2.3': %q", out)
	}
	if !strings.Contains(out, "abc123") {
		t.Errorf("version output missing commit 'abc123': %q", out)
	}
}

func TestRootCmd_Help(t *testing.T) {
	cmd := newRootCmd("dev", "none")
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error: %v", err)
	}
	if !strings.Contains(buf.String(), "secmux") {
		t.Errorf("help output missing 'secmux': %q", buf.String())
	}
}

func TestScanCmd_Help(t *testing.T) {
	cmd := newRootCmd("dev", "none")
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"scan", "--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "scan") {
		t.Errorf("scan help missing 'scan': %q", out)
	}
}

func TestScanCmd_MissingPath(t *testing.T) {
	cmd := newRootCmd("dev", "none")
	cmd.SetArgs([]string{"scan"})
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error when path argument is missing")
	}
}

func TestScanCmd_UnknownScanner(t *testing.T) {
	cmd := newRootCmd("dev", "none")
	errBuf := &bytes.Buffer{}
	cmd.SetErr(errBuf)
	cmd.SetArgs([]string{"scan", "--scanners", "no-such-tool", "/tmp"})
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error for unknown scanner name")
	}
}

func TestScanCmd_UnknownFormat(t *testing.T) {
	// Specify a scanner by name so the command reaches the format-check branch
	// even when no tool binaries are installed.
	cmd := newRootCmd("dev", "none")
	cmd.SetArgs([]string{"scan", "--scanners", "gitleaks", "--format", "xml", "/tmp"})
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error for unknown format")
	}
}
