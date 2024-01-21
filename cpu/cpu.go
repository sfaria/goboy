package cpu

type CPU struct {
	registers Registers
	pc        uint16
	bus       MemoryBus
}
