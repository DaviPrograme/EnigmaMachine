package main

import (
	"EnigmaMachine/enigma"
	"fmt"
)

func main() {
	machine := &enigma.Enigma{}
	machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	str, _ := machine.Decrypter("FFQJPX", []uint8{0})
	fmt.Println(str)
	str, _ = machine.Encrypter("ENIGMA", []uint8{0})
	fmt.Println(str)

	machine.ClearAll()

	machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	machine.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	str, _ = machine.Decrypter("ZYDNI", []uint8{0})
	fmt.Println(str)
	str, _ = machine.Encrypter("ULTRA", []uint8{0})
	fmt.Println(str)

	machine.ClearAll()

	machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
	machine.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	str, _ = machine.Decrypter("QHSGUWIG", []uint8{4})
	fmt.Println(str)
	str, _ = machine.Decrypter("XVPURPLE", []uint8{4})
	fmt.Println(str)

	machine.ClearAll()

	plugs := []string{"A-B", "S-Z", "U-Y", "G-H", "L-Q", "E-N"}
	machine.AddScrambler("AJPCZWRLFBDKOTYUQGENHXMIVS") //2
	machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ") //1
	machine.AddScrambler("TAGBPCSDQEUFVNZHYIXJWLRKOM") //3
	machine.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	machine.SetPlugBoard(plugs)
	str, _ = machine.Decrypter("GYHRVFLRXY", []uint8{0, 4, 1})
	fmt.Println(str)
	str, _ = machine.Encrypter("BLITZKRIEG", []uint8{0, 4, 1})
	fmt.Println(str)

	machine.ClearAll()

}
