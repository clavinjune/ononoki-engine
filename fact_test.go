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
	"testing"

	"github.com/clavinjune/ononoki"
	"github.com/stretchr/testify/require"
)

const (
	FactID         string             = "F001"
	FactName       string             = "new fact"
	FactProperty   string             = "example"
	FactComparator ononoki.Comparator = ononoki.ComparatorGT
	FactThreshold  int64              = 20
)

func TestNewFact(t *testing.T) {
	r := require.New(t)

	f := ononoki.NewFact(FactProperty, FactComparator, FactThreshold)

	r.Equal("", f.ID)
	r.Equal("", f.Name)
	r.Equal(FactProperty, f.Property)
	r.Equal(FactComparator, f.Comparator)
	r.Equal(FactThreshold, f.Threshold)
}

func TestFactWithOpts(t *testing.T) {
	r := require.New(t)

	f := ononoki.NewFact(FactProperty, FactComparator, FactThreshold,
		ononoki.FactWithID[int64](FactID),
		ononoki.FactWithName[int64](FactName),
	)

	r.Equal(FactID, f.ID)
	r.Equal(FactName, f.Name)
	r.Equal(FactProperty, f.Property)
	r.Equal(FactComparator, f.Comparator)
	r.Equal(FactThreshold, f.Threshold)
}

func TestFact_Verify(t *testing.T) {
	tt := []struct {
		Name         string
		Map          map[string]any
		ExpectedBool bool
		ExpectedErr  error
	}{
		{
			Name:         "map is nil",
			Map:          nil,
			ExpectedBool: false,
			ExpectedErr:  ononoki.ErrFactDataEmpty,
		},
		{
			Name:         "map no property",
			Map:          map[string]any{},
			ExpectedBool: false,
			ExpectedErr:  ononoki.ErrFactPropertyNotFound,
		},
		{
			Name: "map property is nil",
			Map: map[string]any{
				FactProperty: nil,
			},
			ExpectedBool: false,
			ExpectedErr:  ononoki.ErrFactPropertyValueNil,
		},
		{
			Name: "map property type is mismatched",
			Map: map[string]any{
				FactProperty: "random",
			},
			ExpectedBool: false,
			ExpectedErr:  ononoki.ErrFactPropertyTypeMismatched,
		},
		{
			Name: "map property less than threshold",
			Map: map[string]any{
				FactProperty: FactThreshold - 1,
			},
			ExpectedBool: false,
			ExpectedErr:  nil,
		},
		{
			Name: "map property verified",
			Map: map[string]any{
				FactProperty: FactThreshold + 1,
			},
			ExpectedBool: true,
			ExpectedErr:  nil,
		},
	}

	f := ononoki.NewFact(FactProperty, FactComparator, FactThreshold)

	for i := range tt {
		tc := tt[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			ok, err := f.Verify(tc.Map)
			r.Equal(tc.ExpectedBool, ok)
			r.Equal(tc.ExpectedErr, err)
		})
	}
}
