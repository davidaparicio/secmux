//go:build integration

package scanner

import (
	"context"
	"os"
	"os/exec"
	"testing"
	"time"
)

// Integration tests for Scan() methods.
// Run with: go test -tags integration ./internal/scanner/...
//
// These tests require the respective scanner binaries to be installed.

func TestGitleaksScanner_Scan_Integration(t *testing.T) {
	if _, err := exec.LookPath("gitleaks"); err != nil {
		t.Skip("gitleaks not installed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := &GitleaksScanner{}
	findings, err := s.Scan(ctx, t.TempDir())
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings in clean dir, want 0", len(findings))
	}
}

func TestTrufflehogScanner_Scan_Integration(t *testing.T) {
	if _, err := exec.LookPath("trufflehog"); err != nil {
		t.Skip("trufflehog not installed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := &TrufflehogScanner{}
	findings, err := s.Scan(ctx, t.TempDir())
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings in clean dir, want 0", len(findings))
	}
}

func TestDetectSecretsScanner_Scan_Integration(t *testing.T) {
	if _, err := exec.LookPath("detect-secrets"); err != nil {
		t.Skip("detect-secrets not installed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := &DetectSecretsScanner{}
	findings, err := s.Scan(ctx, t.TempDir())
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings in clean dir, want 0", len(findings))
	}
}

func TestGGShieldScanner_Scan_Integration(t *testing.T) {
	if _, err := exec.LookPath("ggshield"); err != nil {
		t.Skip("ggshield not installed")
	}
	if os.Getenv("GITGUARDIAN_API_KEY") == "" {
		t.Skip("GITGUARDIAN_API_KEY not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := &GGShieldScanner{}
	findings, err := s.Scan(ctx, t.TempDir())
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings in clean dir, want 0", len(findings))
	}
}

func TestGitSecretsScanner_Scan_Integration(t *testing.T) {
	if _, err := exec.LookPath("git-secrets"); err != nil {
		t.Skip("git-secrets not installed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := &GitSecretsScanner{}
	findings, err := s.Scan(ctx, t.TempDir())
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}
	if len(findings) != 0 {
		t.Errorf("got %d findings in clean dir, want 0", len(findings))
	}
}
