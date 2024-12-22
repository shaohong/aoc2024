package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	data := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	registers, program := ParseInput(data)
	// fmt.Println(registers)
	// fmt.Println(program)

	if registers['A'] != 729 || registers['B'] != 0 || registers['C'] != 0 {
		t.Error("Register A should be 729")
	}

	expected_program := []int{0, 1, 5, 4, 3, 0}
	if !reflect.DeepEqual(program, expected_program) {
		t.Errorf("got %v, expect %v", program, expected_program)
	}
}

func TestRunProgram(t *testing.T) {
	data := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	registers, program := ParseInput(data)
	computer := Computer{registers: registers, program: program}

	fmt.Println(computer.String())

	// for i := 1; i <= 3; i++ {
	// 	computer.RunNextInstruction()
	// 	fmt.Println(computer.String())
	// }

	computer.RunProgram()
	got := computer.GetOutput()
	expected := []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expect %v", got, expected)
	}

}

func TestOctals(t *testing.T) {

	const a int = 034530
	const b int = 034530 * 8
	fmt.Println(a, b)
}
