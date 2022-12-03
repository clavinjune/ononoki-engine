package ononoki_engine

import (
	"reflect"
)

// Fact holds the core of a certain ruleset
type Fact[T Threshold] struct {
	// ID is an optional unique identifier
	ID string `json:"id,omitempty"`

	// Name is an optional fact identifier
	Name string `json:"name,omitempty"`

	// Property is where the fact comes from
	Property string `json:"property"`

	// Comparator compares Fact.Property value with Fact.Threshold
	Comparator Comparator `json:"comparator"`

	// Threshold is the value to be compared with Fact.Property
	Threshold T `json:"threshold"`
}

// FactOptFunc is used for setup optional attributes of Fact
type FactOptFunc[T Threshold] func(fact *Fact[T])

// FactWithID sets Fact.ID
func FactWithID[T Threshold](id string) FactOptFunc[T] {
	return func(f *Fact[T]) {
		f.ID = id
	}
}

// FactWithName sets Fact.Name
func FactWithName[T Threshold](name string) FactOptFunc[T] {
	return func(f *Fact[T]) {
		f.Name = name
	}
}

// NewFact creates *Fact
func NewFact[T Threshold](property string, comparator Comparator, threshold T, opts ...FactOptFunc[T]) *Fact[T] {
	f := &Fact[T]{
		Property:   property,
		Comparator: comparator,
		Threshold:  threshold,
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f *Fact[T]) Verify(m map[string]any) (bool, error) {
	val, kind, err := f.propertyMetadata(m)
	if err != nil {
		return false, err
	}

	if kind == reflect.Invalid {
		return false, ErrFactPropertyTypeMismatched
	}

	switch t := any(f.Threshold).(type) {
	case int64:
		if kind != reflect.Int64 {
			return false, ErrFactPropertyTypeMismatched
		}

		return Compare(f.Comparator, val.(int64), t), nil
	case float64:
		if kind != reflect.Float64 {
			return false, ErrFactPropertyTypeMismatched
		}

		return Compare(f.Comparator, val.(float64), t), nil
	case string:
		if kind != reflect.String {
			return false, ErrFactPropertyTypeMismatched
		}
		return Compare(f.Comparator, val.(string), t), nil
	}

	return false, ErrUnreachableCode
}

func (f *Fact[T]) propertyMetadata(m map[string]any) (any, reflect.Kind, error) {
	if m == nil {
		return nil, reflect.Invalid, ErrFactDataEmpty
	}

	val, ok := m[f.Property]
	if !ok {
		return nil, reflect.Invalid, ErrFactPropertyNotFound
	}

	tv := reflect.TypeOf(val)
	if tv == nil {
		return nil, reflect.Invalid, ErrFactPropertyValueNil
	}

	if tv.Kind() == reflect.Pointer || tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	return val, tv.Kind(), nil
}
