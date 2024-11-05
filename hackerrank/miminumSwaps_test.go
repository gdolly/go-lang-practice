package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	actual := MinimumSwaps([]int32{2, 3, 4, 1, 5})
	expected := int32(3)
	assert.Equal(t, expected, actual)
}
