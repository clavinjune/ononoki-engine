package ononoki_engine_test

import (
	ononoki "github.com/clavinjune/ononoki-engine"
	"github.com/stretchr/testify/require"
	"testing"
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
