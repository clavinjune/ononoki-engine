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
	ConclusionName string = "new conclusion"
	ConclusionID   string = "C001"
)

func TestNewConclusion(t *testing.T) {
	r := require.New(t)
	c := ononoki.NewConclusion(ConclusionName)

	r.Equal("", c.ID)
	r.Equal(ConclusionName, c.Name)
}

func TestConclusionWithOpts(t *testing.T) {
	r := require.New(t)

	c := ononoki.NewConclusion(ConclusionName,
		ononoki.ConclusionWithID(ConclusionID),
	)

	r.Equal(ConclusionID, c.ID)
	r.Equal(ConclusionName, c.Name)
}
