package classical

import (
	"errors"
	"fmt"
)

type SubstitutionCipher struct {
	Key        map[byte]byte
	ReverseKey map[byte]byte
	MatchCase  bool
}

func NewSubstitutionCipher(key []byte, matchCase bool) (*SubstitutionCipher, error) {
	keyMap, reverseKeyMap, err := buildKeys(key)
	if err != nil {
		return nil, err
	}

	return &SubstitutionCipher{Key: keyMap, ReverseKey: reverseKeyMap, MatchCase: matchCase}, nil
}

func (s *SubstitutionCipher) Encrypt(plaintext []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return nil, errors.New("plaintext is empty")
	}

	return substituteFromMap(plaintext, s.Key, s.MatchCase)
}

func (s *SubstitutionCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, errors.New("ciphertext is empty")
	}

	return substituteFromMap(ciphertext, s.ReverseKey, s.MatchCase)
}

// DFBGHIMLORZXYTUACQPNEJKSVW
// ABCDEFGHIJKLMNOPQRSTUVWXYZ

func buildKeys(key []byte) (map[byte]byte, map[byte]byte, error) {
	if len(key) != 26 {
		return nil, nil, fmt.Errorf("key length must be 26 characters, got %d", len(key))
	}

	alphaMap := make(map[byte]byte, 26)
	reverseMap := make(map[byte]byte, 26)
	for i, c := range key {
		if !isAlphabetical(c) {
			return nil, nil, fmt.Errorf("key contains non-alphabetical character %c", c)
		}

		index := c
		if index <= upperCaseEndIndex {
			index = c - upperCaseStartIndex + lowerCaseStartIndex
		}

		if _, ok := reverseMap[index]; ok {
			return nil, nil, fmt.Errorf("key contains duplicate alpha character at position %d", i)
		}

		alphaMap[byte(i)+lowerCaseStartIndex] = index
		reverseMap[index] = byte(i) + lowerCaseStartIndex
	}

	return alphaMap, reverseMap, nil
}

func substituteFromMap(text []byte, m map[byte]byte, matchCase bool) ([]byte, error) {
	result := make([]byte, len(text))
	for i, b := range text {
		if !isAlphabetical(b) {
			continue
		}

		index := b
		isUpperCase := b <= upperCaseEndIndex
		if isUpperCase {
			index += lowerCaseStartIndex - upperCaseStartIndex
		}

		val, ok := m[index]
		if !ok {
			return nil, fmt.Errorf("key not found")
		}

		if matchCase && isUpperCase {
			val -= lowerCaseStartIndex - upperCaseStartIndex
		}

		result[i] = val
	}

	return result, nil
}
