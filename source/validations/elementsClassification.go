package validations

import (
	"Assembler/references"
	"strings"
)

func CheckConstant(element string) (bool, string, string) {
	if !references.ConstantPtrn.MatchString(element) {
		return false, "", ""
	}

	base, err := checkBaseConstant(element)
	if base == "No valido" || base == "Bad Quotes" {
		return true, base, err
	}

	return true, base, ""
}

func checkBaseConstant(element string) (string, string) {

	if references.StringPtrn.MatchString(element) {
		if references.EndQuotesPtrn.MatchString(element) {
			return "Caracter", ""
		}
		return "Bad Quotes", "Cadena mal declarada"
	}

	if references.BaseConstantPtrn.MatchString(element) {
		b := references.BaseConstantPtrn.FindString(element)

		switch strings.ToUpper(b) {
		case "B":
			return checkBinary(element)
		case "H":
			return checkHexa(element)
		case "O":
			return "Octal", ""
		default:
			return "Decimal", ""
		}
	}

	return "No valido", "Constante no reconocida"
}

func CheckInstruction(element string) (bool, string) {
	_, found := references.Instructions[strings.ToUpper(element)]
	if found {
		return true, "Instruccion"
	}
	return false, ""
}

func CheckRegister(element string) (bool, string) {
	_, fullMatch := references.Registers16bits[strings.ToUpper(element)]
	_, match := references.Registers8bits[strings.ToUpper(element)]
	if fullMatch || match {
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

func checkBinary(element string) (string, string) {
	lastChar := string(element[len(element)-1])
	if lastChar == "b" || lastChar == "B" {
		tempString := element[:len(element)-1]
		if len(tempString) == 8 || len(tempString) == 16 {
			return "Binario", ""
		}
		return "No valido", "Longitud incorrecta (1 o 2 bytes)"
	}
	return "No valido", "Formato binario incorrecto"
}

func checkHexa(element string) (string, string) {

	if string(element[0]) == "0" {
		tempString := element[1:]
		tempString = tempString[:len(tempString)-1]

		validLengths := []int{2, 4, 8}
		for _, length := range validLengths {
			if len(tempString) == length {
				return "Hexadecimal", ""
			}
		}
		return "No valido", "Longitud incorrecta (1 o 2 bytes)"
	}
	return "No valido", "Formato incorrecto (debe iniciar con 0)"
}
