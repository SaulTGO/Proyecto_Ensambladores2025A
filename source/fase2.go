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

// splitSegments Return each line program segments. stackLines, codeLines, dataLines
func splitSegments() {

	var stackLines = make(map[int]string)
	var codeLines = make(map[int]string)
	var dataLines = make(map[int]string)

	var currentSection = ""

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
			stackLines[lineNumber] = line
		} else if currentSection == "Code" {
			codeLines[lineNumber] = line
		} else if currentSection == "Data" {
			dataLines[lineNumber] = line
		}

	}

	Segments.StackLines = stackLines
	Segments.DataLines = dataLines
	Segments.CodeLines = codeLines

	//println("STACK LINES")
	//for lineNumber, line := range stackLines {
	//	println(lineNumber, line)
	//}
	//println("DATA LINES")
	//for lineNumber, line := range dataLines {
	//	println(lineNumber, line)
	//}
	//println("CODE LINES")
	//for lineNumber, line := range codeLines {
	//	println(lineNumber, line)
	//}

}

func validateLines() []LineStatus {

	var tempLineStatus []LineStatus

	// Sort map keys to append correctly
	sortedStackKeys := SortKeys(Segments.StackLines)
	sortedDataKeys := SortKeys(Segments.DataLines)
	sortedCodeKeys := SortKeys(Segments.CodeLines)

	segmentFlag := false
	codeFlag := false
	dataFlag := false

	//Validate Stack Lines
	if len(Segments.StackLines) > 0 {
		for _, lineNumber := range sortedStackKeys {
			line := Segments.StackLines[lineNumber]
			lineElements := getLineElements(lineNumber)
			elementsClassif := getElementsClassif(lineElements)

			if references.StackSegmentPtrn.MatchString(line) && segmentFlag == false {
				segmentFlag = true
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", ""})
				continue
			} else if references.BadStackSegmentPtrn.MatchString(line) && segmentFlag == true {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", "Inicio de segmento ya declarado"})
				continue
			}

			if exists, er := ExistInvalidElement(lineElements, elementsClassif); exists {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", er})
			} else if status, err := validations.ValidateStackLine(elementsClassif); status {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", err})
			} else {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", err})
			}
		}
	}

	//Validate Data Lines
	if len(Segments.DataLines) > 0 {
		for _, lineNumber := range sortedDataKeys {
			line := Segments.DataLines[lineNumber]
			lineElements := getLineElements(lineNumber)
			elementsClassif := getElementsClassif(lineElements)

			if references.DataSegmentPtrn.MatchString(line) && dataFlag == false {
				dataFlag = true
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", ""})
				continue
			} else if references.BadDataSegmentPtrn.MatchString(line) && dataFlag == true {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", "Inicio de segmento ya declarado"})
				continue
			}

			if exists, er := ExistInvalidElement(lineElements, elementsClassif); exists {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", er})
			} else if status, err := validations.ValidateDataLine(elementsClassif); status {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", err})
				if lineElements[0].Classification == "Simbolo" && lineElements[1].Classification == "PseudoInstruccion" {
					AddToSymbolTable(lineElements, false)
				}
			} else {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", err})
			}
		}
	}

	//Validate Code Lines
	if len(Segments.CodeLines) > 0 {
		for _, lineNumber := range sortedCodeKeys {
			line := Segments.CodeLines[lineNumber]
			lineElements := getLineElements(lineNumber)
			elementsClassif := getElementsClassif(lineElements)

			if references.CodeSegmentPtrn.MatchString(line) && codeFlag == false {
				codeFlag = true
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", ""})
				continue
			} else if references.BadCodeSegmentPtrn.MatchString(line) && codeFlag == true {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", "Inicio de segmento ya declarado"})
				continue
			}

			if exists, er := ExistInvalidElement(lineElements, elementsClassif); exists {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", er})
			} else if status, err := validations.ValidateCodeLine(elementsClassif); status {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Correcto", err})
				if lineElements[0].Classification == "Simbolo" {
					AddToSymbolTable(lineElements, true)
				}
			} else {
				tempLineStatus = append(tempLineStatus, LineStatus{lineNumber, line, "Error", err})
			}
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
	if len(lineElements) == 0 {
		return
	}

	if isTag {
		symbol := lineElements[0].Name //The symbol s name is the first element in the line
		SymbolTable = append(SymbolTable, Symbol{symbol, "Etiqueta", "", ""})
	} else {
		symbol := lineElements[0].Name         //The symbol s name is the first element in the line
		kind := lineElements[2].Classification //Get the type of the symbol (constant type)
		value := lineElements[2].Name          //Get the actual value

		// For variables who have more than 1 constant value
		if len(lineElements) > 3 {
			valuesSlice := lineElements[2:]
			values := ""
			for _, val := range valuesSlice {
				values = values + val.Name + ", "
			}

			value = values
		}

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

func ExistInvalidElement(lineElmts []AsmInstruction, elmtsClassif []string) (bool, string) {
	for elmtNumber, elmtClassif := range elmtsClassif {
		if elmtClassif == "No valido" {
			return true, "<" + lineElmts[elmtNumber].Name + ">" + " " + lineElmts[elmtNumber].Error
		}
	}
	return false, ""
}
