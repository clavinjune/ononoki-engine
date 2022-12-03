package ononoki_engine

const (
	// ComparatorEQ compares with ==
	ComparatorEQ Comparator = "COMPARATOR_EQ"

	// ComparatorNE compares with !=
	ComparatorNE Comparator = "COMPARATOR_NE"

	// ComparatorLT compares with <
	ComparatorLT Comparator = "COMPARATOR_LT"

	// ComparatorLE compares with <=
	ComparatorLE Comparator = "COMPARATOR_LE"

	// ComparatorGT compares with >
	ComparatorGT Comparator = "COMPARATOR_GT"

	// ComparatorGE compares with >=
	ComparatorGE Comparator = "COMPARATOR_GE"
)

// Comparator is used for dynamic comparator
type Comparator string

func Compare[T Threshold](c Comparator, a, b T) bool {
	switch c {
	case ComparatorEQ:
		return a == b
	case ComparatorNE:
		return a != b
	case ComparatorLT:
		return a < b
	case ComparatorLE:
		return a <= b
	case ComparatorGT:
		return a > b
	case ComparatorGE:
		return a >= b
	}

	return false
}
