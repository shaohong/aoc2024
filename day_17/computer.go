//https://adventofcode.com/2024/day/17

/*
its program is a list of 3-bit numbers (0 through 7), like 0,1,2,3. The computer also has three registers named A, B, and C,

8 instructions/opcode

Each instruction also reads the 3-bit number after it as an input, a.k.a operand

instruction pointer -> points to next opcode

instruction pointer starts at 0, pointing at the first 3-bit number in the program

the instruction pointer increases by 2 after each instruction is processed (to move past the instruction's opcode and its operand)


Two types of operand:
  * literal
  * combo

Combo operands 0 through 3 represent literal values 0 through 3.
Combo operand 4 represents the value of register A.
Combo operand 5 represents the value of register B.
Combo operand 6 represents the value of register C.
Combo operand 7 is reserved and will not appear in valid programs.


eight instructions:
* opcode 0: adv (division)  int( register_A / (2 ** combo_operand)  ) -> regiter_A

* opcode 1: bxl (bitwise xor) register_B xor literal_operator -> register_B

* opcode 2: bst : (combo_operand % 8) -> register_B

* opcode 3: jnz:  if register_A is not zero, instruction_pointer += literal_operand (and ip not increased by 2 after the jump)

* opcode 4: bxc: (register_B xor register_C)-> register_B

* opcode 5: out: print (combo_operand % 8). (If a program outputs multiple values, they are separated by commas.)

* opcode 6: bdv:  int( register_A / (2 ** combo_operand)  ) -> regiter_B

* opcode 7: cdv: int( register_A / (2 ** combo_operand)  ) -> regiter_C

*/

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ParseInput(data string) (registers map[rune]int, program []int) {
	// Split the input into lines

	registers = make(map[rune]int)
	program = make([]int, 0)

	parts := strings.Split(data, "\n\n")

	part1_lines := strings.Split(parts[0], "\n")

	for i := 0; i < len(part1_lines); i++ {
		var tmp int
		var reg_name rune
		fmt.Sscanf(part1_lines[i], "Register %c: %d", &reg_name, &tmp)
		registers[reg_name] = tmp
	}

	var program_str string
	fmt.Sscanf(parts[1], "Program: %s", &program_str)
	for _, val := range strings.Split(program_str, ",") {
		num, _ := strconv.Atoi(val)
		program = append(program, num)
	}

	return registers, program
}

type Computer struct {
	registers           map[rune]int
	program             []int
	instruction_pointer int
	outputs             []int
}

type Instruction struct {
	opcode  int
	operand int
}

func (c *Computer) LiteralOperand(operand int) int {
	return operand
}

func (c *Computer) ComboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.registers['A']
	case 5:
		return c.registers['B']
	case 6:
		return c.registers['C']
	default:
		panic("Invalid operand")
	}
}

func (c *Computer) ExecuteInstruction(instruction Instruction) {
	fmt.Println("Executing instruction: ", instruction.String())
	// run the current instruction, update computer state after the execution
	switch instruction.opcode {
	case 0: // adv
		numerator := c.registers['A']
		denominator := 1
		for i := 0; i < c.ComboOperand(instruction.operand); i++ {
			denominator *= 2
		}
		c.registers['A'] = numerator / denominator
		c.instruction_pointer += 2

	case 1: // bxl
		c.registers['B'] = c.registers['B'] ^ c.LiteralOperand(instruction.operand)
		c.instruction_pointer += 2

	case 2: // bst
		c.registers['B'] = c.ComboOperand(instruction.operand) % 8
		c.instruction_pointer += 2

	case 3: // jnz
		if c.registers['A'] != 0 {
			c.instruction_pointer = c.LiteralOperand(instruction.operand)
		} else {
			// do nothing
			c.instruction_pointer += 2
		}

	case 4: // bxc
		c.registers['B'] = c.registers['B'] ^ c.registers['C']
		c.instruction_pointer += 2

	case 5: // out
		new_out := c.ComboOperand(instruction.operand) % 8
		c.outputs = append(c.outputs, new_out)
		fmt.Printf("Output: %03bb\n", new_out)
		c.instruction_pointer += 2

	case 6: // bdv
		numerator := c.registers['A']
		denominator := 1
		for i := 0; i < c.ComboOperand(instruction.operand); i++ {
			denominator *= 2
		}
		c.registers['B'] = numerator / denominator
		c.instruction_pointer += 2

	case 7: // cdv
		numerator := c.registers['A']
		denominator := 1
		for i := 0; i < c.ComboOperand(instruction.operand); i++ {
			denominator *= 2
		}
		c.registers['C'] = numerator / denominator
		c.instruction_pointer += 2
	}
}

func (c *Computer) String() string {

	s := fmt.Sprintf("Computer State: reg[A]: %d, reg[B] %d, reg[C]: %d,ip: %d ", c.registers['A'], c.registers['B'], c.registers['C'], c.instruction_pointer)
	s += fmt.Sprintf("outputs: %v\n", c.outputs)
	return s
}

func opcodeToString(opcode int) string {
	switch opcode {
	case 0:
		return "adv"
	case 1:
		return "bxl"
	case 2:
		return "bst"
	case 3:
		return "jnz"
	case 4:
		return "bxc"
	case 5:
		return "out"
	case 6:
		return "bdv"
	case 7:
		return "cdv"
	default:
		return "invalid opcode"
	}
}
func (i Instruction) String() string {
	s := fmt.Sprintf("%d,%d opcode: %s, operand: %d", i.opcode, i.operand, opcodeToString(i.opcode), i.operand)
	return s
}

func (c *Computer) RunNextInstruction() {
	opcode := c.program[c.instruction_pointer]
	operand := c.program[c.instruction_pointer+1]
	instruction := Instruction{opcode, operand}

	c.ExecuteInstruction(instruction)
	fmt.Println(c.String())
}

func (c *Computer) RunProgram() string {
	for c.instruction_pointer < len(c.program) {
		c.RunNextInstruction()
	}

	outputs_str := make([]string, 0)
	for _, val := range c.outputs {
		outputs_str = append(outputs_str, fmt.Sprintf("%d", val))
	}
	return strings.Join(outputs_str, ",")
}

func (c *Computer) GetOutput() []int {
	return c.outputs
}

func main() {
	// Read the input
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	registers, program := ParseInput(string(data))
	fmt.Println("registers: ", registers)
	fmt.Println("program: ", program)

	computer := Computer{registers, program, 0, make([]int, 0)}
	cout := computer.RunProgram()
	fmt.Println(cout)

	fmt.Println("Part 2: ")
	A := solveForA(program)
	fmt.Println("A: ", A)
}
