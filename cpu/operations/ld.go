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
