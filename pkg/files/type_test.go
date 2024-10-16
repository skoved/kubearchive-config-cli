// Copyright skoved

package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsYaml(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		expected bool
	}{
		{
			name:     "txt file is not a yaml file",
			file:     "testdata/text.txt",
			expected: false,
		},
		{
			name:     "detect a yaml file",
			file:     "testdata/cm.yaml",
			expected: true,
		},
		{
			name:     "invalid yaml files not detected",
			file:     "testdata/invalid.yaml",
			expected: false,
		},
	}
	for _, tc := range testCases {
		res := IsYaml(tc.file)
		assert.Equal(t, tc.expected, res)
	}
}
