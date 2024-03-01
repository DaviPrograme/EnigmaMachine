# Enigma-Machine

## What is this project?

This project is a package built in  [Golang](https://go.dev/) to emulate the operation of the Enigma Machine, a machine that was used in World War II to encrypt messages.

## What motivated me to create this project?

Anyone who knows me knows that I am a big fan of games and passionate about programming, and I love when I see opportunities to combine the two. Well, I found this opportunity to combine both in a game called [Cypher](https://store.steampowered.com/app/746710/Cypher/), eThis game consists of several challenges in the field of cryptography, thus getting to know a good part of its origin, starting with techniques such as [Steganography](https://pt.wikipedia.org/wiki/Esteganografia), which is the art of hiding messages within another message, also going through the [Caesar Cipher](https://pt.wikipedia.org/wiki/Cifra_de_C%C3%A9sar), [Vigenère Cipher](https://pt.wikipedia.org/wiki/Cifra_de_Vigen%C3%A8re), Enigma Machine (focus of this project), reaching [Digital Cryptography.](https://www.usna.edu/Users/cs/wcbrown/courses/si110AY13S/lec/l27/lec.html). BWell, when I was playing and reached the challenges of the Enigma Machine, I realized that I didn't actually know how it works (despite being a fan of the movie [The Imitation Game](https://pt.wikipedia.org/wiki/O_Jogo_da_Imita%C3%A7%C3%A3o) lol), eand for this reason I started my studies with the intention of satisfying this curiosity, and at some point I thought: "more than just solving the challenges proposed by the game why not create my own version of the Enigma Machine?", and well, it was basically this question that brought me here.

## What is the importance of learning about the Enigma Machine?

Well, if you are a curious person like me, just satisfying this thirst would be a good reason, but I can think of a few more reasons to learn about the Enigma Machine, which are as follows:

* n society, every creation made today is based on yesterday's learning, just as understanding steganography, Caesar and Vigenère Ciphers help understand the creation of the Enigma Machine, understanding about the Enigma Machine helps understand more about recent cryptography techniques such as [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard), [RSA](https://pt.wikipedia.org/wiki/RSA_(sistema_criptogr%C3%A1fico)),  [SHA-256](https://pt.wikipedia.org/wiki/SHA-2);
* The Enigma Machine was an important factor during World War II, understanding about the machine helps to understand more about the context of some things on this subject;
* In the past, when there was a conflict between nations, the confrontation mainly occurred in the physical realm, today war mainly occurs in the realm of information/data, knowing how to protect data has become a matter of national security. Understanding about the Enigma Machine (and consequently, due to the previous points about cryptography itself) helps understand how these conflicts occur in the realm of nations and corporations in modern times.

## Summary of the History of the Enigma Machine

![Maquina-Enigma-Alemanha-nazismo](https://github.com/DaviPrograme/Enigma-Machine/assets/56012877/f5eb23db-0097-49bf-a9db-66f9f85e90e9)


The Enigma Machine is one of the most famous and intriguing cryptographic machines in history. Developed in Germany during the 1920s, it was originally designed for commercial and governmental use, but eventually became a central piece of German military cryptography during World War II.

The Enigma was invented by the German engineer Arthur Scherbius, who patented its first version in 1918. However, it was only after its acquisition by the German company Chiffriermaschinen Aktiengesellschaft (Chiffriermaschinen AG) that the machine was refined and marketed under the name "Enigma" in 1923.

The Enigma machine consisted of a keyboard, a set of scramblers (usually three to five) that could be exchanged and configured in various ways, a lampboard that displayed the encrypted letters, and a reflector that directed the circuit back through the scramblers, increasing the complexity of the cryptography. Each time a key was pressed, the scramblers rotated, altering the letter substitution.

During World War II, the German armed forces used the Enigma machine to encrypt their communications, relying on its supposed inviolability to keep their military secrets safe. However, the encryption of the Enigma was broken by the herculean efforts of the Allied cryptanalysts, especially by the team at Bletchley Park, led by Alan Turing.

The breaking of the Enigma encryption was a monumental achievement, crucial to the Allied war effort. This was mainly achieved through the development of machines called "bombs", which simulated thousands of possible settings of the Enigma machine until finding the correct one for a specific message. In addition, intelligence efforts that provided clues about the daily settings of the Enigma, as well as statistical analysis of patterns in encrypted messages, contributed to the success.

Ultimately, the breaking of the Enigma allowed the Allies to intercept and decipher a significant amount of enemy communications, providing a crucial advantage that contributed to the final victory in World War II.

## Main Parts of the Enigma Machine

* **Keyboard:** The keyboard of the Enigma machine consists of a set of keys, each representing a letter of the alphabet. When a key is pressed, an electrical circuit is activated, sending an electrical signal through the machine;

* **Plugboard:** The plugboard is a part of the Enigma machine that allows the exchange of letters before and after passing through the scramblers. It consists of a set of cables that are connected between pairs of letters. For example, if "A" is connected to "F" on the plugboard, whenever "A" is pressed on the keyboard, it will be replaced by "F" before entering the scramblers;

* **Scramblers:** The scramblers are the heart of the Enigma machine. Each scrambler is a cylindrical disk with electrical contacts on each side, representing the letters of the alphabet. The scramblers are placed in a specific order in the machine and can be configured in various ways. Each time a key is pressed, the scramblers rotate, altering the letter substitution. This greatly increases the complexity of the cryptography;

* **Reflector**: The reflector is a crucial part of the Enigma machine that contributes to the complexity of the cryptography. Located after the scramblers, the reflector serves to direct the electrical circuit back through the scramblers after the signal passes through them. It consists of a set of electrical contacts that reflect the electric current back through the scramblers, adding an additional layer of confusion to the cipher;
  
* **Lampboard:** The lampboard is where the encrypted letters are displayed. After passing through the scramblers and the plugboard, the letters are illuminated on the lampboard, revealing the encrypted letter corresponding to the letter typed on the keyboard. This allows the operator to read the encrypted message and transcribe it for delivery or storage.

## How does the Enigma Machine work?

![funcionamento_maquina_enigma](https://github.com/DaviPrograme/Enigma-Machine/assets/56012877/54f1bd7a-2762-4de4-be93-866cd2713015)

I will now list a summarized step-by-step of how the Enigma Machine works:

* **Step 1:** The entire process of the Enigma Machine starts at the keyboard, when a key is pressed a signal corresponding to the letter is released, this signal first passes through the plugboard;

* **Step 2:** Receiving the signal from the previous step, the plugboard performs a validation to know how to proceed, for example, if the signal sent by the keyboard was the letter "A" and on the plugboard the letter "A" is plugged with the letter "F", in this case the signal will be changed to represent the letter "F", otherwise the signal will be sent as the letter "A" itself, in both cases the signal will be sent to the next step, which are the scramblers;

* **Step 3:** The signal passed by the previous step reaches the scramblers, here is the "heart" of the whole process because here is where most of the shuffling takes place. The scrambler is a cylindrical piece where around its perimeter has the representation of the 26 characters from "A" to "Z", every time a signal reaches the rightmost scrambler it changes its position, so for example, if before the signal the scrambler was in the letter "A" it moves to letter "B", receives the signal to the next scrambler as the letter "B" and the next scramblers repeat the same process except that, the next scrambler will only change its position when the rightmost scrambler completes a full rotation, thus passing through all the letters, and in this way the signal is transmitted to the next scramblers where each one will perform the alteration of the character until reaching the next step, which is the reflector;

* **Step 4:** The signal from the previous step now reaches the reflector, which is nothing more than another scrambler, but with a small difference, instead of passing the data forward it makes the alteration of the signal, thus changing the representation of the character once again and returns this data through the path it came from. In this way, it will return through all the scramblers having thus one more round of alterations until passing through the leftmost scrambler going to the next step, which is the plugboard again;

* **Step 5:**  In this step, it will repeat the process from step 2 evaluating which was the character representation that the scramblers sent, if the letter in question is plugged into another letter the character will be altered once again, otherwise it will not undergo alterations and will be sent to the last step, which is the lampboard;

* **Step 6:** In the lampboard, it receives the signal from the previous step and lights up a bulb representing the encrypted character, and in this way the process is terminated until the next key is pressed restarting the cycle.


## About the Enigma Package

My idea was to transport the main elements of the physical machine to the code, which are:

* **Scrambler:** which is the rotor, responsible for shuffling;
* **Reflector:** which is the reflector, the element that returns the signal to the scramblers with one more alteration;
* **Plugboard:**  which is the panel where alterations are made to the signal before and after passing through the scramblers.

To use the project just download this repository and import it as in the line below:

```bash
import "EnigmaMachine/enigma"
```

### How does the package work?

The first thing to do is to create an instance of the enigma object which can be done as follows:

```bash
machine := &enigma.Enigma{}
```

After instantiating the project, it is necessary to configure scramblers (mandatory), reflector (if necessary), and plugboard (if necessary), let's start with the scrambler. To add a scrambler you can do as follows:

```bash
err := machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
```

eThis function takes as a parameter a string that must contain only characters from "A" to "Z", these characters can be uppercase or lowercase but must have 26 letters representing all the characters of the alphabet without repetition, in case any error occurs it is possible to check the return of the method. It is possible to add more than one scrambler by repeating the same process with a different string, in this way the complexity of the shuffling is increased.

Furthermore, if necessary, it is necessary to add the reflector, this can be done as follows:

```bash
err := machine.InsertReflector("UWYGADFPVZBECKMTHXSLRINQOJ")
```

The **InsertReflector** method follows the same rules as AddScrambbler, cwith the only difference being that in this case it is not possible to add more than one reflector, if you call the function again with another string it will replace the previous value.

If necessary, it is also possible to insert the plugboard as follows:

```bash
err := machine.SetPlugBoard([]string{"A-B", "S-Z", "U-Y", "G-H", "L-Q", "E-N"})
```

The **SetPlugBoard**   method takes a slice of strings as input, the representations of the characters can also be represented without the dash, it will work in the same way.

With all this configured, we can encrypt a message as follows:

```bash
str, err := machine.Encrypter("ENIGMA", []uint8{0})
```

The **Encrypter**  function takes two parameters, first the message to be encrypted, then a slice of uint8 representing the initial position of each scrambler, so for example, putting the value as 0 represents that it will start at the letter "A", 1 represents the letter "B" and so on. If you insert more than one scrambler it will have to receive the same number of initial positions representations as in the example below:

```bash
plugs := []string{"A-B", "S-Z", "U-Y", "G-H", "L-Q", "E-N"}
machine.AddScrambler("AJPCZWRLFBDKOTYUQGENHXMIVS")
machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ") 
machine.AddScrambler("TAGBPCSDQEUFVNZHYIXJWLRKOM") 
machine.InsertReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
str, err = machine.Encrypter("BLITZKRIEG", []uint8{0, 4, 1})
```

In the above case, I passed the representations of 0 "A", 4 "E", 1 "B", if I had passed a greater or smaller amount of initial states it would generate an error.

To decrypt something it can be done as follows:

```bash
str, err :=machine.Decrypter("GYHRVFLRXY", []uint8{0, 4, 1})
```

The **Decrypter** function follows the same logic as the **Encrypter** c method with the only difference that in the text it expects an encrypted message, otherwise the logic is the same :)




  





