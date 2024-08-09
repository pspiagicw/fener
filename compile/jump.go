package compile

import "github.com/pspiagicw/fener/code"

func (c *Compiler) JumpOptimizer() {
	for start, end := range c.jumpTable {
		bytes := code.ToBytes(c.instructions[:end])
		c.instructions[start].Operands = []int{len(bytes)}
	}
}
