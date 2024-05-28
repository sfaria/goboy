package cpu

type Registers struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8
	F uint8
	H uint8
	L uint8
}

func (r Registers) GetZeroFlag() bool {
	return getFlag(r, 7)
}

func (r *Registers) SetZeroFlag(val bool) {
	setFlag(r, 7, val)
}

func (r Registers) GetSubtractFlag() bool {
	return getFlag(r, 6)
}

func (r *Registers) SetSubtractFlag(val bool) {
	setFlag(r, 6, val)
}

func (r Registers) GetHalfCarryFlag() bool {
	return getFlag(r, 5)
}

func (r *Registers) SetHalfCarryFlag(val bool) {
	setFlag(r, 5, val)
}

func (r Registers) GetCarryFlag() bool {
	return getFlag(r, 4)
}

func (r *Registers) SetCarryFlag(val bool) {
	setFlag(r, 4, val)
}

func (r Registers) GetAF() uint16 {
	return getVirtual(r.A, r.F)
}

func (r *Registers) SetAF(val uint16) {
	setVirtual(&r.A, &r.F, val)
}

func (r Registers) GetBC() uint16 {
	return getVirtual(r.B, r.C)
}

func (r *Registers) SetBC(val uint16) {
	setVirtual(&r.B, &r.C, val)
}

func (r Registers) GetDE() uint16 {
	return getVirtual(r.D, r.E)
}

func (r *Registers) SetDE(val uint16) {
	setVirtual(&r.D, &r.E, val)
}

func (r Registers) GetHL() uint16 {
	return getVirtual(r.H, r.L)
}

func (r *Registers) SetHL(val uint16) {
	setVirtual(&r.H, &r.L, val)
}

func getVirtual(r1 uint8, r2 uint8) uint16 {
	return (uint16(r1) << 8) | uint16(r2)
}

func setVirtual(r1 *uint8, r2 *uint8, val uint16) {
	*r1 = uint8((val & 0xFF00) >> 8)
	*r2 = uint8((val & 0x00FF))
}

func getFlag(r Registers, bitIndex uint8) bool {
	return ((r.F >> bitIndex) & 0x01) != 0
}

func setFlag(r *Registers, bitIndex uint8, val bool) {
	if val {
		r.F = r.F | (uint8(1) << bitIndex)
	} else {
		r.F = r.F & ^(uint8(1) << bitIndex)
	}
}
