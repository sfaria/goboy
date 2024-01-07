package operations

import (
	"goboy/cpu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWithoutOverflow(t *testing.T) {
	var r = cpu.Registers{A: 0, B: 32}
	var val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(32), val)
	assert.Equal(t, uint8(32), r.A)
	assert.False(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetCarryFlag())
	assert.False(t, r.GetHalfCarryFlag())

	// zero flag
	r = cpu.Registers{A: 0, B: 0}
	val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(0), val)
	assert.Equal(t, uint8(0), r.A)
	assert.True(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetCarryFlag())
	assert.False(t, r.GetHalfCarryFlag())
}

func TestAddWithOverflow(t *testing.T) {
	var r = cpu.Registers{A: 255, B: 1}
	var val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(0), val)
	assert.Equal(t, uint8(0), r.A)
	assert.True(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.True(t, r.GetCarryFlag())
	assert.True(t, r.GetHalfCarryFlag())
}

func TestAddWithHalfCarry(t *testing.T) {
	var r = cpu.Registers{A: 143, B: 1}
	var val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(144), val)
	assert.Equal(t, uint8(144), r.A)
	assert.False(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetCarryFlag())
	assert.True(t, r.GetHalfCarryFlag())
}

func TestAddWithATargetReturnsError(t *testing.T) {
	var r = cpu.Registers{}
	var val, err = Add(&r, A)

	assert.Equal(t, uint8(0), val)
	assert.True(t, err != nil)
}
