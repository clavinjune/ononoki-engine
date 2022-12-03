package ononoki_engine

import (
	"errors"
	"fmt"
)

var (
	ErrBase                       = errors.New("ononoki")
	ErrUnreachableCode            = fmt.Errorf("%w: unreachable code reached", ErrBase)
	ErrFact                       = fmt.Errorf("%w: fact", ErrBase)
	ErrFactDataEmpty              = fmt.Errorf("%w: empty data", ErrFact)
	ErrFactPropertyNotFound       = fmt.Errorf("%w: property not found", ErrFact)
	ErrFactPropertyValueNil       = fmt.Errorf("%w: property value nil", ErrFact)
	ErrFactPropertyTypeMismatched = fmt.Errorf("%w: property type mismatched", ErrFact)

	ErrRule             = fmt.Errorf("%w: rule", ErrBase)
	ErrRuleInvalid      = fmt.Errorf("%w: invalid rule", ErrRule)
	ErrRuleNoConclusion = fmt.Errorf("%w: no conclusion", ErrRule)
)
