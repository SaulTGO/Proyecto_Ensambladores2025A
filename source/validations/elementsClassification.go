package validations

import (
	"Assembler/references"
	"regexp"
	"strings"
)

func CheckConstant(element string) (bool, string) {
	base := ""
	constantPtrn := regexp.MustCompile(`(^\d)|(^")|(^')`)
	if constantPtrn.MatchString(element) {
		base = checkBaseConstant(element)
		return true, base
	}
	return false, ""

}

func checkBaseConstant(element string) string {
	baseConstantPtrn := regexp.MustCompile(`\w$`)
	base := ""
	if baseConstantPtrn.MatchString(element) {
		b := baseConstantPtrn.FindString(element)

		if strings.EqualFold(b, "b") {
			base = "Binario"
		} else if strings.EqualFold(b, "h") {
			base = "Hexadecimal"
		} else if strings.EqualFold(b, "o") {
			base = "Octal"
		} else if strings.EqualFold(b, "") {

		} else {
			base = "Decimal"
		}
	}
	return base
}

func CheckInstruction(element string) (bool, string) {
	_, found := references.Instructions[element]
	if found {
		return true, "Instruccion"
	}
	return false, ""
}

func CheckRegister(element string) (bool, string) {
	_, found := references.Registers16bits[element]
	if found {
		return true, "Registro"
	} else if _, found := references.Registers8bits[element]; found {
		return true, "Registro"
	}
	return false, ""
}

func CheckPseudoInstruction(element string) (bool, string) {
	_, found := references.PseudoInstructions[element]
	if found {
		return true, "PseudoInstruccion"
	}
	return false, ""
}

func CheckSegment(element string) (bool, string) {
	segmentPtrns := regexp.MustCompile(`(?i)(code segment)|(?i)(data segment)|(?i)(stack segment)
	|(?i)(\.code segment)|(?i)(\.data segment)|(?i)(\.stack segment)|(?i)(\.model small)`)

	if segmentPtrns.MatchString(element) {
		return true, "PseudoInstruccion"
	}
	return false, ""
}
