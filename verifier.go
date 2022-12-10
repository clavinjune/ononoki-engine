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

// Verifier verifies whether the given data is valid or not
type Verifier interface {
	// Verify verifies whether the given data is valid or not
	Verify(map[string]any) (bool, error)
}

// VerifierFunc an adapter to Verifier
type VerifierFunc func(map[string]any) (bool, error)

// Verify verifies whether the given data is valid or not
func (v VerifierFunc) Verify(m map[string]any) (bool, error) {
	return v(m)
}
