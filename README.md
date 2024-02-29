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


