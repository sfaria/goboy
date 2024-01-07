package operations

import (
	"fmt"
	"goboy/cpu"
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

func (target RegisterTarget) getValue(r cpu.Registers) (uint8, error) {
	switch target {
	case A:
		return r.A, nil
	case B:
		return r.B, nil
	case C:
		return r.C, nil
	case D:
		return r.D, nil
	case E:
		return r.E, nil
	case H:
		return r.H, nil
	case L:
		return r.L, nil
	}

	return 0, fmt.Errorf("unknown register target: %d", target)
}
