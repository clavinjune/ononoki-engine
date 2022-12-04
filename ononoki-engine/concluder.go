package ononoki_engine

// Concluder returns Conclusion from the given data
type Concluder interface {
	Conclude(map[string]any) (*Conclusion, error)
}

// ConcluderFunc an adapter to Concluder
type ConcluderFunc func(map[string]any) (*Conclusion, error)

func (c ConcluderFunc) Conclude(m map[string]any) (*Conclusion, error) {
	return c(m)
}
