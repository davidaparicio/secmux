package scanner

import "context"

// Severity levels for findings.
type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityMedium   Severity = "medium"
	SeverityLow      Severity = "low"
	SeverityInfo     Severity = "info"
)

// Finding is a single secret detected by a scanner.
type Finding struct {
	Scanner     string   `json:"scanner"`
	File        string   `json:"file"`
	Line        int      `json:"line"`
	Rule        string   `json:"rule"`
	Description string   `json:"description"`
	Secret      string   `json:"secret,omitempty"`
	Severity    Severity `json:"severity"`
}

// Verifiable is optionally implemented by scanners that support disabling
// external credential verification (e.g. detect-secrets --no-verify).
type Verifiable interface {
	SetNoVerify(bool)
}

// Scanner is the interface every scanner must implement.
type Scanner interface {
	// Name returns the unique identifier for this scanner.
	Name() string
	// IsAvailable reports whether the scanner binary (and any required config) is present.
	IsAvailable() bool
	// Scan runs the scanner against path and returns findings.
	Scan(ctx context.Context, path string) ([]Finding, error)
}

// Registry holds all registered scanners.
type Registry struct {
	scanners []Scanner
}

// NewRegistry creates an empty registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Register adds a scanner to the registry.
func (r *Registry) Register(s Scanner) {
	r.scanners = append(r.scanners, s)
}

// Get returns the scanner with the given name, if any.
func (r *Registry) Get(name string) (Scanner, bool) {
	for _, s := range r.scanners {
		if s.Name() == name {
			return s, true
		}
	}
	return nil, false
}

// Available returns scanners whose IsAvailable() returns true.
func (r *Registry) Available() []Scanner {
	var out []Scanner
	for _, s := range r.scanners {
		if s.IsAvailable() {
			out = append(out, s)
		}
	}
	return out
}

// All returns every registered scanner regardless of availability.
func (r *Registry) All() []Scanner {
	return r.scanners
}

// DefaultRegistry returns a registry pre-populated with all built-in scanners.
func DefaultRegistry() *Registry {
	r := NewRegistry()
	r.Register(&GitleaksScanner{})
	r.Register(&TrufflehogScanner{})
	r.Register(&DetectSecretsScanner{})
	r.Register(&GitSecretsScanner{})
	r.Register(&GGShieldScanner{})
	return r
}
