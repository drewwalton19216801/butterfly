package cpu

// Basic6502 is a basic 6502 CPU

type Basic6502 struct {
	// Registers
	A  uint8  // Accumulator
	X  uint8  // X register
	Y  uint8  // Y register
	SP uint8  // Stack pointer
	PC uint16 // Program counter
	P  uint8  // Processor status
}

// NewBasic6502 creates a new Basic6502 CPU
func NewBasic6502() *Basic6502 {
	return &Basic6502{}
}

// Reset resets the CPU
func (cpu *Basic6502) Reset() {
	// Reset the A, X, and Y registers to 0
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	// Reset the stack pointer to 0xFD
	cpu.SP = 0xFD
	// Reset the program counter to 0x0000
	cpu.PC = 0x0000 // TODO: Read the reset vector from memory
	// Reset the processor status to 0x24
	cpu.P = 0x24
}
