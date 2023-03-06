package cpu

// Basic6502 is a basic 6502 CPU

type Basic6502 struct {
	// Registers
	Registers Registers
}

type Registers struct {
	A  uint8  // Accumulator
	X  uint8  // X register
	Y  uint8  // Y register
	SP uint8  // Stack pointer
	PC uint16 // Program counter
	P  uint8  // Processor status
}

// Processor status flag bits
const (
	CarryFlag      = 0x01 // Carry flag
	ZeroFlag       = 0x02 // Zero flag
	InterruptFlag  = 0x04 // Interrupt disable flag
	DecimalFlag    = 0x08 // Decimal mode flag
	BreakFlag      = 0x10 // Break command flag
	OverflowFlag   = 0x40 // Overflow flag
	NegativeFlag   = 0x80 // Negative flag
	UnusedFlag     = 0x20 // Unused flag
	AllFlags       = 0xFF // All flags
	AllFlagsExcept = 0x1F // All flags except unused flag
)

// SetFlag sets a flag in the processor status register
func (cpu *Basic6502) SetFlag(flag uint8, value bool) {
	if value {
		cpu.Registers.P |= flag
	} else {
		cpu.Registers.P &^= flag
	}
}

// GetFlag gets a flag from the processor status register
func (cpu *Basic6502) GetFlag(flag uint8) bool {
	return cpu.Registers.P&flag != 0
}

// NewBasic6502 creates a new Basic6502 CPU
func NewBasic6502() *Basic6502 {
	return &Basic6502{}
}

// Reset resets the CPU
func (cpu *Basic6502) Reset() {
	// Reset the A, X, and Y registers to 0
	cpu.Registers.A = 0
	cpu.Registers.X = 0
	cpu.Registers.Y = 0
	// Reset the stack pointer to 0xFD
	cpu.Registers.SP = 0xFD
	// Reset the program counter to 0x0000
	cpu.Registers.PC = 0x0000 // TODO: Read the reset vector from memory
	// Reset the processor status to 0x24
	cpu.Registers.P = 0x24
}
