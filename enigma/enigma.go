package enigma

import (
	"errors"
	"fmt"
	"strings"
)

type Enigma struct {
	scramblers []string
	ringGroove string
	reflector  string
	plugboard  map[string]string
}

func moveLeftRunes(runes []rune) {
	runeStartToEnd := runes[0]
	lenRunes := len(runes)

	for index, _ := range runes {
		if index == lenRunes-1 {
			runes[index] = runeStartToEnd
			continue
		}
		runes[index] = runes[index+1]
	}
}

func (e *Enigma) setRingGroove(ringGroove string) error {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ringGroove = strings.ReplaceAll(strings.ToUpper(ringGroove), " ", "")
	if ringGroove == "" {
		return errors.New("must contain characters from A to Z")
	}
	cpyRingGroove := strings.Clone(ringGroove)

	for _, char := range alphabet {
		cpyRingGroove = strings.ReplaceAll(cpyRingGroove, string(char), "")
	}
	if len(cpyRingGroove) != 0 {
		return errors.New("must only contain characters A to Z")
	}
	e.ringGroove = ringGroove
	return nil
}

func (e *Enigma) clearRingGroove() {
	e.ringGroove = ""
}

func (e Enigma) CountScramblers() int {
	return len(e.scramblers)
}

func (e *Enigma) AddScrambler(scrambler string) error {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scrambler = strings.ToUpper(scrambler)
	cpyScrambler := strings.Clone(scrambler)

	if len(scrambler) != 26 {
		return errors.New("[SCRAMBLER] The size must be 26 to contain all characters from A to Z")
	}
	for _, char := range alphabet {
		cpyScrambler = strings.Replace(cpyScrambler, string(char), "", 1)
	}
	if len(cpyScrambler) > 0 {
		return errors.New("[SCRAMBLER] There are repeated or invalid characters, the scrambler must only contain characters from A to Z. \nList of repeated or invalid characters: [" + cpyScrambler + "].")
	}
	e.scramblers = append(e.scramblers, scrambler)
	return nil
}

func (e *Enigma) ClearScramblers() {
	e.scramblers = []string{}
}

func (e *Enigma) DeleteScrambler(indexDelete uint8) {
	if int(indexDelete) < len(e.scramblers) {
		if indexDelete == 0 {
			e.scramblers = e.scramblers[1:len(e.scramblers)]
		} else {
			start := e.scramblers[:indexDelete]
			cpy := start
			end := e.scramblers[indexDelete+1 : len(e.scramblers)]
			cpy = append(cpy, end...)
			e.scramblers = cpy
		}
	}
}

func (e Enigma) ShowScramblers() {
	fmt.Printf("amount scramblers: %d\n", len(e.scramblers))
	fmt.Println("scramblers: [")
	for _, scrambler := range e.scramblers {
		fmt.Println(scrambler)
	}
	fmt.Println("]")
}

func (e *Enigma) InsertReflector(reflector string) error {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	reflector = strings.ToUpper(reflector)
	cpyReflector := strings.Clone(reflector)

	if len(reflector) != 26 {
		return errors.New("[REFLECTOR] The size must be 26 to contain all characters from A to Z")
	}
	for _, char := range alphabet {
		cpyReflector = strings.Replace(cpyReflector, string(char), "", 1)
	}
	if len(cpyReflector) > 0 {
		return errors.New("[REFLECTOR] There are repeated or invalid characters, the scrambler must only contain characters from A to Z. \nList of repeated or invalid characters: [" + cpyReflector + "].")
	}
	e.reflector = reflector
	return nil
}

func (e *Enigma) ClearReflector() {
	e.reflector = ""
}

