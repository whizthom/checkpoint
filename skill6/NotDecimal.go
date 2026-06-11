package main


// Instructions
// Write a function called NotDecimal() that takes as an argument a string in form of a float number with the 
// decimal point and returns a string converted to int without the decimal point 
// (you will have to multiply it by 10^n to remove the .).

// If the number doesn't have a decimal point or there is only a zero after the . return the number followed by a newline \n.
// If the argument is empty return a newline \n.
// If the argument is not a number return it followed by a newline \n.

func NotDecimal(dec string) string {
	// empty string → return newline
	if len(dec) == 0 {
		return "\n"
	}

	// find the decimal point
	dotIndex := -1
	for i, ch := range dec {
		if ch == '.' {
			dotIndex = i
			break
		}
	}

	// no decimal point → return as is
	if dotIndex == -1 {
		return dec + "\n"
	}

	// get the part before and after the dot
	whole := dec[:dotIndex]
	decimal := dec[dotIndex+1:]

	// validate: check every character is a digit (allowing leading minus)
	for i, ch := range dec {
		if ch == '-' && i == 0 {
			continue
		}
		if ch == '.' {
			continue
		}
		if ch < '0' || ch > '9' {
			return dec + "\n"
		}
	}

	// if decimal part is empty or only zeros → return whole part
	allZero := true
	for _, ch := range decimal {
		if ch != '0' {
			allZero = false
			break
		}
	}
	if allZero {
		return whole + "\n"
	}

	// remove trailing zeros from decimal
	end := len(decimal)
	for end > 0 && decimal[end-1] == '0' {
		end--
	}
	decimal = decimal[:end]

	// if whole is just "0" or "-0" drop the zero
	if whole == "0" {
		whole = ""
	} else if whole == "-0" {
		whole = "-"
	}

	return whole + decimal + "\n"
}