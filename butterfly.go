package main

import (
	"fmt"

	"github.com/drewwalton19216801/butterfly/cpu"
)

func main() {
	cpu := cpu.NewBasic6502()
	cpu.Reset()
	// Print the CPU registers in hex
	fmt.Printf("A: %02X X: %02X Y: %02X SP: %02X PC: %04X P: %02X\n", cpu.Registers.A, cpu.Registers.X, cpu.Registers.Y, cpu.Registers.SP, cpu.Registers.PC, cpu.Registers.P)
}
