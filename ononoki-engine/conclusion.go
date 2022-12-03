package ononoki_engine

// Conclusion holds the result of a certain ruleset
type Conclusion struct {
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
