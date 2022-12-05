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

// Concluder returns Conclusion from the given data
type Concluder interface {
	// Conclude returns Conclusion from the given data
	Conclude(map[string]any) (*Conclusion, error)
}

// ConcluderFunc an adapter to Concluder
type ConcluderFunc func(map[string]any) (*Conclusion, error)

// Conclude returns Conclusion from the given data
func (c ConcluderFunc) Conclude(m map[string]any) (*Conclusion, error) {
	return c(m)
}
