package scanner

import (
	"context"
	"testing"
)

func TestRegistry_RegisterAndGet(t *testing.T) {
	r := NewRegistry()
	g := &GitleaksScanner{}
	r.Register(g)

	got, ok := r.Get("gitleaks")
	if !ok {
		t.Fatal("expected to find gitleaks scanner")
	}
	if got.Name() != "gitleaks" {
		t.Errorf("got name %q, want %q", got.Name(), "gitleaks")
	}
}

func TestRegistry_GetUnknown(t *testing.T) {
	r := NewRegistry()
	_, ok := r.Get("does-not-exist")
	if ok {
		t.Fatal("expected not to find unknown scanner")
	}
}

func TestRegistry_Available_FiltersUnavailable(t *testing.T) {
	r := NewRegistry()
	r.Register(&stubScanner{name: "unavailable", available: false})
	r.Register(&stubScanner{name: "available", available: true})

	avail := r.Available()
	if len(avail) != 1 {
		t.Fatalf("got %d available scanners, want 1", len(avail))
	}
	if avail[0].Name() != "available" {
		t.Errorf("got %q, want available", avail[0].Name())
	}
}

func TestRegistry_All(t *testing.T) {
	r := NewRegistry()
	r.Register(&stubScanner{name: "a"})
	r.Register(&stubScanner{name: "b"})

	if len(r.All()) != 2 {
		t.Errorf("got %d scanners, want 2", len(r.All()))
	}
}

func TestDefaultRegistry_HasExpectedScanners(t *testing.T) {
	r := DefaultRegistry()
	expected := []string{"gitleaks", "trufflehog", "detect-secrets", "git-secrets", "ggshield"}
	for _, name := range expected {
		if _, ok := r.Get(name); !ok {
			t.Errorf("default registry missing scanner %q", name)
		}
	}
}

// stubScanner satisfies Scanner for test use.
type stubScanner struct {
	name      string
	available bool
	findings  []Finding
	err       error
}

func (s *stubScanner) Name() string      { return s.name }
func (s *stubScanner) IsAvailable() bool { return s.available }
func (s *stubScanner) Scan(_ context.Context, _ string) ([]Finding, error) {
	return s.findings, s.err
}
