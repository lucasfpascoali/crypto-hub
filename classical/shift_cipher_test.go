package classical

import "testing"

const caesarCipherExpected = 3

func TestNewCaesarCipher(t *testing.T) {
	var tests = []struct {
		MatchCase bool
	}{
		{true},
		{false},
	}

	for _, test := range tests {
		c := NewCaesarCipher(test.MatchCase)
		if c == nil {
			t.Error("NewCaesarCipher returned nil")
			return
		}

		if c.Key != caesarCipherExpected {
			t.Errorf("NewCaesarCipher returned wrong key. Expected %d, got %d", caesarCipherExpected, c.Key)
		}

		if c.MatchCase != test.MatchCase {
			t.Errorf("NewCaesarCipher returned wrong MatchCase. Expected %t, got %t", test.MatchCase, c.MatchCase)
		}
	}
}

func TestNewShiftCipher(t *testing.T) {
	var tests = []struct {
		InputKey    uint64
		ExpectedKey int
		MatchCase   bool
	}{
		{1, 1, false},
		{13, 13, true},
		{25, 25, true},
		{10000, 10000 % 26, true},
		{26, 0, true},
		{0, 0, false},
		{27, 1, false},
	}

	for _, test := range tests {
		s := NewShiftCipher(test.InputKey, test.MatchCase)
		if s == nil {
			t.Error("NewShiftCipher returned nil")
			return
		}

		if s.Key != test.ExpectedKey {
			t.Errorf("NewShiftCipher returned wrong Key. Expected %d, got %d", test.ExpectedKey, s.Key)
		}

		if s.MatchCase != test.MatchCase {
			t.Errorf("NewShiftCipher returned wrong MatchCase. Expected %t, got %t", test.MatchCase, s.MatchCase)
		}
	}
}

func TestShiftCipher_Encrypt(t *testing.T) {
	tests := []struct {
		s         *ShiftCipher
		plaintext []byte
		isError   bool
	}{
		{NewShiftCipher(1, false), []byte(""), true},
		{NewShiftCipher(1, false), []byte("ThisISaTest"), false},
		{NewCaesarCipher(false), []byte(""), true},
		{NewCaesarCipher(false), []byte("abcdefghijklmnopqrstuvwxyz"), false},
	}

	for _, test := range tests {
		ciphertext, err := test.s.Encrypt(test.plaintext)
		if test.isError && err == nil {
			t.Error("Encrypt returned nil error when it shouldn't have")
			return
		}

		if !test.isError && err != nil {
			t.Errorf("Encrypt returned unexpected error when it shouldn't have: %s", err)
			return
		}

		expectedCiphertext := string(shiftLetters(test.plaintext, test.s.Key, test.s.MatchCase))
		if string(ciphertext) != expectedCiphertext {
			t.Error("Encrypt returned a value different than shiftLetters function with same parameters")
		}
	}
}

func TestShiftCipher_Decrypt(t *testing.T) {
	tests := []struct {
		s          *ShiftCipher
		ciphertext []byte
		isError    bool
	}{
		{NewShiftCipher(1, false), []byte(""), true},
		{NewShiftCipher(1, false), []byte("UijtJTbUftu"), false},
		{NewCaesarCipher(false), []byte(""), true},
		{NewCaesarCipher(false), []byte("dEFghIJklm NOpqrSTu vwxyzabc"), false},
	}

	for _, test := range tests {
		plaintext, err := test.s.Encrypt(test.ciphertext)
		if test.isError && err == nil {
			t.Error("Encrypt returned nil error when it shouldn't have")
			return
		}

		if !test.isError && err != nil {
			t.Errorf("Encrypt returned unexpected error when it shouldn't have: %s", err)
			return
		}

		expectedPlaintext := string(shiftLetters(test.ciphertext, test.s.Key, test.s.MatchCase))
		if string(plaintext) != expectedPlaintext {
			t.Error("Encrypt returned a value different than shiftLetters function with same parameters")
		}
	}
}

func TestShiftLetters(t *testing.T) {
	tests := []struct {
		text      string
		result    string
		shift     int
		matchCase bool
	}{
		{"", "", 0, false},
		{"ThisISaTest", "uijtjtbuftu", 1, false},
		{"ThisISaTest", "UijtJTbUftu", 1, true},
		{"HelloWorld", "helloworld", 26, false},
		{"HelloWorld", "HelloWorld", 26, true},
		{"we will meet at midnight", "hp htww xppe le xtoytrse", 11, false},
		{"wE will Meet at midnighT", "hP htww Xppe le xtoytrsE", 11, true},
		{"abcdefghijklmnopqrstuvwxyz", "defghijklmnopqrstuvwxyzabc", 3, false},
		{"aBCdeFGhij KLmnoPQr stuvwxyz", "dEFghIJklm NOpqrSTu vwxyzabc", 3, true},
		{"uijtjtbuftu", "thisisatest", -1, false},
		{"UijtJTbUftu", "ThisISaTest", -1, true},
		{"hp htww xppe le xtoytrse", "we will meet at midnight", -11, false},
		{"hP htww Xppe le xtoytrsE", "wE will Meet at midnighT", -11, true},
		{"dEFghIJklm NOpqrSTu vwxyzabc", "aBCdeFGhij KLmnoPQr stuvwxyz", -3, true},
	}

	for _, test := range tests {
		result := shiftLetters([]byte(test.text), test.shift, test.matchCase)
		if string(result) != test.result {
			t.Errorf("ShiftLetters returned wrong result. Expected %s, got %s", test.result, string(result))
		}
	}
}
