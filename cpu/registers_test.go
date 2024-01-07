package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAF(t *testing.T) {
	var r = Registers{
		A: 128, // 10000000
		F: 170, // 10101010
	}

	var actual = r.GetAF()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetAF(t *testing.T) {
	var r = Registers{}
	r.SetAF(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.A) // 10000000
	assert.Equal(t, uint8(170), r.F) // 10101010
}

func TestGetBC(t *testing.T) {
	var r = Registers{
		B: 128, // 10000000
		C: 170, // 10101010
	}

	var actual = r.GetBC()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetBC(t *testing.T) {
	var r = Registers{}
	r.SetBC(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.B) // 10000000
	assert.Equal(t, uint8(170), r.C) // 10101010
}

func TestGetDE(t *testing.T) {
	var r = Registers{
		D: 128, // 10000000
		E: 170, // 10101010
	}

	var actual = r.GetDE()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetDE(t *testing.T) {
	var r = Registers{}
	r.SetDE(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.D) // 10000000
	assert.Equal(t, uint8(170), r.E) // 10101010
}

func TestGetHL(t *testing.T) {
	var r = Registers{
		H: 128, // 10000000
		L: 170, // 10101010
	}

	var actual = r.GetHL()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetHL(t *testing.T) {
	var r = Registers{}
	r.SetHL(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.H) // 10000000
	assert.Equal(t, uint8(170), r.L) // 10101010
}

func TestGetZeroFlag(t *testing.T) {
	var r = Registers{F: 128}

	var actual = r.GetZeroFlag()
	assert.True(t, actual)

	r = Registers{}
	actual = r.GetZeroFlag()
	assert.False(t, actual)
}

func TestSetZeroFlag(t *testing.T) {
	var r = Registers{F: 128}

	assert.True(t, r.GetZeroFlag())

	r.SetZeroFlag(false)
	assert.False(t, r.GetZeroFlag())

	r.SetZeroFlag(true)
	assert.True(t, r.GetZeroFlag())
}

func TestGetSubtractFlag(t *testing.T) {
	var r = Registers{F: 64}

	var actual = r.GetSubtractFlag()
	assert.True(t, actual)

	r = Registers{}
	actual = r.GetSubtractFlag()
	assert.False(t, actual)
}

func TestSetSubtractFlag(t *testing.T) {
	var r = Registers{F: 64}

	assert.True(t, r.GetSubtractFlag())

	r.SetSubtractFlag(false)
	assert.False(t, r.GetSubtractFlag())

	r.SetSubtractFlag(true)
	assert.True(t, r.GetSubtractFlag())
}

func TestGetHalfCarryFlag(t *testing.T) {
	var r = Registers{F: 32}

	var actual = r.GetHalfCarryFlag()
	assert.True(t, actual)

	r = Registers{}
	actual = r.GetHalfCarryFlag()
	assert.False(t, actual)
}

func TestSetHalfCarryFlag(t *testing.T) {
	var r = Registers{F: 32}

	assert.True(t, r.GetHalfCarryFlag())

	r.SetHalfCarryFlag(false)
	assert.False(t, r.GetHalfCarryFlag())

	r.SetHalfCarryFlag(true)
	assert.True(t, r.GetHalfCarryFlag())
}

func TestGetCarryFlag(t *testing.T) {
	var r = Registers{F: 16}

	var actual = r.GetCarryFlag()
	assert.True(t, actual)

	r = Registers{}
	actual = r.GetCarryFlag()
	assert.False(t, actual)
}

func TestSetCarryFlag(t *testing.T) {
	var r = Registers{F: 16}

	assert.True(t, r.GetCarryFlag())

	r.SetCarryFlag(false)
	assert.False(t, r.GetCarryFlag())

	r.SetCarryFlag(true)
	assert.True(t, r.GetCarryFlag())
}

func TestSetFlagsFunctionIndependently(t *testing.T) {
	var r = Registers{}

	assert.False(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetHalfCarryFlag())
	assert.False(t, r.GetCarryFlag())

	r.SetZeroFlag(true)
	assert.True(t, r.GetZeroFlag())
	assert.False(t, r.GetSubtractFlag())
	assert.False(t, r.GetHalfCarryFlag())
	assert.False(t, r.GetCarryFlag())

	r.SetSubtractFlag(true)
	assert.True(t, r.GetZeroFlag())
	assert.True(t, r.GetSubtractFlag())
	assert.False(t, r.GetHalfCarryFlag())
	assert.False(t, r.GetCarryFlag())

	r.SetHalfCarryFlag(true)
	assert.True(t, r.GetZeroFlag())
	assert.True(t, r.GetSubtractFlag())
	assert.True(t, r.GetHalfCarryFlag())
	assert.False(t, r.GetCarryFlag())

	r.SetCarryFlag(true)
	assert.True(t, r.GetZeroFlag())
	assert.True(t, r.GetSubtractFlag())
	assert.True(t, r.GetHalfCarryFlag())
	assert.True(t, r.GetCarryFlag())
}
