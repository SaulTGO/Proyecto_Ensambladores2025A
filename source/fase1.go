package main

import (
	"Assembler/references"
	"Assembler/validations"
	"sort"
	"strconv"
	"strings"
)

func AnalizeSourceCode(sourceCode map[int]string) {
	//Remove comments and empty lines
	ParsedSourceCode := removeComments(sourceCode)
	ParsedSourceCode = removeEmptyLines(ParsedSourceCode)
	ParsedSourceCode = removeSpaces(ParsedSourceCode)
	classificateLines(ParsedSourceCode)
}

func classificateLines(parsedSourceCode map[int]string) {
	sortedKeys := SortKeys(parsedSourceCode)
	elementNumber := 0
	var elmNumber *int = &elementNumber

	for _, k := range sortedKeys {
		line := parsedSourceCode[k]
		lineNumber := k
		elements := splitLine(line)

		for _, element := range elements {
			*elmNumber = *elmNumber + 1
			tempInst := classificateElement(element, strconv.Itoa(lineNumber+1), elementNumber)
			tempInst.AppendAsmInstruction()
		}
	}
}

func classificateElement(element string, lineNumber string, elementNumber int) AsmInstruction {
	if isSegment, seg := validations.CheckSegment(element); isSegment {
		return SetAsmInstruction(element, seg, lineNumber, elementNumber)
	} else if isConstant, base := validations.CheckConstant(element); isConstant {
		tempSlice := []string{"Constante", base}
		return SetAsmInstruction(element, strings.Join(tempSlice, " "), lineNumber, elementNumber)
	} else if isInstruction, inst := validations.CheckInstruction(element); isInstruction {
		return SetAsmInstruction(element, inst, lineNumber, elementNumber)
	} else if isRegister, reg := validations.CheckRegister(element); isRegister {
		return SetAsmInstruction(element, reg, lineNumber, elementNumber)
	} else if isPseudoInstruction, psinst := validations.CheckPseudoInstruction(element); isPseudoInstruction {
		return SetAsmInstruction(element, psinst, lineNumber, elementNumber)
	} else if isInvalidInstruction, invInst := validations.CheckInvalidInstruction(element); isInvalidInstruction {
		return SetAsmInstruction(element, invInst, lineNumber, elementNumber)
	} else if isSimbol, symbol := validations.CheckSymbol(element); isSimbol {
		return SetAsmInstruction(element, symbol, lineNumber, elementNumber)
	}
	return SetAsmInstruction(element, "No valido", lineNumber, elementNumber)

}

func splitLine(line string) []string {
	tempString := ""
	var tempSlice []string

	if references.FullStringPtrn.MatchString(line) {
		tempString = references.FullStringPtrn.FindAllString(line, -1)[0]
		line = strings.Replace(line, tempString, "", -1)
	}
	if references.PseudoElementsPtrn.MatchString(line) {
		return append(tempSlice, line)
	}

	tempSlice = strings.Split(line, " ")
	for i, element := range tempSlice {
		tempSlice[i] = strings.TrimSpace(element)
		tempSlice[i] = strings.TrimSuffix(tempSlice[i], ",")
		tempSlice[i] = strings.TrimSuffix(tempSlice[i], ":")
	}

	tempSlice = removeSliceLines(tempSlice)
	tempSlice = removeSliceSpaces(tempSlice)
	if tempString != "" {
		return append(tempSlice, tempString)
	}
	return tempSlice
}

func removeComments(sourceCode map[int]string) map[int]string {
	for i, line := range sourceCode {
		newLine := strings.Split(line, ";")[0]
		sourceCode[i] = newLine
	}
	return sourceCode
}

func removeSpaces(sourceCode map[int]string) map[int]string {
	for i, line := range sourceCode {
		newLine := strings.TrimSpace(line)
		sourceCode[i] = newLine
	}
	return sourceCode
}

func removeEmptyLines(sourceCode map[int]string) map[int]string {
	var r = make(map[int]string)
	for i, str := range sourceCode {
		if str != "" {
			r[i+1] = str
		}
	}
	return r
}

func removeSliceSpaces(parsedSourceCode []string) []string {
	var r []string
	for _, str := range parsedSourceCode {
		if true {
			r = append(r, str)
		}
	}
	return r
}

func removeSliceLines(parsedSourceCode []string) []string {
	var r []string
	for _, str := range parsedSourceCode {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func SortKeys(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
