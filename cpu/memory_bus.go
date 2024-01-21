package cpu

type MemoryBus struct {
	bus [0xFFFF]uint8
}

func (m MemoryBus) ReadByte(address uint16) uint8 {
	return m.bus[address]
}
