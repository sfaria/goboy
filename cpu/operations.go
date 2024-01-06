package cpu

import (
	"fmt"
)

type RegisterTarget int

const (
	A RegisterTarget = iota
	B
	C
	D
	E
	H
	L
)

func Add(r *Registers, target RegisterTarget) (uint8, error) {
	if target == A {
		return 0, fmt.Errorf("register a cannot be the target of an add instruction")
	}

	var val, err = getValue(r, target)
	if err != nil {
		return 0, err
	}

	var resultWithOverflow uint16 = (uint16(r.a) & 0x00FF) + (uint16(val) & 0x00FF)
	var didCarry = resultWithOverflow > 255
	var result uint8 = uint8(resultWithOverflow & 0x00FF)
	var didHalfCarry = ((r.a & 0x0F) + (val & 0x0F)) > 0x0F

	r.a = result
	r.SetZeroFlag(result == 0)
	r.SetSubtractFlag(false)
	r.SetCarryFlag(didCarry)
	r.SetHalfCarryFlag(didHalfCarry)

	return r.a, nil
}

func getValue(r *Registers, target RegisterTarget) (uint8, error) {
	switch target {
	case A:
		return r.a, nil
	case B:
		return r.b, nil
	case C:
		return r.c, nil
	case D:
		return r.d, nil
	case E:
		return r.e, nil
	case H:
		return r.h, nil
	case L:
		return r.l, nil
	}

	return 0, fmt.Errorf("unknown register target: %d", target)
}
