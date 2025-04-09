package main

import (
	"Assembler/validations"
	"regexp"
	"strconv"
	"strings"
)

func AnalizeSourceCode(sourceCode []string) {
	//Remove comments and empty lines
	ParsedSourceCode := removeComments(sourceCode)
	ParsedSourceCode = removeEmptyLines(ParsedSourceCode)
	ParsedSourceCode = removeSpaces(ParsedSourceCode)
	classificateLines(ParsedSourceCode)
}

func classificateLines(parsedSourceCode []string) {
	for lineNumber, line := range parsedSourceCode {

		elements := splitLine(line)

		for _, element := range elements {
			tempInst := classificateElement(element, strconv.Itoa(lineNumber+1))
			tempInst.AppendAsmInstruction()
		}
	}
}

func classificateElement(element string, lineNumber string) AsmInstruction {
	if isSegment, seg := validations.CheckSegment(element); isSegment {
		return SetAsmInstruction(element, seg, lineNumber)
	} else if isConstant, base := validations.CheckConstant(element); isConstant {
		tempSlice := []string{"Constante", base}
		return SetAsmInstruction(element, strings.Join(tempSlice, " "), lineNumber)
	} else if isInstruction, inst := validations.CheckInstruction(element); isInstruction {
		return SetAsmInstruction(element, inst, lineNumber)
	} else if isRegister, reg := validations.CheckRegister(element); isRegister {
		return SetAsmInstruction(element, reg, lineNumber)
	} else if isPseudoInstruction, psinst := validations.CheckPseudoInstruction(element); isPseudoInstruction {
		return SetAsmInstruction(element, psinst, lineNumber)
	}
	return SetAsmInstruction(element, "No valido", lineNumber)

}

func splitLine(line string) []string {
	stringsPtrns := regexp.MustCompile(`("+(.)+")|('+(.)+')`)
	segmentPtrns := regexp.MustCompile(`(?i)(code segment)|(?i)(data segment)|(?i)(stack segment)
	|(?i)(\.code segment)|(?i)(\.data segment)|(?i)(\.stack segment)|(?i)(\.model small)`)
	tempString := ""
	var tempSlice []string

	if stringsPtrns.MatchString(line) {
		tempString = stringsPtrns.FindAllString(line, -1)[0]
		line = strings.Replace(line, tempString, "", -1)
	}
	if segmentPtrns.MatchString(line) {
		return append(tempSlice, line)
	}

	tempSlice = strings.Split(line, " ")
	for i, element := range tempSlice {
		tempSlice[i] = strings.TrimSpace(element)
		tempSlice[i] = strings.TrimSuffix(tempSlice[i], ",")
	}

	tempSlice = removeEmptyLines(tempSlice)
	tempSlice = removeSliceSpaces(tempSlice)
	if tempString != "" {
		return append(tempSlice, tempString)
	}
	return tempSlice
}

func removeComments(sourceCode []string) []string {
	for i, line := range sourceCode {
		newLine := strings.Split(line, ";")[0]
		sourceCode[i] = newLine
	}
	return sourceCode
}

func removeSpaces(sourceCode []string) []string {
	for i, line := range sourceCode {
		newLine := strings.TrimSpace(line)
		sourceCode[i] = newLine
	}
	return sourceCode
}

func removeEmptyLines(sourceCode []string) []string {
	var r []string
	for _, str := range sourceCode {
		if str != "" {
			r = append(r, str)
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
