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
	return ((r.f >> 7) & 0x0F) != 0
}

func (r *Registers) SetZeroFlag(value bool) {
	if value {
		r.f = r.f | (uint8(1) << 7)
	} else {
		r.f = r.f & ^(uint8(1) << 7)
	}
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
