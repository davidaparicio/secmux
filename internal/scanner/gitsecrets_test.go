package scanner

import (
	"testing"
)

func TestGitSecretsScanner_Name(t *testing.T) {
	s := &GitSecretsScanner{}
	if s.Name() != "git-secrets" {
		t.Errorf("got %q, want git-secrets", s.Name())
	}
}

func TestGitSecretsScanner_IsAvailable(t *testing.T) {
	s := &GitSecretsScanner{}
	_ = s.IsAvailable()
}

var _ Scanner = (*GitSecretsScanner)(nil)