func (e *Enigma) SetPlugBoard(plugs []string) error {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	plugboard := make(map[string]string)
	var letters []string

	for index, plug := range plugs {
		plug = strings.ReplaceAll(strings.ToUpper(plug), " ", "")
		if strings.Contains(plug, "-") {
			letters = strings.Split(plug, "-")
		} else {
			letters = strings.Split(plug, "")
		}
		if len(letters) != 2 || len(letters[0]) != 1 || len(letters[1]) != 1 {
			return fmt.Errorf("configuration error found in plug index [%d] plug \"%s\"", index, plug)
		}
		if plugboard[letters[0]] != "" {
			return fmt.Errorf("duplicate key index [%d] in the letter \"%s\"", index, letters[0])
		}
		if plugboard[letters[1]] != "" {
			return fmt.Errorf("duplicate key index [%d] in the letter \"%s\"", index, letters[1])
		}
		if !strings.Contains(alphabet, letters[0]) || !strings.Contains(alphabet, letters[1]) {
			return fmt.Errorf("only characters A to Z are allowed")
		}
		plugboard[letters[0]] = letters[1]
		plugboard[letters[1]] = letters[0]
	}
	e.plugboard = plugboard
	return nil
}

func (e *Enigma) ClearPlugBoard() {
	e.plugboard = map[string]string{}
}

func (e *Enigma) ClearAll() {
	e.ClearScramblers()
	e.ClearReflector()
	e.ClearPlugBoard()
}

func (e Enigma) ShowPlugBoard() {
	values := ""

	fmt.Println("Plugboard: {")
	for key, value := range e.plugboard {
		if strings.Contains(values, key) {
			continue
		}
		fmt.Printf("%s <-> %s,\n", key, value)
		values = values + value
	}
	fmt.Println("}")
}

func (e Enigma) encrypterChar(indexScrambler uint, alphabets, scramblers [][]rune, indexSearch uint8) uint8 {
	if int(indexScrambler) < len(e.scramblers) {
		letter := alphabets[indexScrambler][indexSearch]
		index := uint8(strings.Index(string(scramblers[indexScrambler]), string(letter)))
		index = e.encrypterChar(indexScrambler+1, alphabets, scramblers, index)
		if e.reflector != "" {
			letter = scramblers[indexScrambler][index]
			index = uint8(strings.Index(string(alphabets[indexScrambler]), string(letter)))
			return index
		}
		return index
	}
	if e.reflector == "" {
		return indexSearch
	}
	alphabetInit := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letter := alphabetInit[indexSearch]
	index := strings.Index(e.reflector, string(letter))
	return uint8(index)
}

func (e Enigma) decrypterChar(alphabets, scramblers [][]rune, char rune) uint8 {
	alphabetString := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var index int

	if e.reflector == "" {
		index = strings.Index(alphabetString, string(char))
		for loops := len(scramblers) - 1; loops >= 0; loops-- {
			letter := scramblers[loops][index]
			index = strings.Index(string(alphabets[loops]), string(letter))
		}
		return uint8(index)
	}
	return uint8(e.encrypterChar(0, alphabets, scramblers, uint8(strings.Index(alphabetString, string(char)))))
}

func (e Enigma) swapPlugBoard(charSearch rune) rune {
	letter := string(charSearch)
	if e.plugboard[letter] != "" {
		return []rune(e.plugboard[letter])[0]
	}
	return charSearch
}

func (e Enigma) scramblerMadeCompleteTurn(indexScrambler int, spinScramblersBefore, spinScramblersAfter []uint8) bool {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if indexScrambler == -1 {
		return true
	}
	if len(e.ringGroove) > indexScrambler {
		charRingGroove := strings.Split(e.ringGroove, "")[indexScrambler]
		charNow := strings.Split(alphabet, "")[spinScramblersAfter[indexScrambler]]
		indexBeforeCharRingGroove := (strings.Index(alphabet, charRingGroove) - 1)
		if indexBeforeCharRingGroove == -1 {
			indexBeforeCharRingGroove = 25
		}
		if charRingGroove == charNow && spinScramblersBefore[indexScrambler] == uint8(indexBeforeCharRingGroove) {
			return true
		}
		return false
	}
	if spinScramblersBefore[indexScrambler] == 25 && spinScramblersAfter[indexScrambler] == 0 {
		return true
	}
	return false
}

