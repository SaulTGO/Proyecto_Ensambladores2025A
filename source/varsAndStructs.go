package main

var SourceCode = make(map[int]string)

//var SourceCode []string

//var ParsedSourceCode []asmInstruction

var InstructionsSlice []AsmInstruction

type AsmInstruction struct {
	Name           string
	Classification string
	LineNumber     string
	ElementNumber  int
}

func (asmInst AsmInstruction) AppendAsmInstruction() {
	InstructionsSlice = append(InstructionsSlice, asmInst)
}

func SetAsmInstruction(name string, classification string, numberLine string, elementNumber int) AsmInstruction {
	tempInst := AsmInstruction{
		Name:           name,
		Classification: classification,
		LineNumber:     numberLine,
		ElementNumber:  elementNumber,
	}

	return tempInst
}

type Data struct {
	SourceCode        map[int]string
	InstructionsSlice []AsmInstruction
	FileProcessed     bool
}
