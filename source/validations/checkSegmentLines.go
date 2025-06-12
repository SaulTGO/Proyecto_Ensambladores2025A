package validations

// ValidateStackLine Return the Status and the error (if exists) using ONLY the classification
func ValidateStackLine(elementsClasif []string) (bool, string) {

	if len(elementsClasif) == 1 {
		if elementsClasif[0] == "PseudoInstruccion" {
			return true, ""
		}
	}

	if len(elementsClasif) == 3 {
		if elementsClasif[0] == "PseudoInstruccion" {
			if elementsClasif[1] == "Constante Decimal" {
				if elementsClasif[2] == "PseudoInstruccion" {
					return true, ""
				} else {
					return false, "Sintaxis incorrecta. (Ej: DB 10 DUP(0)"
				}
			} else {
				return false, "Sintaxis incorrecta. (Ej: DB 10 DUP(0)"
			}
		} else {
			return false, "Sintaxis incorrecta. (Ej: DB 10 DUP(0)"
		}
	}
	return false, ""
}

func ValidateDataLine(elementsClasif []string) (bool, string) {

	if len(elementsClasif) == 1 {
		if elementsClasif[0] == "PseudoInstruccion" {
			return true, ""
		}
	}

	// Check var declaration
	if len(elementsClasif) == 3 {
		if elementsClasif[0] == "Simbolo" {
			if elementsClasif[1] == "PseudoInstruccion" {
				if isValidConstant(elementsClasif[2]) {
					return true, ""
				}
			}
		}
	}

	//Check multiple var values
	if len(elementsClasif) >= 4 {
		if elementsClasif[0] == "Simbolo" {
			if elementsClasif[1] == "PseudoInstruccion" {
				// Check if every constant is valid
				for _, element := range elementsClasif[2:] {
					if isValidConstant(element) {
						continue
					} else {
						return false, "Constante no valida"
					}
				}
				return true, ""
			}
		}
	}

	return false, ""
}

func ValidateCodeLine(elementsClasif []string) (bool, string) {
	if len(elementsClasif) == 0 {
		return false, ""
	}

	if elementsClasif[0] == "No valido" {
		return false, "Instruccion invalida"
	}

	// Check tags or no operands instructions
	if len(elementsClasif) == 1 {
		if elementsClasif[0] == "Simbolo" {
			return true, ""
		} else if elementsClasif[0] == "Instruccion" {
			return true, ""
		} else if elementsClasif[0] == "PseudoInstruccion" {
			return true, ""
		} else {
			return false, "Intruccion no reconocida"
		}
	}

	// Check jumps and one operand instructions
	if len(elementsClasif) == 2 {
		if elementsClasif[0] == "Instruccion" {
			if elementsClasif[1] == "Simbolo" || elementsClasif[1] == "Registro" {
				return true, ""
			} else {
				return false, "Operando no reconocido"
			}
		} else {
			return false, "Elemento no reconocido"
		}
	}

	// Check two operands instructions
	if len(elementsClasif) == 3 {
		if elementsClasif[0] == "Instruccion" {
			return true, ""
		} else {
			return false, "Elemento no reconocido"
		}
	}
	return false, ""
}

var validConstants = []string{"Constante Decimal", "Constante Hexadecimal", "Constante Caracter", "Constante Octal", "Constante Binario"}

func isValidConstant(element string) bool {
	for _, valid := range validConstants {
		if element == valid {
			return true
		}
	}
	return false
}
