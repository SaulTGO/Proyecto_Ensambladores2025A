package validations

import (
	"Assembler/references"
	"strings"
)

func CheckConstant(element string) (bool, string) {
	base := ""

	if references.ConstantPtrn.MatchString(element) {
		base = checkBaseConstant(element)
		return true, base
	}
	return false, ""

}

func checkBaseConstant(element string) string {

	base := ""
	if references.BaseConstantPtrn.MatchString(element) {
		b := references.BaseConstantPtrn.FindString(element)

		if strings.EqualFold(b, "b") {
			base = "Binario"
		} else if strings.EqualFold(b, "h") {
			base = "Hexadecimal"
		} else if strings.EqualFold(b, "o") {
			base = "Octal"
		} else {
			base = "Decimal"
		}
	} else if references.StringPtrn.MatchString(element) {
		base = "Caracter"
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
	isPseudo := references.PseudoElementsPtrn.MatchString(element)
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
	_, found := references.InvalidInstructions[element]
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
