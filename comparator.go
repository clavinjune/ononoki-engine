// Copyright 2022 clavinjune/ononoki-engine
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ononoki

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

// Compare compares the given value and threshold
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
