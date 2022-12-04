package ononoki_engine

type Rule struct {
	// ID is an optional unique identifier
	ID string `json:"id,omitempty"`

	// Name is an optional fact identifier
	Name string `json:"name,omitempty"`

	// Verifier is an optional verifier to verify whether the rule is valid or not
	Verifier Verifier

	// Conclusion is an optional conclusion
	// if Conclusion is not nil, means this is the leaf of the decision tree
	Conclusion *Conclusion

	// Children is sets of Rule to make the decision tree
	Children []Concluder
}

// RuleOptFunc is used for setup optional attributes of Rule
type RuleOptFunc func(*Rule)

// RuleWithID sets Rule.ID
func RuleWithID(id string) RuleOptFunc {
	return func(r *Rule) {
		r.ID = id
	}
}

// RuleWithName sets Rule.Name
func RuleWithName(name string) RuleOptFunc {
	return func(r *Rule) {
		r.Name = name
	}
}

// NewRuleLeaf creates *Rule on leaf which has only Conclusion not children
func NewRuleLeaf(verifier Verifier, conclusion *Conclusion, opts ...RuleOptFunc) *Rule {
	return NewRule(nil, verifier, conclusion, opts...)
}

// NewRuleRoot creates the root of the decision tree
func NewRuleRoot(children []Concluder, opts ...RuleOptFunc) *Rule {
	return NewRule(children, nil, nil, opts...)
}

// NewRule creates *Rule
func NewRule(children []Concluder, verifier Verifier, conclusion *Conclusion, opts ...RuleOptFunc) *Rule {
	r := &Rule{
		Verifier:   verifier,
		Conclusion: conclusion,
		Children:   children,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *Rule) Conclude(m map[string]any) (*Conclusion, error) {
	n := len(r.Children)
	if r.Verifier == nil && n == 0 {
		return nil, ErrRuleInvalid
	}

	if r.Verifier != nil {
		ok, err := r.Verifier.Verify(m)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}
	} else {
		return r.traverse(m)
	}

	if n == 0 && r.Conclusion != nil {
		return r.Conclusion, nil
	}

	return r.traverse(m)
}

func (r *Rule) traverse(m map[string]any) (*Conclusion, error) {
	for _, child := range r.Children {
		conclusion, err := child.Conclude(m)
		if err != nil {
			return nil, err
		}

		if conclusion != nil {
			return conclusion, nil
		}
	}

	return nil, ErrRuleNoConclusion
}
