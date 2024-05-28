package operations

import (
	"fmt"
	"goboy/cpu"
)

func LDnnn(r *cpu.Registers, target RegisterTarget, n uint8) (uint8, error) {
	if target == A {
		return 0, fmt.Errorf("register a cannot be the target of an LD instruction")
	}

	return target.SetValue(r, n)
}

func LDr1r2(r *cpu.Registers, r1 RegisterTarget, r2 RegisterTarget) (uint8, error) {
	return 0, fmt.Errorf("Unimplemented")
}
