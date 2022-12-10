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

package ononoki_test

import (
	"reflect"
	"testing"

	"github.com/clavinjune/ononoki"
	"github.com/stretchr/testify/require"
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
			Name:       "ComparatorEQ(string) false",
			Comparator: ononoki.ComparatorEQ,
			Value:      "a",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorEQ(string) true",
			Comparator: ononoki.ComparatorEQ,
			Value:      "a",
			Threshold:  "a",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorNE(string) false",
			Comparator: ononoki.ComparatorNE,
			Value:      "a",
			Threshold:  "a",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorNE(string) true",
			Comparator: ononoki.ComparatorNE,
			Value:      "a",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorLT(string) true",
			Comparator: ononoki.ComparatorLT,
			Value:      "a",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorLT(string) false 1",
			Comparator: ononoki.ComparatorLT,
			Value:      "b",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorLT(string) false 2",
			Comparator: ononoki.ComparatorLT,
			Value:      "c",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorLE(string) true 1",
			Comparator: ononoki.ComparatorLE,
			Value:      "a",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorLE(string) true 2",
			Comparator: ononoki.ComparatorLE,
			Value:      "b",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorLE(string) false",
			Comparator: ononoki.ComparatorLE,
			Value:      "c",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorGT(string) true",
			Comparator: ononoki.ComparatorGT,
			Value:      "c",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorGT(string) false 1",
			Comparator: ononoki.ComparatorGT,
			Value:      "b",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorGT(string) false 2",
			Comparator: ononoki.ComparatorGT,
			Value:      "a",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorGE(string) true 1",
			Comparator: ononoki.ComparatorGE,
			Value:      "b",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorGE(string) true 2",
			Comparator: ononoki.ComparatorGE,
			Value:      "c",
			Threshold:  "b",
			Expected:   true,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorGE(string) false",
			Comparator: ononoki.ComparatorGE,
			Value:      "a",
			Threshold:  "b",
			Expected:   false,
			Type:       reflect.String,
		},
		{
			Name:       "ComparatorEQ(int) false",
			Comparator: ononoki.ComparatorEQ,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorEQ(int) true",
			Comparator: ononoki.ComparatorEQ,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorNE(int) false",
			Comparator: ononoki.ComparatorNE,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorNE(int) true",
			Comparator: ononoki.ComparatorNE,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorLT(int) true",
			Comparator: ononoki.ComparatorLT,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorLT(int) false 1",
			Comparator: ononoki.ComparatorLT,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorLT(int) false 2",
			Comparator: ononoki.ComparatorLT,
			Value:      int64(2),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorLE(int) true 1",
			Comparator: ononoki.ComparatorLE,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorLE(int) true 2",
			Comparator: ononoki.ComparatorLE,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorLE(int) false",
			Comparator: ononoki.ComparatorLE,
			Value:      int64(3),
			Threshold:  int64(2),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorGT(int) true",
			Comparator: ononoki.ComparatorGT,
			Value:      int64(3),
			Threshold:  int64(2),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorGT(int) false 1",
			Comparator: ononoki.ComparatorGT,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorGT(int) false 2",
			Comparator: ononoki.ComparatorGT,
			Value:      int64(1),
			Threshold:  int64(2),
			Expected:   false,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorGE(int) true 1",
			Comparator: ononoki.ComparatorGE,
			Value:      int64(1),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorGE(int) true 2",
			Comparator: ononoki.ComparatorGE,
			Value:      int64(3),
			Threshold:  int64(1),
			Expected:   true,
			Type:       reflect.Int64,
		},
		{
			Name:       "ComparatorGE(int) false",
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
