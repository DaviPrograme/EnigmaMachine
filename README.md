# Enigma-Machine

## O que é este projeto?

Este projeto se trata de um pacote construido em Golang para emular o funcionamento da Máquiana Enigma, máquina esta que foi utilizada na segunda guerra mundial como forma de criptografar mensagens

## O que me motivou a criar este projeto?

Quem me conhece sabe que sou muito fã de jogos e sou apaixonado por programação, e amo quando vejo oportunidades de unir as duas coisas. Bem, encontrei esta oportunidade de unir as duas coisas em um jogo chamado [Cypher](https://store.steampowered.com/app/746710/Cypher/), este jogo consiste em vários desafios no campo da criptografia, conhecendo assim boa parte da sua origem, começando por técnicas como a [Esteganografia](https://pt.wikipedia.org/wiki/Esteganografia) que é a arte de ocultar mensagens dentro de outra mensagem, passando também pelas [Cifra de Cézar](https://pt.wikipedia.org/wiki/Cifra_de_C%C3%A9sar), [Cifra de Vigenère](https://pt.wikipedia.org/wiki/Cifra_de_Vigen%C3%A8re), Máquina Enigma (foco deste projeto), chegando até a [Criptografia Digital](https://www.usna.edu/Users/cs/wcbrown/courses/si110AY13S/lec/l27/lec.html) Bem, quando eu estava jogando e cheguei nos desafios da Máquina Enigma notei que nao sabia como ela funciona de fato (apesar de ser fã do filme [Jogo da Imitação](https://pt.wikipedia.org/wiki/O_Jogo_da_Imita%C3%A7%C3%A3o) rsrs), e por este motivo iniciei os meu estudos com o intuito de saciar esta curiosidade, e em um dado momento pensei: "mais do que apenas resolver os desafios propostos pelo jogo por que não crio a minha versão da Máquina Enigma para resolver estas questões por mim?", e bem, foi basicamente esta pergunta que me trouxe até aqui rs.

## Qual a importância de aprender sobre a Máquina Enigma?

Bem, caso seja uma pessoa curiosa assim como eu o simples fato de saciar esta sede ja seria um bom motivo, mas consigo pensar em mais alguns motivos para aprender sobre a Máquina Enigma que são os seguintes:

* Na sociedade toda criação feita hoje é baseada no aprendizado de ontem, da mesma forma que entender sobre a esteganografia, Cifras de Cezar e Vigenère ajudam a entender o por quê da criação  da Máquina Enigma, entender sobre a Máquina Enigma ajuda entender sobre tecnicas de criptografia mais recentes como por exemplo [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard), [RSA](https://pt.wikipedia.org/wiki/RSA_(sistema_criptogr%C3%A1fico)),  [SHA-256](https://pt.wikipedia.org/wiki/SHA-2);
* A Máquina Enigma foi um importante fator durante a Segunda guerra Mundial, entender sobre a máquina ajuda entender mais o contexto de algumas coisas sobre este tema;
* Antigamente quando existia um conflito entre nações o embate acontecia principalmente no âmbito físico, nos dias de hoje a guerra acontece principalmente no âmbito das informações/dados, saber proteger os dados virou assunto de segurança nacional. Entender sobre a Máquina Enigma (e por consequência devido aos pontos anteriores sobre a criptografia em si) ajuda entender como esses embates acontecem no âmbito das nações e corporações.

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

Listarei agora o passo a passo de como a Máquina Enigma funciona:

* **Passo 1:** Todo o processo da Máquina Enigma começa no teclado, quando uma tecla é pressionada é liberado um sinal que corresponde a letra, esse sinal passa primeiro pela placa de conexões;

* **Passo 2:** Recebendo o sinal do Passo anterior, a placa de conexões faz uma validação para saber como proceder, por exemplo, se o sinal enviado pelo teclado foi a letra "A" e na placa de conexões a letra "A" esta plugada com a letra "F", neste caso o sinal vai ser alterado para representar a letra "F", caso contrario o sinal será enviado de como letra "A" mesmo, em ambos os caso o sinal vai ser enviado para o próximo passo, que são os rotores;

* **Passo 3:** O sinal repassado pelo passo anterior chega nos rotores, aqui é o "coração" de todo o processo pois aqui é onde acontece a maior parte dos embaralhamentos. O rotor é uma peça cilindrica onde no seu entorno tem a representação dos 26 caracteres de "A" a "Z", toda vez que um sinal chega ao rotor mais a direita ele muda sua posição, então por exemplo, se antes do sinal o rotor estava na letra "A" ele passa para letra "B", recebe o sinal para o proximo rotor como sendo a letra "B" e os proximos rotores repetem o mesmo processo com excessão de que, o próximo rotor só vai mudar a sua posição quando o rotor da sua direita completar um giro completo, passando assim por todas as letras, e dessa forma o sinal e trasmitido para os proximos rotores onde cada um realizara a alteração do caracter até chegar no proximo passo, que é o refletor;

* **Passo 4:** O sinal do passo anterior chega agora no refletor, que nada mais é que mais um rotor, mas com um pequena diferença, ao invés de ele passar o dado para frente ele faz a alteração do sinal, alterando assim mais uma vez a representação do caracter em questão e retorna esse dado pelo caminho que ele veio. Dessa forma, ele vai retornar por todos os rotres tendo assim mais uma rodada de alterações até passar pelo ultimo rotor mais a esquerda indo assim para o proximo passo que é a placa de conexões novamente;

* **Passo 5:** Neste passo, ele repetira o processo do passo 2 avaliando qual foi a representação de caracter que os rotores enviaram, caso a letra em questão esteja plugada em outra letra o caracter será alterado mais uma vez, caso contrário não sofrerá alterações e sera enviado para o último passo, que é o painel de lâmpadas;

* **Passo 6:** No painel de lampadas ele recebbe o sinal do passo anterior e acende uma lampada que representa o caracter cifrado, e dessa forma o processo é encerrado até que a proxima tecla seja pressionada reiniciando o ciclo.

  





