package polymonrphism_test

import (
	"fmt"
	"testing"
)

type code string

type Programmer interface {
	WriteHelloWorld() code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() code {
	return "fmt.Println(\"Hello world!\")"
}

type PhpProgrammer struct {
}

func (p *PhpProgrammer) WriteHelloWorld() code {
	return "echo Hello world!"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestDuotai(t *testing.T) {
	php := new(PhpProgrammer)
	// php := &PhpProgrammer{}
	goProg := new(GoProgrammer)
	writeFirstProgram(php)
	writeFirstProgram(goProg)
}
