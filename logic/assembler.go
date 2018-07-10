package logic

type Assembler interface {
	Assemble(content string) error
}
