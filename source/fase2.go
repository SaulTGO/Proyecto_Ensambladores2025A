package main

import (
	"Assembler/references"
	"Assembler/validations"
	"strings"
)

func SemanticAnalysis() {

	splitSegments() // Fill the var: Segments
	LinesValidation = validateLines()

}

func splitSegments() {
	var currentSection string
	sortedKeys := SortKeys(ParsedSourceCode)

	for _, lineNumber := range sortedKeys {
		line := ParsedSourceCode[lineNumber]

		if references.StackSegmentPtrn.MatchString(line) {
			currentSection = "Stack"
		} else if references.CodeSegmentPtrn.MatchString(line) {
			currentSection = "Code"
		} else if references.DataSegmentPtrn.MatchString(line) {
			currentSection = "Data"
		}

		if currentSection == "Stack" {
			StackLines[lineNumber] = line
		} else if currentSection == "Code" {
			CodeLines[lineNumber] = line
		} else if currentSection == "Data" {
			DataLines[lineNumber] = line
		}

	}

	Segments.StackLines = StackLines
	Segments.DataLines = DataLines
	Segments.CodeLines = CodeLines

}

func validateLines() []LineStatus {

	var tempLineStatus []LineStatus

	// Sort map keys to append correctly
	sortedStackKeys := SortKeys(Segments.StackLines)
	sortedDataKeys := SortKeys(Segments.DataLines)
	sortedCodeKeys := SortKeys(Segments.CodeLines)

	//Validate Stack Lines
	for _, lineNumber := range sortedStackKeys {
		line := Segments.StackLines[lineNumber]
		lineElements := getLineElements(lineNumber)
		elementsClassif := getElementsClassif(lineElements)

		if status, err := validations.ValidateStackLine(elementsClassif); status {
			tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", err})
		} else {
			tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", err})
		}
	}

	//Validate Data Lines
	for _, lineNumber := range sortedDataKeys {
		line := Segments.DataLines[lineNumber]
		lineElements := getLineElements(lineNumber)
		elementsClassif := getElementsClassif(lineElements)

		if status, err := validations.ValidateDataLine(elementsClassif); status {
			tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", err})
			if lineElements[0].Classification == "Simbolo" && lineElements[1].Classification == "PseudoInstruccion" {
				AddToSymbolTable(lineElements, false)
			}
		} else {
			tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", err})
		}
	}

	//Validate Code Lines
	for _, lineNumber := range sortedCodeKeys {
		line := Segments.CodeLines[lineNumber]
		lineElements := getLineElements(lineNumber)
		elementsClassif := getElementsClassif(lineElements)

		if status, err := validations.ValidateCodeLine(elementsClassif); status {
			tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", err})
			if lineElements[0].Classification == "Simbolo" {
				AddToSymbolTable(lineElements, true)
			}
		} else {
			tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", err})
		}
	}

	return tempLineStatus

}

func getLineElements(lineNumber int) []AsmInstruction {
	var tempSlice []AsmInstruction
	for _, element := range InstructionsSlice {
		if element.LineNumber == lineNumber {
			tempSlice = append(tempSlice, element)
			continue
		}
	}
	return tempSlice
}

func getElementsClassif(lineElements []AsmInstruction) []string {
	var tempSlice []string

	for _, element := range lineElements {
		tempSlice = append(tempSlice, element.Classification)
	}

	return tempSlice
}

func AddToSymbolTable(lineElements []AsmInstruction, isTag bool) {

	if isTag {
		symbol := lineElements[0].Name //The symbol s name is the first element in the line
		SymbolTable = append(SymbolTable, Symbol{symbol, "Etiqueta", "", ""})
	} else {
		symbol := lineElements[0].Name         //The symbol s name is the first element in the line
		kind := lineElements[2].Classification //Get the type of the symbol (constant type)
		value := lineElements[2].Name          //Get the actual value

		var size = ""
		var toSize = &size

		if strings.ToUpper(lineElements[1].Name) == "DB" {
			*toSize = "Byte"
		} else if strings.ToUpper(lineElements[1].Name) == "DW" {
			*toSize = "Word"
		}

		SymbolTable = append(SymbolTable, Symbol{symbol, kind, value, size})
	}
}
