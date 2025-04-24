package validations

import (
	"Assembler/references"
	"strings"
)

func CheckConstant(element string) (bool, string) {
	base := ""

	if references.ConstantPtrn.MatchString(element) {
		base = checkBaseConstant(element)
		if base == "No valido" {
			return false, "No valido"
		}
		return true, base
	}
	return false, ""
}

func checkBaseConstant(element string) string {
	if references.StringPtrn.MatchString(element) {
		if references.EndQuotesPtrn.MatchString(element) {
			return "Caracter"
		}
		return "Bad Quotes"
	} else if references.BaseConstantPtrn.MatchString(element) {
		b := references.BaseConstantPtrn.FindString(element)

		if strings.EqualFold(b, "b") || strings.EqualFold(b, "B") {
			return "Binario"
		} else if strings.EqualFold(b, "h") || strings.EqualFold(b, "H") {
			return checkHexa(element)
		} else if strings.EqualFold(b, "o") || strings.EqualFold(b, "O") {
			return "Octal"
		} else {
			return "Decimal"
		}
	}
	return ""
}

func CheckInstruction(element string) (bool, string) {
	_, found := references.Instructions[strings.ToUpper(element)]
	if found {
		return true, "Instruccion"
	}
	return false, ""
}

func CheckRegister(element string) (bool, string) {
	_, found := references.Registers16bits[strings.ToUpper(element)]
	if found {
		return true, "Registro"
	} else if _, found := references.Registers8bits[element]; found {
		return true, "Registro"
	}
	return false, ""
}

func CheckPseudoInstruction(element string) (bool, string) {
	_, found := references.PseudoInstructions[strings.ToUpper(element)]
	isPseudo := references.PseudoElementsPtrn.MatchString(strings.ToUpper(element))
	if found || isPseudo {
		return true, "PseudoInstruccion"
	}
	if references.DupPtrn.MatchString(element) {
		return true, "PseudoInstruccion"
	}
	return false, ""
}

func CheckSegment(element string) (bool, string) {
	if references.PseudoElementsPtrn.MatchString(element) {
		return true, "PseudoInstruccion"
	}
	return false, ""
}

func CheckInvalidInstruction(element string) (bool, string) {
	_, found := references.InvalidInstructions[strings.ToUpper(element)]
	if found {
		return true, "No valido"
	}
	return false, ""
}

func CheckSymbol(element string) (bool, string) {
	if references.SymbolPtrn.MatchString(element) {
		return true, "Simbolo"
	}
	return false, ""
}

func checkHexa(element string) string {
	if string(element[0]) == "0" {
		tempString := element[1:]
		tempString = tempString[:len(tempString)-1]
		if len(tempString) == 2 || len(tempString) == 4 || len(tempString) == 8 {
			return "Hexadecimal"
		}
	}
	return "No valido"
}
