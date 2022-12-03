package ononoki_engine

// Concluder returns Conclusion from given data
type Concluder interface {
	Conclude(map[string]any) (*Conclusion, error)
}
