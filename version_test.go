package version_test

import (
	"fmt"
	"testing"

	"github.com/geoffreybauduin/version"
	"github.com/stretchr/testify/assert"
)

func TestVersion_Is(t *testing.T) {
	for _, testData := range []struct {
		CurrentVersion   string
		ConditionToCheck string
		IsResult         bool
	}{
		{"1.0.0", "=1.0.0", true},
		{"1.0.0", ">1.0.0", false},
		{"1.0.0", "<1.0.0", false},
		{"1.0.0", "<=1.0.0", true},
		{"1.0.0", ">=1.0.0", true},
		{"1.0.1", "1.0.1", true},
		{"1.0.0", "=1.0.1", false},
		{"1.0.0", "<1.0.1", true},
		{"1.0.1", ">1.0.0", true},
		// 2 exhausted
		{"1.0.0", ">1.0", false},
		{"1.0.0", ">=1.0", true},
		{"1.0.0", "=1.0", true},
		{"1.0.1", "=1.0", true}, // not specific so it should work
		{"1.0.0", "<1.0", false},
		{"1.0.0", "<=1.0", true},
		{"1.0.1", "<=1.0", true}, // same here, not specific
	} {
		t.Run(fmt.Sprintf("testing if %s %s returns %t", testData.CurrentVersion, testData.ConditionToCheck, testData.IsResult), func(t *testing.T) {
			v, err := version.New(testData.CurrentVersion)
			if assert.NoError(t, err) {
				assert.Equal(t, testData.IsResult, v.Is(testData.ConditionToCheck))
			}
		})
	}
}
