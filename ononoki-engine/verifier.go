package ononoki_engine

// Verifier verifies whether the given data is valid or not
type Verifier interface {
	Verify(map[string]any) (bool, error)
}

// VerifierFunc an adapter to Verifier
type VerifierFunc func(map[string]any) (bool, error)

func (v VerifierFunc) Verify(m map[string]any) (bool, error) {
	return v(m)
}
