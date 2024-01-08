package operations

import (
	"goboy/cpu"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDnnn(t *testing.T) {
	var r = cpu.Registers{}

	for target := B; target < L; target = target + 1 {
		var n uint8 = uint8(rand.Intn(256))
		var val, ldErr = LDnnn(&r, target, n)
		assert.Equal(t, n, val)
		assert.Nil(t, ldErr)

		var actual, err = target.GetValue(r)
		assert.Nil(t, err)
		assert.Equal(t, n, actual)
	}
}

func TestLDnnnDoesNotAllowAccumulator(t *testing.T) {
	var r = cpu.Registers{}
	var val, err = LDnnn(&r, A, uint8(rand.Intn(256)))
	assert.Equal(t, val, uint8(0))
	assert.NotNil(t, err)
}
