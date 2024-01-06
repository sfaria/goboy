package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWithoutOverflow(t *testing.T) {
	var r = Registers{a: 0, b: 32}
	var val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(32), val)
	assert.Equal(t, uint8(32), r.a)
	assert.False(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetCarryFlag())
	assert.False(t, r.GetHalfCarryFlag())

	// zero flag
	r = Registers{a: 0, b: 0}
	val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(0), val)
	assert.Equal(t, uint8(0), r.a)
	assert.True(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetCarryFlag())
	assert.False(t, r.GetHalfCarryFlag())
}

func TestAddWithOverflow(t *testing.T) {
	var r = Registers{a: 255, b: 1}
	var val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(0), val)
	assert.Equal(t, uint8(0), r.a)
	assert.True(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.True(t, r.GetCarryFlag())
	assert.True(t, r.GetHalfCarryFlag())
}

func TestAddWithHalfCarry(t *testing.T) {
	var r = Registers{a: 143, b: 1}
	var val, err = Add(&r, B)

	assert.NoError(t, err)
	assert.Equal(t, uint8(144), val)
	assert.Equal(t, uint8(144), r.a)
	assert.False(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetCarryFlag())
	assert.True(t, r.GetHalfCarryFlag())
}

func TestAddWithATargetReturnsError(t *testing.T) {
	var r = Registers{}
	var val, err = Add(&r, A)

	assert.Equal(t, uint8(0), val)
	assert.True(t, err != nil)
}
