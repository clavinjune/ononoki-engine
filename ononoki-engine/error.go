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

import (
	"errors"
	"fmt"
)

var (
	// ErrBase bases all ononoki-engine error
	ErrBase = errors.New("ononoki-engine")

	// ErrUnreachableCode means there's a bug on ononoki-engine' logic since client shouldn't meet this error
	ErrUnreachableCode = fmt.Errorf("%w: unreachable code reached", ErrBase)

	// ErrFact bases all Fact's error
	ErrFact = fmt.Errorf("%w: fact", ErrBase)

	// ErrFactDataEmpty means the data hasn't initialized
	ErrFactDataEmpty = fmt.Errorf("%w: empty data", ErrFact)

	// ErrFactPropertyNotFound means the fact can't found the given property on the data
	ErrFactPropertyNotFound = fmt.Errorf("%w: property not found", ErrFact)

	// ErrFactPropertyValueNil means the property value is empty, it's different with ErrFactPropertyNotFound
	ErrFactPropertyValueNil = fmt.Errorf("%w: property value nil", ErrFact)

	// ErrFactPropertyTypeMismatched means you have wrong type on the property data
	ErrFactPropertyTypeMismatched = fmt.Errorf("%w: property type mismatched", ErrFact)

	// ErrRule bases all Rule's error
	ErrRule = fmt.Errorf("%w: rule", ErrBase)
	// ErrRuleInvalid means the rule attributes don't seem to be correct
	ErrRuleInvalid = fmt.Errorf("%w: invalid rule", ErrRule)

	// ErrRuleNoConclusion means the rule has been traversed until the leaves but there's no verified fact there
	ErrRuleNoConclusion = fmt.Errorf("%w: no conclusion", ErrRule)
)
