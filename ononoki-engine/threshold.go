package ononoki_engine

// Threshold limits the Fact.Threshold data type
type Threshold interface {
	~int64 | ~float64 | ~string
}
