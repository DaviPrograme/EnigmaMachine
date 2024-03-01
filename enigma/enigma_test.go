package enigma

import "testing"

const standardErrorMSG = "Expected Value: %v     Returned Value: %v"

func TestMoveLeftRunes(t *testing.T) {
	baseTest := "ABCD"

	expectedValue1 := "BCDA"
	case1 := []rune(baseTest)
	moveLeftRunes(case1)
	if string(case1) != expectedValue1 {
		t.Errorf(standardErrorMSG, expectedValue1, string(case1))
	}

	expectedValue2 := "CDAB"
	case2 := []rune(baseTest)
	moveLeftRunes(case2)
	moveLeftRunes(case2)
	if string(case2) != expectedValue2 {
		t.Errorf(standardErrorMSG, expectedValue2, string(case2))
	}

	expectedValue3 := "ABCD"
	case3 := []rune(baseTest)
	moveLeftRunes(case3)
	moveLeftRunes(case3)
	moveLeftRunes(case3)
	moveLeftRunes(case3)
	if string(case3) != expectedValue3 {
		t.Errorf(standardErrorMSG, expectedValue3, string(case3))
	}

}

func TestCountScramblers(t *testing.T) {
	var e Enigma

	expectedValue0 := 0
	if e.CountScramblers() != expectedValue0 {
		t.Errorf(standardErrorMSG, expectedValue0, e.CountScramblers())
	}

	expectedValue1 := 1
	e.scramblers = append(e.scramblers, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if e.CountScramblers() != expectedValue1 {
		t.Errorf(standardErrorMSG, expectedValue1, e.CountScramblers())
	}

	expectedValue2 := 2
	e.scramblers = append(e.scramblers, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if e.CountScramblers() != expectedValue2 {
		t.Errorf(standardErrorMSG, expectedValue2, e.CountScramblers())
	}
}

func TestAddScramblers(t *testing.T) {
	var e Enigma

	expectedCountScramblers0 := 0
	err := e.AddScrambler("BCDEFGHIJKLMNOPQRSTUVWXYZ")
	if len(e.scramblers) != expectedCountScramblers0 || err == nil {
		t.Error("It is not allowed to add scramblers with less than 26 characters as they must contain all characters from A to Z")
	}

	err = e.AddScrambler("AABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if len(e.scramblers) != expectedCountScramblers0 || err == nil {
		t.Error("Repeated characters are not allowed")
	}

	err = e.AddScrambler("ABCDEFGHIJKL-NOPQRSTUVWXYZ")
	if len(e.scramblers) != expectedCountScramblers0 || err == nil {
		t.Error("Scramblers with characters other than A to Z are not allowed")
	}

	expectedCountScramblers1 := 1
	err = e.AddScrambler("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if len(e.scramblers) != expectedCountScramblers1 || err != nil {
		t.Error("It is allowed to add scramblers with uppercase characters A to Z")
	}

	expectedCountScramblers2 := 2
	err = e.AddScrambler("abcdefghijklmnopqrstuvwxyz")
	if len(e.scramblers) != expectedCountScramblers2 || err != nil {
		t.Error("It is allowed to add scramblers with lowercase characters A to Z")
	}

	expectedCountScramblers3 := 3
	err = e.AddScrambler("ABCDEFGHIJKLMnopqrstuvwxyz")
	if len(e.scramblers) != expectedCountScramblers3 || err != nil {
		t.Error("It is allowed to add scramblers with uppercase and lowercase characters A to z")
	}

}

func TestClearScramblers(t *testing.T) {
	var e Enigma

	e.scramblers = append(e.scramblers, "42", "SAOPAULO")
	e.ClearScramblers()
	if len(e.scramblers) != 0 {
		t.Errorf(standardErrorMSG, 0, len(e.scramblers))
	}
}

func TestDeleteScrambler(t *testing.T) {
	var e Enigma

	e.scramblers = []string{"A", "B", "C", "D", "E"}
	expectedCountScramblers0 := 5
	e.DeleteScrambler(5)
	if len(e.scramblers) != expectedCountScramblers0 {
		t.Errorf(standardErrorMSG, expectedCountScramblers0, len(e.scramblers))
	}

	expectedCountScramblers1 := 4
	e.DeleteScrambler(0)
	if len(e.scramblers) != expectedCountScramblers1 {
		t.Errorf(standardErrorMSG, expectedCountScramblers1, len(e.scramblers))
	}
	containValue := false
	valueSought := "A"
	for _, scrambler := range e.scramblers {
		if scrambler == valueSought {
			containValue = true
			break
		}
	}
	if containValue {
		t.Errorf("Did not exclude expected value \"%s\"", valueSought)
	}

	expectedCountScramblers3 := 3
	e.DeleteScrambler(1)
	if len(e.scramblers) != expectedCountScramblers3 {
		t.Errorf(standardErrorMSG, expectedCountScramblers3, len(e.scramblers))
	}
	containValue = false
	valueSought = "C"
	for _, scrambler := range e.scramblers {
		if scrambler == valueSought {
			containValue = true
			break
		}
	}
	if containValue {
		t.Errorf("Did not exclude expected value \"%s\"", valueSought)
	}

	expectedCountScramblers2 := 2
	e.DeleteScrambler(2)
	if len(e.scramblers) != expectedCountScramblers2 {
		t.Errorf(standardErrorMSG, expectedCountScramblers2, len(e.scramblers))
	}
	containValue = false
	valueSought = "E"
	for _, scrambler := range e.scramblers {
		if scrambler == valueSought {
			containValue = true
			break
		}
	}
	if containValue {
		t.Errorf("Did not exclude expected value \"%s\"", valueSought)
	}
}

func TestInsertReflector(t *testing.T) {
	var e Enigma

	expectedReflector0 := ""
	err := e.InsertReflector("BCDEFGHIJKLMNOPQRSTUVWXYZ")
	if e.reflector != expectedReflector0 || err == nil {
		t.Error("It is not allowed to add reflector with less than 26 characters as they must contain all characters from A to Z")
	}

	err = e.InsertReflector("AABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if e.reflector != expectedReflector0 || err == nil {
		t.Error("Repeated characters are not allowed")
	}

	err = e.InsertReflector("ABCDEFGHIJKL-NOPQRSTUVWXYZ")
	if e.reflector != expectedReflector0 || err == nil {
		t.Error("Reflector with characters other than A to Z are not allowed")
	}

	e.reflector = ""
	expectedReflector1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	err = e.InsertReflector("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if e.reflector != expectedReflector1 || err != nil {
		t.Error("It is allowed to add reflector with uppercase characters A to Z")
	}

	e.reflector = ""
	err = e.InsertReflector("abcdefghijklmnopqrstuvwxyz")
	if e.reflector != expectedReflector1 || err != nil {
		t.Error("It is allowed to add reflector with lowercase characters A to Z")
	}

	e.reflector = ""
	err = e.InsertReflector("ABCDEFGHIJKLMnopqrstuvwxyz")
	if e.reflector != expectedReflector1 || err != nil {
		t.Error("It is allowed to add reflector with uppercase and lowercase characters A to z")
	}

}

func TestClearReflector(t *testing.T) {
	var e Enigma

	e.reflector = "42SAOPAULO"
	e.ClearReflector()
	if e.reflector != "" {
		t.Errorf(standardErrorMSG, 0, len(e.scramblers))
	}
}

func TestSetPlugBoard(t *testing.T) {
	var e Enigma

	expectedValue := 0
	e.SetPlugBoard([]string{"12", "34", "56"})
	if len(e.plugboard) != expectedValue {
		t.Errorf("No numbers allowed")
	}

	e.SetPlugBoard([]string{"AB", "C1", "DT"})
	if len(e.plugboard) != expectedValue {
		t.Errorf("No numbers allowed")
	}

	e.SetPlugBoard([]string{"*/", "%&", "@#"})
	if len(e.plugboard) != expectedValue {
		t.Errorf("Other characters that are not part of the alphabet are not allowed")
	}

	e.SetPlugBoard([]string{"AB", "CA", "DT"})
	if len(e.plugboard) != expectedValue {
		t.Errorf("Repeated characters are not allowed")
	}

	e.plugboard = map[string]string{}
	e.SetPlugBoard([]string{"AB", "CD", "EF"})
	if len(e.plugboard) == 0 {
		t.Errorf("Uppercase characters are allowed")
	}

	e.plugboard = map[string]string{}
	e.SetPlugBoard([]string{"ab", "cd", "ef"})
	if len(e.plugboard) == 0 {
		t.Errorf("Lowercase characters are allowed")
	}

	e.plugboard = map[string]string{}
	e.SetPlugBoard([]string{" AB", "C D", "EF "})
	if len(e.plugboard) == 0 {
		t.Errorf("Spaces are allowed")
	}

	e.plugboard = map[string]string{}
	e.SetPlugBoard([]string{"A-B", "C-D", "E-F "})
	if len(e.plugboard) == 0 {
		t.Errorf("Dashes are allowed to separate characters")
	}

	e.plugboard = map[string]string{}
	e.SetPlugBoard([]string{"A-B", "C-DZ", "E-F "})
	if len(e.plugboard) != 0 {
		t.Errorf("No more than two characters are allowed")
	}
}

func TestDecrypter(t *testing.T) {
	var e Enigma

	expectedValue := "ENIGMA"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	str, err := e.Decrypter("FFQJPX", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "ULTRA"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	e.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	str, err = e.Decrypter("ZYDNI", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "XVPURPLE"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	e.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	str, err = e.Decrypter("QHSGUWIG", []uint8{4})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "BLITZKRIEG"
	plugs := []string{"A-B", "S-Z", "U-Y", "G-H", "L-Q", "E-N"}
	e.AddScrambler("AJPCZWRLFBDKOTYUQGENHXMIVS")
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	e.AddScrambler("TAGBPCSDQEUFVNZHYIXJWLRKOM")
	e.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	e.SetPlugBoard(plugs)
	str, err = e.Decrypter("GYHRVFLRXY", []uint8{0, 4, 1})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	e.AddScrambler("AJPCZWRLFBDKOTYUQGENHXMIVS")
	_, err = e.Decrypter("GYHRVFLRXY", []uint8{0, 4})
	if err == nil {
		t.Errorf("The number of initial states of the scrambler must be equal to the number of scrambler")
	}

	e.ClearAll()

	expectedValue = "ENI-GMA"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	str, err = e.Decrypter("FFQ-JPX", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "ENI GMA"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	str, err = e.Decrypter("FFQ JPX", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}
}

func TestEncrypter(t *testing.T) {
	var e Enigma

	expectedValue := "FFQJPX"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	str, err := e.Encrypter("ENIGMA", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "ZYDNI"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	e.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	str, err = e.Decrypter("ULTRA", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "QHSGUWIG"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	e.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	str, err = e.Decrypter("XVPURPLE", []uint8{4})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "GYHRVFLRXY"
	plugs := []string{"A-B", "S-Z", "U-Y", "G-H", "L-Q", "E-N"}
	e.AddScrambler("AJPCZWRLFBDKOTYUQGENHXMIVS")
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	e.AddScrambler("TAGBPCSDQEUFVNZHYIXJWLRKOM")
	e.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	e.SetPlugBoard(plugs)
	str, err = e.Decrypter("BLITZKRIEG", []uint8{0, 4, 1})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	e.AddScrambler("AJPCZWRLFBDKOTYUQGENHXMIVS")
	_, err = e.Decrypter("GYHRVFLRXY", []uint8{0, 4})
	if err == nil {
		t.Errorf("The number of initial states of the scrambler must be equal to the number of scrambler")
	}

	e.ClearAll()

	expectedValue = "FFQ-JPX"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	str, err = e.Encrypter("ENI-GMA", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}

	e.ClearAll()

	expectedValue = "FFQ JPX"
	e.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	str, err = e.Encrypter("ENI GMA", []uint8{0})
	if str != expectedValue || err != nil {
		t.Errorf(standardErrorMSG, expectedValue, str)
	}
}
