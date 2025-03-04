package classical

const alphabetSize = 26
const upperCaseStartIndex = 65
const upperCaseEndIndex = 90
const lowerCaseStartIndex = 97
const lowerCaseEndIndex = 122

func isAlphabetical(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}

func mod(n, m int) int {
	return (n%m + m) % m
}
