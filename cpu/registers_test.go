package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAF(t *testing.T) {
	var r = Registers{
		a: 128, // 10000000
		f: 170, // 10101010
	}

	var actual = r.GetAF()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetAF(t *testing.T) {
	var r = Registers{}
	r.SetAF(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.a) // 10000000
	assert.Equal(t, uint8(170), r.f) // 10101010
}

func TestGetBC(t *testing.T) {
	var r = Registers{
		b: 128, // 10000000
		c: 170, // 10101010
	}

	var actual = r.GetBC()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetBC(t *testing.T) {
	var r = Registers{}
	r.SetBC(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.b) // 10000000
	assert.Equal(t, uint8(170), r.c) // 10101010
}

func TestGetDE(t *testing.T) {
	var r = Registers{
		d: 128, // 10000000
		e: 170, // 10101010
	}

	var actual = r.GetDE()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetDE(t *testing.T) {
	var r = Registers{}
	r.SetDE(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.d) // 10000000
	assert.Equal(t, uint8(170), r.e) // 10101010
}

func TestGetHL(t *testing.T) {
	var r = Registers{
		h: 128, // 10000000
		l: 170, // 10101010
	}

	var actual = r.GetHL()
	var expected uint16 = 32938 // 1000000010101010
	assert.Equal(t, expected, actual)
}

func TestSetHL(t *testing.T) {
	var r = Registers{}
	r.SetHL(32938) // 1000000010101010

	assert.Equal(t, uint8(128), r.h) // 10000000
	assert.Equal(t, uint8(170), r.l) // 10101010
}
