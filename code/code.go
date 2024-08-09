package code

import "encoding/binary"

//go:generate stringer -type=OpCode

type OpCode byte

const (
	PUSH OpCode = iota
	ADD
	SUB
	DIV
	MUL

	TRUE
	FALSE
	AND
	OR
	GT
	LT
	EQ
	NOT
	NEQ

	JCMP
	JMP
	JT

	SET
	GET
)

type Instruction struct {
	Opcode   OpCode
	Operands []int
}

type Definition struct {
	Opcode      OpCode
	NumOperands int
}

var definitions = map[OpCode]Definition{
	PUSH: {PUSH, 1},

	ADD: {ADD, 0},
	SUB: {SUB, 0},
	DIV: {DIV, 0},
	MUL: {MUL, 0},

	TRUE:  {TRUE, 0},
	FALSE: {FALSE, 0},
	AND:   {AND, 0},
	OR:    {OR, 0},
	GT:    {GT, 0},
	LT:    {LT, 0},
	EQ:    {EQ, 0},
	NOT:   {NOT, 0},
	NEQ:   {NEQ, 0},

	SET: {SET, 1},
	GET: {GET, 1},

	JCMP: {JCMP, 1},
	JMP:  {JMP, 1},
}

func Make(op OpCode, operands ...int) *Instruction {
	def, ok := definitions[op]

	if !ok {
		return nil
	}

	if def.NumOperands != len(operands) {
		return nil
	}

	instr := &Instruction{Opcode: def.Opcode, Operands: operands}

	return instr
}

func ToBytes(instructions []*Instruction) []byte {
	bytes := []byte{}

	for _, i := range instructions {
		bytes = append(bytes, byte(i.Opcode))
		for _, operand := range i.Operands {
			bytes = binary.BigEndian.AppendUint16(bytes, uint16(operand))
		}
	}

	return bytes
}
