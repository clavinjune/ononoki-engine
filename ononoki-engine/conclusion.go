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

// Conclusion holds the result of a certain ruleset
type Conclusion struct {
	_ struct{}
	// ID is an optional unique identifier
	ID string `json:"id,omitempty"`

	// Name is a conclusion identifier
	Name string `json:"name"`
}

// ConclusionOptFunc is used for setup optional attributes of Conclusion
type ConclusionOptFunc func(*Conclusion)

// ConclusionWithID sets Conclusion.ID
func ConclusionWithID(id string) ConclusionOptFunc {
	return func(c *Conclusion) {
		c.ID = id
	}
}

// NewConclusion creates *Conclusion
func NewConclusion(name string, opts ...ConclusionOptFunc) *Conclusion {
	c := &Conclusion{
		Name: name,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
