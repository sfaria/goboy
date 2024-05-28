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

func (target RegisterTarget) GetValue(r cpu.Registers) (uint8, error) {
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

func (target RegisterTarget) SetValue(r *cpu.Registers, val uint8) (uint8, error) {
	switch target {
	case A:
		r.A = val
		return r.A, nil
	case B:
		r.B = val
		return r.B, nil
	case C:
		r.C = val
		return r.C, nil
	case D:
		r.D = val
		return r.D, nil
	case E:
		r.E = val
		return r.E, nil
	case H:
		r.H = val
		return r.H, nil
	case L:
		r.L = val
		return r.L, nil
	}

	return 0, fmt.Errorf("unknown register target: %d", target)
}
