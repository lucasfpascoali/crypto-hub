package classical

import "errors"

type ShiftCipher struct {
	Key       byte
	MatchCase bool
}

func NewShiftCipher(key uint64, matchCase bool) *ShiftCipher {
	return &ShiftCipher{Key: byte(key % 26), MatchCase: matchCase}
}

func (s *ShiftCipher) Encrypt(plaintext []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return nil, errors.New("plaintext is empty")
	}

	var ciphertext = make([]byte, len(plaintext))
	for i, b := range plaintext {
		if !isAlphabetical(b) {
			continue
		}

		startIndex := lowerCaseStartIndex
		if s.MatchCase && b <= upperCaseEndIndex {
			startIndex = upperCaseStartIndex
		}

		ciphertext[i] = startIndex + ((b - startIndex + s.Key) % 26)
	}

	return ciphertext, nil
}

func (s *ShiftCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, errors.New("ciphertext is empty")
	}

	var plaintext = make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		if !isAlphabetical(b) {
			continue
		}

		startIndex := lowerCaseStartIndex
		if s.MatchCase && b <= upperCaseEndIndex {
			startIndex = upperCaseStartIndex
		}

		plaintext[i] = startIndex + ((b - startIndex - s.Key) % 26)
	}

	return plaintext, nil
}
