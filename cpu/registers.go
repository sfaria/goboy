package cpu

type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8
	h uint8
	l uint8
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
	return getVirtual(r.a, r.f)
}

func (r *Registers) SetAF(val uint16) {
	setVirtual(&r.a, &r.f, val)
}

func (r Registers) GetBC() uint16 {
	return getVirtual(r.b, r.c)
}

func (r *Registers) SetBC(val uint16) {
	setVirtual(&r.b, &r.c, val)
}

func (r Registers) GetDE() uint16 {
	return getVirtual(r.d, r.e)
}

func (r *Registers) SetDE(val uint16) {
	setVirtual(&r.d, &r.e, val)
}

func (r Registers) GetHL() uint16 {
	return getVirtual(r.h, r.l)
}

func (r *Registers) SetHL(val uint16) {
	setVirtual(&r.h, &r.l, val)
}

func getVirtual(r1 uint8, r2 uint8) uint16 {
	return (uint16(r1) << 8) | uint16(r2)
}

func setVirtual(r1 *uint8, r2 *uint8, val uint16) {
	*r1 = uint8((val & 0xFF00) >> 8)
	*r2 = uint8((val & 0x00FF))
}

func getFlag(r Registers, bitIndex uint8) bool {
	return ((r.f >> bitIndex) & 0x01) != 0
}

func setFlag(r *Registers, bitIndex uint8, val bool) {
	if val {
		r.f = r.f | (uint8(1) << bitIndex)
	} else {
		r.f = r.f & ^(uint8(1) << bitIndex)
	}
}
