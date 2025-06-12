package main

var SourceCode = make(map[int]string)

var ParsedSourceCode = make(map[int]string)

var InstructionsSlice []AsmInstruction

type AsmInstruction struct {
	Name           string
	Classification string
	LineNumber     int
	Error          string
}

func (asmInst AsmInstruction) AppendAsmInstruction() {
	InstructionsSlice = append(InstructionsSlice, asmInst)
}

func SetAsmInstruction(name string, classification string, numberLine int, error string) AsmInstruction {
	tempInst := AsmInstruction{
		Name:           name,
		Classification: classification,
		LineNumber:     numberLine,
		Error:          error,
	}

	return tempInst
}

type Data struct {
	SourceCode        map[int]string
	InstructionsSlice []AsmInstruction
	FileProcessed     bool
}

// Fase 2

//var StackLines = make(map[int]string)
//var CodeLines = make(map[int]string)
//var DataLines = make(map[int]string)

type ProgramSegments struct {
	StackLines map[int]string
	DataLines  map[int]string
	CodeLines  map[int]string
}

var Segments ProgramSegments

type LineStatus struct {
	LineNumber int
	Line       string
	Status     string
	Error      string
}

var LinesValidation []LineStatus

type Symbol struct {
	Name  string
	Kind  string
	Value string
	Size  string
}

var SymbolTable []Symbol

type Fase2Data struct {
	SourceCode    map[int]string
	Validations   []LineStatus
	SymbolTable   []Symbol
	FileProcessed bool
}