func (e Enigma) Decrypter(msg string, initialStateScramblers []uint8) (decryptedMsg string, err error) {
	countScramblers := len(e.scramblers)
	countInitialStateScramblers := len(initialStateScramblers)
	if countScramblers != countInitialStateScramblers {
		err = fmt.Errorf("difference between number of scramblers (%d) and initialStateScramblers (%d)", countScramblers, countInitialStateScramblers)
		return decryptedMsg, err
	}
	msg = strings.ToUpper(msg)
	alphabets := [][]rune{}
	scramblers := [][]rune{}
	spinScramblers := []uint8{}
	spinScramblers = append(spinScramblers, initialStateScramblers...)
	alphabetStringInit := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetRuneInit := []rune(alphabetStringInit)
	for count := 0; count < countScramblers; count++ {
		scramblers = append(scramblers, []rune(e.scramblers[count]))
		alphabets = append(alphabets, []rune(strings.Clone(alphabetStringInit)))
	}
	for index, _ := range scramblers {
		for count := initialStateScramblers[index] % 26; count > 0; count-- {
			moveLeftRunes(alphabets[index])
			moveLeftRunes(scramblers[index])
		}
	}
	for _, char := range msg {
		if !strings.Contains(alphabetStringInit, string(char)) {
			decryptedMsg = decryptedMsg + string(char)
			continue
		}
		spinScramblersBefore := []uint8{}
		spinScramblersBefore = append(spinScramblersBefore, spinScramblers...)
		for indexScramblers, _ := range scramblers {
			if e.scramblerMadeCompleteTurn(indexScramblers-1, spinScramblersBefore, spinScramblers) {
				moveLeftRunes(alphabets[indexScramblers])
				moveLeftRunes(scramblers[indexScramblers])
				spinScramblers[indexScramblers] = (spinScramblers[indexScramblers] + 1) % 26
			}
		}
		index := e.decrypterChar(alphabets, scramblers, e.swapPlugBoard(char))
		encryptedChar := alphabetRuneInit[index]
		decryptedMsg = decryptedMsg + string(e.swapPlugBoard(encryptedChar))
	}
	return decryptedMsg, nil
}

func (e Enigma) Encrypter(msg string, initialStateScramblers []uint8) (encryptedMsg string, err error) {
	countScramblers := len(e.scramblers)
	countInitialStateScramblers := len(initialStateScramblers)
	if countScramblers != countInitialStateScramblers {
		err = fmt.Errorf("difference between number of scramblers (%d) and initialStateScramblers (%d)", countScramblers, countInitialStateScramblers)
		return encryptedMsg, err
	}
	msg = strings.ToUpper(msg)
	alphabets := [][]rune{}
	scramblers := [][]rune{}
	spinScramblers := []uint8{}
	spinScramblers = append(spinScramblers, initialStateScramblers...)
	alphabetStringInit := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetRuneInit := []rune(alphabetStringInit)
	for count := 0; count < countScramblers; count++ {
		scramblers = append(scramblers, []rune(e.scramblers[count]))
		alphabets = append(alphabets, []rune(strings.Clone(alphabetStringInit)))
	}
	for index, _ := range scramblers {
		for count := initialStateScramblers[index] % 26; count > 0; count-- {
			moveLeftRunes(alphabets[index])
			moveLeftRunes(scramblers[index])
		}
	}
	for _, char := range msg {
		if !strings.Contains(alphabetStringInit, string(char)) {
			encryptedMsg = encryptedMsg + string(char)
			continue
		}
		spinScramblersBefore := []uint8{}
		spinScramblersBefore = append(spinScramblersBefore, spinScramblers...)
		for indexScramblers, _ := range scramblers {
			if e.scramblerMadeCompleteTurn(indexScramblers-1, spinScramblersBefore, spinScramblers) {
				moveLeftRunes(alphabets[indexScramblers])
				moveLeftRunes(scramblers[indexScramblers])
				spinScramblers[indexScramblers] = (spinScramblers[indexScramblers] + 1) % 26
			}
		}
		index := e.encrypterChar(0, alphabets, scramblers, uint8(strings.Index(alphabetStringInit, string(e.swapPlugBoard(char)))))
		encryptedChar := alphabetRuneInit[index]
		encryptedMsg = encryptedMsg + string(e.swapPlugBoard(encryptedChar))
	}
	return encryptedMsg, err
}
