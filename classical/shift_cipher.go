package classical

import "errors"

type ShiftCipher struct {
	Key       int
	MatchCase bool
}

func NewCaesarCipher(matchCase bool) *ShiftCipher {
	return NewShiftCipher(3, matchCase)
}

func NewShiftCipher(key uint64, matchCase bool) *ShiftCipher {
	return &ShiftCipher{Key: mod(int(key), alphabetSize), MatchCase: matchCase}
}

func (s *ShiftCipher) Encrypt(plaintext []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return nil, errors.New("plaintext is empty")
	}

	return shiftLetters(plaintext, s.Key, s.MatchCase), nil
}

func (s *ShiftCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, errors.New("ciphertext is empty")
	}

	return shiftLetters(ciphertext, -s.Key, s.MatchCase), nil
}

func shiftLetters(text []byte, shift int, matchCase bool) []byte {
	result := make([]byte, len(text))
	for i, b := range text {
		if !isAlphabetical(b) {
			result[i] = b
			continue
		}

		caseIndex := lowerCaseStartIndex
		if b <= upperCaseEndIndex {
			caseIndex = upperCaseStartIndex
		}

		if matchCase {
			result[i] = byte(caseIndex + mod(int(b)-caseIndex+shift, alphabetSize))
			continue
		}

		result[i] = byte(lowerCaseStartIndex + mod(int(b)-caseIndex+shift, alphabetSize))
	}

	return result
}
