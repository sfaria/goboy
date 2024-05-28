package operations

import (
	"fmt"
	"goboy/cpu"
)

func Add(r *cpu.Registers, target RegisterTarget) (uint8, error) {
	if target == A {
		return 0, fmt.Errorf("register a cannot be the target of an ADD instruction")
	}

	var val, err = target.GetValue(*r)
	if err != nil {
		return 0, err
	}

	var resultWithOverflow uint16 = (uint16(r.A) & 0x00FF) + (uint16(val) & 0x00FF)
	var didCarry = resultWithOverflow > 255
	var result uint8 = uint8(resultWithOverflow & 0x00FF)
	var didHalfCarry = ((r.A & 0x0F) + (val & 0x0F)) > 0x0F

	r.A = result
	r.SetZeroFlag(result == 0)
	r.SetSubtractFlag(false)
	r.SetCarryFlag(didCarry)
	r.SetHalfCarryFlag(didHalfCarry)

	return r.A, nil
}
