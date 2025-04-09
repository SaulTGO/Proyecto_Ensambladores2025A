package main

// var SourceCode map[string]string
var SourceCode []string

//var ParsedSourceCode []asmInstruction

var InstructionsSlice []AsmInstruction

type AsmInstruction struct {
	Name           string
	Classification string
	LineNumber     string
}

func (asmInst AsmInstruction) AppendAsmInstruction() {
	InstructionsSlice = append(InstructionsSlice, asmInst)
}

func SetAsmInstruction(name string, classification string, numberLine string) AsmInstruction {
	tempInst := AsmInstruction{
		Name:           name,
		Classification: classification,
		LineNumber:     numberLine,
	}

	return tempInst
}

type Data struct {
	SourceCode        []string
	InstructionsSlice []AsmInstruction
	FileProcessed     bool
}
