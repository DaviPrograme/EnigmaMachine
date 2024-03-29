# Enigma-Machine

## O que é este projeto?

Este projeto se trata de um pacote construido em [Golang](https://go.dev/) para emular o funcionamento da Máquiana Enigma, máquina esta que foi utilizada na segunda guerra mundial como forma de criptografar mensagens

## O que me motivou a criar este projeto?

Quem me conhece sabe que sou muito fã de jogos e sou apaixonado por programação, e amo quando vejo oportunidades de unir as duas coisas. Bem, encontrei esta oportunidade de unir as duas coisas em um jogo chamado [Cypher](https://store.steampowered.com/app/746710/Cypher/), este jogo consiste em vários desafios no campo da criptografia, conhecendo assim boa parte da sua origem, começando por técnicas como a [Esteganografia](https://pt.wikipedia.org/wiki/Esteganografia) que é a arte de ocultar mensagens dentro de outra mensagem, passando também pelas [Cifra de Cézar](https://pt.wikipedia.org/wiki/Cifra_de_C%C3%A9sar), [Cifra de Vigenère](https://pt.wikipedia.org/wiki/Cifra_de_Vigen%C3%A8re), Máquina Enigma (foco deste projeto), chegando até a [Criptografia Digital](https://www.usna.edu/Users/cs/wcbrown/courses/si110AY13S/lec/l27/lec.html). Bem, quando eu estava jogando e cheguei nos desafios da Máquina Enigma notei que nao sabia como ela funciona de fato (apesar de ser fã do filme [Jogo da Imitação](https://pt.wikipedia.org/wiki/O_Jogo_da_Imita%C3%A7%C3%A3o) rsrs), e por este motivo iniciei os meu estudos com o intuito de saciar esta curiosidade, e em um dado momento pensei: "mais do que apenas resolver os desafios propostos pelo jogo por que não crio a minha versão da Máquina Enigma?", e bem, foi basicamente esta pergunta que me trouxe até aqui rs.

## Qual a importância de aprender sobre a Máquina Enigma?

Bem, caso seja uma pessoa curiosa assim como eu o simples fato de saciar esta sede ja seria um bom motivo, mas consigo pensar em mais alguns motivos para aprender sobre a Máquina Enigma, que são os seguintes:

* Na sociedade, toda criação feita hoje é baseada no aprendizado de ontem, da mesma forma que entender sobre a esteganografia, Cifras de Cezar e Vigenère ajudam a entender o por quê da criação  da Máquina Enigma, entender sobre a Máquina Enigma ajuda entender sobre tecnicas de criptografia mais recentes como por exemplo [DES](https://pt.wikipedia.org/wiki/Data_Encryption_Standard), [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard), [RSA](https://pt.wikipedia.org/wiki/RSA_(sistema_criptogr%C3%A1fico)),  [SHA-256](https://pt.wikipedia.org/wiki/SHA-2);
* A Máquina Enigma foi um importante fator durante a Segunda guerra Mundial, entender sobre a máquina ajuda entender mais o contexto de algumas coisas sobre este tema;
* Antigamente quando existia um conflito entre nações o embate acontecia principalmente no âmbito físico, nos dias de hoje a guerra acontece principalmente no âmbito das informações/dados, saber proteger os dados virou assunto de segurança nacional. Entender sobre a Máquina Enigma (e por consequência, devido aos pontos anteriores sobre a criptografia em si) ajuda entender como esses embates acontecem no âmbito das nações e corporações nos tempos modernos.

## Resumo da História da Maquina Enigma

![Maquina-Enigma-Alemanha-nazismo](https://github.com/DaviPrograme/Enigma-Machine/assets/56012877/f5eb23db-0097-49bf-a9db-66f9f85e90e9)


A máquina Enigma é uma das mais famosas e intrigantes máquinas criptográficas da história. Desenvolvida na Alemanha durante a década de 1920, foi projetada originalmente para uso comercial e governamental, mas eventualmente se tornou uma peça central da criptografia militar alemã durante a Segunda Guerra Mundial.

A Enigma foi inventada pelo engenheiro alemão Arthur Scherbius, que patenteou sua primeira versão em 1918. No entanto, foi somente após sua aquisição pela empresa alemã Chiffriermaschinen Aktiengesellschaft (Chiffriermaschinen AG) que a máquina foi refinada e comercializada com o nome "Enigma", em 1923.

A máquina Enigma consistia em um teclado, um conjunto de rotores (geralmente três a cinco) que podiam ser trocados e configurados de várias maneiras, um painel de lâmpadas que exibiam as letras cifradas e um refletor que direcionava o circuito de volta através dos rotores, aumentando a complexidade da criptografia. Cada vez que uma tecla era pressionada, os rotores giravam, alterando a substituição da letra.

Durante a Segunda Guerra Mundial, as forças armadas alemãs utilizaram a máquina Enigma para cifrar suas comunicações, confiando em sua suposta inviolabilidade para manter seus segredos militares seguros. No entanto, a criptografia da Enigma foi quebrada pelos esforços hercúleos dos criptoanalistas aliados, especialmente pela equipe em Bletchley Park, liderada por Alan Turing.

A quebra da criptografia Enigma foi uma realização monumental, crucial para o esforço de guerra dos Aliados. Isso foi alcançado principalmente por meio do desenvolvimento de máquinas chamadas "bombas", que simulavam milhares de configurações possíveis da máquina Enigma até encontrar a correta para uma mensagem específica. Além disso, contribuíram para o sucesso os esforços de inteligência que forneceram pistas sobre as configurações diárias da Enigma, bem como a análise estatística dos padrões nas mensagens cifradas.

Em última análise, a quebra da Enigma permitiu que os Aliados interceptassem e decifrassem uma quantidade significativa de comunicações inimigas, fornecendo uma vantagem crucial que contribuiu para a vitória final na Segunda Guerra Mundial.

## Principais Partes da Máquina Enigma

* **Teclado:** O teclado da máquina Enigma consiste em um conjunto de teclas, cada uma representando uma letra do alfabeto. Quando uma tecla é pressionada, um circuito elétrico é ativado, enviando um sinal elétrico através da máquina;

* **Placa de Conexões:** O plugboard é uma parte da Enigma que permite a troca de letras antes e depois de passarem pelos rotores. Ele consiste em um conjunto de cabos que são conectados entre pares de letras. Por exemplo, se o "A" estiver conectado ao "F", sempre que o "A" for pressionado no teclado, ele será substituído pelo "F" antes de entrar nos rotores;

* **Rotores:** Os rotores são o coração da máquina Enigma. Cada rotor é um disco metálico com contatos elétricos em cada lado, representando as letras do alfabeto. Os rotores são colocados em uma ordem específica na máquina e podem ser configurados de várias maneiras. Cada vez que uma tecla é pressionada, os rotores giram, alterando a substituição da letra. Isso aumenta drasticamente a complexidade da criptografia;

* **Refletor**: O refletor é uma parte crucial da máquina Enigma que contribui para a complexidade da criptografia. Localizado após os rotores, o refletor serve para direcionar o circuito elétrico de volta através dos rotores após o sinal passar por eles. Ele consiste em um conjunto de contatos elétricos que refletem a corrente elétrica de volta através dos rotores, adicionando uma camada adicional de confusão à cifra.
  
* **Painel de Lâmpadas:** O lampboard é onde as letras cifradas são exibidas. Após passarem pelos rotores e pelo plugboard, as letras são iluminadas no lampboard, revelando a letra cifrada correspondente à letra digitada no teclado. Isso permite que o operador leia a mensagem cifrada e a transcreva para entrega ou armazenamento.

## Como a Máquina Enigma Funciona?

![funcionamento_maquina_enigma](https://github.com/DaviPrograme/Enigma-Machine/assets/56012877/54f1bd7a-2762-4de4-be93-866cd2713015)

Listarei agora o passo a passo resumido de como a Máquina Enigma funciona:

* **Passo 1:** Todo o processo da Máquina Enigma começa no teclado, quando uma tecla é pressionada é liberado um sinal que corresponde a letra, esse sinal passa primeiro pela placa de conexões;

* **Passo 2:** Recebendo o sinal do passo anterior, a placa de conexões faz uma validação para saber como proceder, por exemplo, se o sinal enviado pelo teclado foi a letra "A" e na placa de conexões a letra "A" esta plugada com a letra "F", neste caso o sinal vai ser alterado para representar a letra "F", caso contrario o sinal será enviado como letra "A", em ambos os caso o sinal vai ser enviado para o próximo passo, que são os rotores;

* **Passo 3:** O sinal repassado pelo passo anterior chega nos rotores, aqui é o "coração" de todo o processo pois aqui é onde acontece a maior parte dos embaralhamentos. O rotor é uma peça cilindrica onde no seu entorno tem a representação dos 26 caracteres de "A" a "Z", toda vez que um sinal chega ao rotor mais a direita ele muda sua posição, então por exemplo, se antes do sinal o rotor estava na letra "A" ele passa para letra "B", recebe o sinal para o proximo rotor como sendo a letra "B" e os proximos rotores repetem o mesmo processo com excessão de que, o próximo rotor só vai mudar a sua posição quando o rotor da sua direita completar um giro completo, passando assim por todas as letras, e dessa forma o sinal e trasmitido para os proximos rotores onde cada um realizara a alteração do caracter até chegar no proximo passo, que é o refletor;

* **Passo 4:** O sinal do passo anterior chega agora no refletor, que nada mais é que mais um rotor, mas com um pequena diferença, ao invés de ele passar o dado para frente ele faz a alteração do sinal, alterando assim mais uma vez a representação do caracter em questão e retorna esse dado pelo caminho que ele veio. Dessa forma, ele vai retornar por todos os rotres tendo assim mais uma rodada de alterações até passar pelo ultimo rotor mais a esquerda indo assim para o proximo passo que é a placa de conexões novamente;

* **Passo 5:** Neste passo, ele repetira o processo do passo 2 avaliando qual foi a representação de caracter que os rotores enviaram, caso a letra em questão esteja plugada em outra letra o caracter será alterado mais uma vez, caso contrário não sofrerá alterações e sera enviado para o último passo, que é o painel de lâmpadas;

* **Passo 6:** No painel de lampadas ele recebbe o sinal do passo anterior e acende uma lampada que representa o caracter cifrado, e dessa forma o processo é encerrado até que a proxima tecla seja pressionada reiniciando o ciclo.


Caso tenha ficado dúvidas, segue um [video](https://www.youtube.com/watch?v=ybkkiGtJmkM&t=709s&ab_channel=JaredOwen) que me ajudou muito a entender o funcionamento interno da máquina.


## Sobre o Pacote Enigma

A minha ideia foi trasnportar os principais elementos da máquina fisica para o código, que são:

* **Scrambler:** que é o rotor, responsável pelo embaralhamento;
* **Reflector:** que é o refletor, o elemento que retorna o sinal para os rotores com mais uma alteração;
* **Plugboard:** que é o painel onde são realizadas alterações do sinal antes e depois de passar pelos rotores.

Para utilizar o projeto basta baixar este repositorio e fazer o import como na linha abaixo:

```bash
import "EnigmaMachine/enigma"
```

### Como funciona o pacote?

A primeira coisa a se fazer é criar uma instancia do objeto enigma que pode ser realizado da seguinte forma:

```bash
machine := &enigma.Enigma{}
```

Após ter instaciado o projeto, é necessario configurar scramblers (obrigatorio), reflector (caso seja necessário) e o plugboard (caso seja necessário), vamos começar pelo scrambler. Para adicionar um scrambler voce pode fazer da seguinte forma:

```bash
err := machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ")
```

esta função recebe como parãmetro uma string que deve conter apenas caracteres de "A" a "Z", esses caracteres podem ser maiusculas ou minusculas mas tem que ter 26 letras representando todos os caracteres do alfabeto sem repetição, caso ocorra algum erro é possivel verificar no retorno do metodo. É possivel adicionar mais de um scrambler repetindo o mesmo processo com uma string diferente, dessa forma aumenta a complexidade do embaralhamento

Além disso, caso seja necessario adicionar o reflector, isso pode ser realizado da seguinte forma:

```bash
err := machine.InsertReflector("UWYGADFPVZBECKMTHXSLRINQOJ")
```

O metodo **InsertReflector** respeita as mesmas regras como AddScrambbler, com a única diferença sendo que nesse caso não é possivel adicionar mais de um refletor, caso você chame a função novamente com outra string ela realizara a substituição do valor anterior.

Caso seja necessário também, é possível inserir o plugboard da seguinte forma:

```bash
err := machine.SetPlugBoard([]string{"A-B", "S-Z", "U-Y", "G-H", "L-Q", "E-N"})
```

O método **SetPlugBoard**  ele recebe um slice de strings, as representações dos caracteres podem ser representados também sem o traço, funcionará da mesma forma.

Com tudo isso configurado podemos encriptar uma mensagem da seguinte forma:

```bash
str, err := machine.Encrypter("ENIGMA", []uint8{0})
```

A função **Encrypter** recebe dois parâmetros, primeiro a mensagem a ser criptografada, depois um slice de uint8 que representa a posição inicial de cada scrambler, então por exemplo, colocar o valor como 0 representa que ele vai começar na letra "A", 1 representa a letra "B" e assim sucessivamente. Caso voce insira mais de um scrambler ele tera que receber a mesma quantidade de representações de posiciões iniciais conforma o exemplo abaixo:

```bash
machine.AddScrambler("AJPCZWRLFBDKOTYUQGENHXMIVS")
machine.AddScrambler("UWYGADFPVZBECKMTHXSLRINQOJ") 
machine.AddScrambler("TAGBPCSDQEUFVNZHYIXJWLRKOM") 
str, err = machine.Encrypter("BLITZKRIEG", []uint8{0, 4, 1})
```

No caso acima eu passei as representações de 0 "A", 4 "E", 1 "B", caso eu tivesse passando uma quantidade maior ou menor  de estados iniciais ele geraria um erro.

Para decriptar algo pode ser feito da seguinte forma:

```bash
str, err :=machine.Decrypter("GYHRVFLRXY", []uint8{0, 4, 1})
```

A função **Decrypter** segue a mesma lógica do metodo **Encrypter** com a unica diferrença que no  texto ele espera uma mensagem cifrada, de resto a lógica é a mesma :)
  





