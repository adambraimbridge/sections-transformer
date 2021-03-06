package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransform(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name    string
		term    term
		section section
	}{
		{"Transform term to section", term{
			CanonicalName: "Africa Section",
			RawID:         "Nstein_GL_AFTM_GL_164835"},
			section{
				UUID:      "adb4f804-c3b6-3eca-8708-5edeec653a27",
				PrefLabel: "Africa Section",
				AlternativeIdentifiers: alternativeIdentifiers{
					TME:   []string{"TnN0ZWluX0dMX0FGVE1fR0xfMTY0ODM1-U2VjdGlvbnM="},
					Uuids: []string{"adb4f804-c3b6-3eca-8708-5edeec653a27"},
				},
				Type: "Section"}},
	}

	for _, test := range tests {
		expectedSection := transformSection(test.term, "Sections")
		assert.Equal(test.section, expectedSection, fmt.Sprintf("%s: Expected section incorrect", test.name))
	}

}
