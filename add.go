package asm

// AddRegisterNumber adds a number to the given register.
func (a *Assembler) AddRegisterNumber(registerNameTo string, number uint64) {
	if registerNameTo == "al" {
		a.WriteBytes(0x04, byte(number))
		return
	}

	a.numberToRegister(0x83, 0x80, 0, false, true, false, false, registerNameTo, number)
}

// AddRegisterRegister adds a register value into another register.
func (a *Assembler) AddRegisterRegister(registerNameTo string, registerNameFrom string) {
	a.registerToRegister([]byte{0x01}, []byte{0x00}, registerNameTo, registerNameFrom, false)
}
