package classical

const upperCaseStartIndex = byte(65)
const upperCaseEndIndex = byte(90)
const lowerCaseStartIndex = byte(97)
const lowerCaseEndIndex = byte(122)

func isAlphabetical(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}
