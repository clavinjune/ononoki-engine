package ononoki_engine_test

import (
	"fmt"
	ononoki "github.com/clavinjune/ononoki-engine"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	tt := []struct {
		Name       string
		Comparator ononoki.Comparator
		Value      any
		Threshold  any
		Expected   bool
		Type       reflect.Kind
	}{
		{
			Name:       "unknown Comparator",
			Comparator: "",
			Value:      "",
			Threshold:  "",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorEQ(string) false"),
			Comparator: ononoki.ComparatorEQ,
			Value:      "a",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorEQ(string) true"),
			Comparator: ononoki.ComparatorEQ,
			Value:      "a",
			Threshold:  "a",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorNE(string) false"),
			Comparator: ononoki.ComparatorNE,
			Value:      "a",
			Threshold:  "a",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorNE(string) true"),
			Comparator: ononoki.ComparatorNE,
			Value:      "a",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorLT(string) true"),
			Comparator: ononoki.ComparatorLT,
			Value:      "a",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorLT(string) false 1"),
			Comparator: ononoki.ComparatorLT,
			Value:      "b",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorLT(string) false 2"),
			Comparator: ononoki.ComparatorLT,
			Value:      "c",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorLE(string) true 1"),
			Comparator: ononoki.ComparatorLE,
			Value:      "a",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorLE(string) true 2"),
			Comparator: ononoki.ComparatorLE,
			Value:      "b",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorLE(string) false"),
			Comparator: ononoki.ComparatorLE,
			Value:      "c",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorGT(string) true"),
			Comparator: ononoki.ComparatorGT,
			Value:      "c",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorGT(string) false 1"),
			Comparator: ononoki.ComparatorGT,
			Value:      "b",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorGT(string) false 2"),
			Comparator: ononoki.ComparatorGT,
			Value:      "a",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorGE(string) true 1"),
			Comparator: ononoki.ComparatorGE,
			Value:      "b",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorGE(string) true 2"),
			Comparator: ononoki.ComparatorGE,
			Value:      "c",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorGE(string) false"),
			Comparator: ononoki.ComparatorGE,
			Value:      "a",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       fmt.Sprintf("ComparatorEQ(int) false"),
			Comparator: ononoki.ComparatorEQ,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorEQ(int) true"),
			Comparator: ononoki.ComparatorEQ,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorNE(int) false"),
			Comparator: ononoki.ComparatorNE,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorNE(int) true"),
			Comparator: ononoki.ComparatorNE,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorLT(int) true"),
			Comparator: ononoki.ComparatorLT,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorLT(int) false 1"),
			Comparator: ononoki.ComparatorLT,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorLT(int) false 2"),
			Comparator: ononoki.ComparatorLT,
			Value:      int64(2),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorLE(int) true 1"),
			Comparator: ononoki.ComparatorLE,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorLE(int) true 2"),
			Comparator: ononoki.ComparatorLE,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorLE(int) false"),
			Comparator: ononoki.ComparatorLE,
			Value:      int64(3),
			Threshold:  int64(2),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorGT(int) true"),
			Comparator: ononoki.ComparatorGT,
			Value:      int64(3),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorGT(int) false 1"),
			Comparator: ononoki.ComparatorGT,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorGT(int) false 2"),
			Comparator: ononoki.ComparatorGT,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorGE(int) true 1"),
			Comparator: ononoki.ComparatorGE,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorGE(int) true 2"),
			Comparator: ononoki.ComparatorGE,
			Value:      int64(3),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       fmt.Sprintf("ComparatorGE(int) false"),
			Comparator: ononoki.ComparatorGE,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   false,
			Type:       reflect.Int64,
		},

		// ToDo: add more test cases
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			actual := false
			switch tc.Type {
			case reflect.Int64:
				actual = ononoki.Compare(tc.Comparator, tc.Value.(int64), tc.Threshold.(int64))
			case reflect.Float64:
				actual = ononoki.Compare(tc.Comparator, tc.Value.(float64), tc.Threshold.(float64))
			case reflect.String:
				actual = ononoki.Compare[string](tc.Comparator, tc.Value.(string), tc.Threshold.(string))
			default:
				r.Fail(ononoki.ErrUnreachableCode.Error())
			}

			r.Equal(tc.Expected, actual)
		})
	}
}
