# trainingAlgorithms

# 1.

Na loja de animais à venda, existem alguns tipos de ração disponíveis para compra, sendo eles:

-> Ração para cachorro;
-> Ração para gato; e
-> Ração universal.

O estoque dessas rações está representado em um mapa, onde a chave é o nome da ração e o valor correspondente é a quantidade disponível em estoque.

Polycarp possui x cães e y gatos. Gostaríamos de determinar se é possível para ele comprar comida suficiente para todos os seus animais na loja. Cada um dos seus cães e gatos deve receber um pacote de ração adequado para sua espécie.

# Exemplo de entrada:

Quantidade de cachorros: 2
Quantidade de gatos: 3
Estoque: map[string]int{
"dog": 2,
"cat": 2,
"universal: 2",
}

# Exemplo de saída:

É possível: true

# Explicação:

Nesse caso, é possível comprar 2 pacotes de ração para cachorro e 3 pacotes de ração para gato, totalizando 5 pacotes de ração. Como o estoque possui 6 pacotes de ração, é possível comprar a quantidade necessária para alimentar todos os animais.

# Exemplo de entrada:

Quantidade de cachorros: 3
Quantidade de gatos: 3
Estoque: map[string]int{
"dog": 2,
"cat": 2,
"universal: 1",
}

# Exemplo de saída:

É possível: false

# Explicação:

Nesse caso, é possível comprar 2 pacotes de ração para cachorro, 2 pacotes de ração para gato e 1 pacote de ração universal, totalizando 5 pacotes de ração. Como o estoque possui 5 pacotes de ração, não é possível comprar a quantidade necessária para alimentar todos os animais.

# 2.

O torneio de programação do CEUB ocorrerá em breve. Neste ano, equipes de quatro pessoas estáo autorizadas a participar.

No UniCEUB, temos um grupo de participantes que inclui programadores e matemáticos. Gostaríamos de saber o número máximo de equipes que podem ser formadas, considerando as seguintes regras:

-> Cada equipe deve ser composta por exatamente quatro estudantes;
-> Equipes compostas apenas por quatro matemáticos ou apenas por quatro programadores não têm um bom desempenho, portanto, decidiu-se não formar tais equipes;
-> Assim, cada equipe deve ter pelo menos um programador e pelo menos um matemático; e
-> Cada pessoa só pode fazer parte de uma equipe.

Escreva um programa que receba como entrada uma lista de participantes e retorne o número máximo de equipes que podem ser formadas, respeitando as regras mencionadas.

# Exemplo de entrada:

Participantes: []Participant{
        {
          Name: "Pedro",
          Role: "Programmer",
        },
        {
          Name: "Vanessa",
          Role: "Programmer",
        },
        {
          Name: "Tônia",
          Role: "Mathematician",
        },
        {
          Name: "Jõao",
          Role: "Mathematician",
        },
}

# Exemplo de saída:

Número máximo de equipes: 1

# Explicação:

Nesse caso, apenas uma equipe pode ser formada e a divisão dos grupos é a mesma lista de participantes.

# 3.

O cozinheiro Remy preparou uma refeição para si mesmo e, enquanto almoçava, decidiu assistir a um vídeo no RataTube, No entanto, ele tem um tempo limitado de t segundos para o almoço. Ele pediu a sua ajuda para escolher o vídeo ideal.

O RataTube possui um feed de n vídeos, cada um representado por uma estrutura de vídeos, contendo informações sobre sua duração em segundos e o nível de entretenimento. O feed é inicialmente aberto no primeiro vídeo e Remy pode pular para o próximo vídeo em 1 segundo (caso exista). Ele pode pular vídeos quantas vezes desejar, inclusive não pular nenhum.

Sua tarefa é auxiliar Remy a escolher um vídeo que ele possa abrir e assistir dentro do tempo disponível, t segundos. Se houver vários vídeos que se encaixem nessa condição, ele deseja escolher o vídeo com o maior nível de entretenimento. Retorne qualquer vídeo apropriado ou exiba um erro caso não haja um vídeo adequado dentro do tempo disponível. E, caso não seja possível assistir nenhum vídeo, retorne um erro.

# Exemplo de entrada:

t = 5
videos = []Video{
        {
          ID: 1,
          Duration: 2,
          Entertainment: 3,
        },
        {
          ID: 2,
          Duration: 3,
          Entertainment: 4,
        },
        {
          ID: 3,
          Duration: 1,
          Entertainment: 2,
        },
}

# Exemplo de saída:

Video{
    ID: 2,
    Duration: 3,
    Entertainment: 4,
}

# Explicação:

Nesse caso, o vídeo com maior nível de entretenimento que pode ser assistido dentro do tempo disponível é o segundo vídeo.

# Exemplo de entrada:

t = 10
videos = []Video{
        {
          ID: 1,
          Duration: 6,
          Entertainment: 8,
        },
        {
          ID: 2,
          Duration: 7,
          Entertainment: 5,
        },
        {
          ID: 3,
          Duration: 9,
          Entertainment: 9,
        },
        {
          ID: 4,
          Duration: 4,
          Entertainment: 6,
        },
}

# Exemplo de saída:

Video{
    ID: 1,
    Duration: 6,
    Entertainment: 8,
}

# Explicação:

Nesse caso, o vídeo com maior nível de entretenimento que pode ser assistido dentro do tempo disponível é o primeiro vídeo. Isso se deve ao fato de que é necessário gastar 1 segundo para pular para o próximo vídeo, e o terceiro vídeo, que possui o maior nível de entretenimento, não pode ser assistido dentro do tempo disponível.

# 4.

Receba uma lista de camisetas, cada uma contendo o preço e o tamanho. O tamanho da camisa é representado por uma string, que pode ser "M" ou uma combinação de caracteres "X" seguida por "S" ou "L".

Por exemplo, as strings "M", "XXL", "S" e "XXXXXXXS" podem representar tamanhos de algumas camisas. Já as strings "XM", "LL" e "SX" não representam tamanhos válidos.

O objetivo é calcular a média entre o preço da maior camisa e o preço da menor camisa da lista.

A comparação entre os tamanhos das camisas deve seguir as seguintes regras:

-> Qualquer tamanho pequeno (independentemente da quantidade de caracteres "X") é menor que o tamanho médio e qualquer tamanho grande;
-> Qualquer tamanho grande (independentemente da quantidade de caracteres "X") é maior que o tamanho médio e qualquer tamanho pequeno;
-> Quanto mais caracteres "X" antes de "S", menor é o tamanho; e
-> Quanto mais caracteres "X" antes de "L", maior é o tamanho.

Por exemplo:

1. "XXXS" < "XS"
2. "XXXL" > "XL"
3. "XL" > "M"
4. "XXL" = "XXl"
5. "XXXXXS" < "M"
6. "XL" > "XXXS"

Dessa forma, ao receber a lista de camisas com seus respectivos preços e tamanhos, você deve calcular a média de preços da maior e da menor camisa. E, caso não seja possível calcular a média, retorne um erro.

# Exemplo de entrada:

Camisas = []Shirt{
        {
          Size: "M",
          Price: 10,
        },
        {
          Size: "XXL",
          Price: 20,
        },
        {
          Size: "S",
          Price: 7,
        },
        {
          Size: "XXXXXXXS",
          Price: 5,
        },
}
        
# Exemplo de saída:

Maior: 20
Menor: 5

# Explicação:

Nesse caso, a média da maior camisa (XXL) é de 20 e a média da menor camisa (XXXXXXXS) é de 5.

# Exemplo de entrada:

Camisas = []Shirt{
        {
          Size: "S",
          Price: 10,
        },
        {
          Size: "XXL",
          Price: 20,
        },
        {
          Size: "S",
          Price: 4,
        },
        {
          Size: "XXXXXXXS",
          Price: 5,
        },
        {
          Size: "XXXXXXXL",
          Price: 25,
        },
}
        
# Exemplo de saída:

Maior: 15
Menor: 7

# Explicação:

Nesse caso, a média da maior camisa (XXXXXXXL) é de 15 e a média da menor camisa (S) é de 7.

# 5.

Um novo serviço de e-mail, chamado "CEUBdesk", será inaugurado no CEUB em um futuro próximo. A administração do site quer lançar o projeto o mais rápido possível, por isso eles pedem a sua ajuda. Você é sugerido(a) a implementar o protótipo do sistema de registro do site. O sistema deve funcionar com o seguinte princípio:

Cada vez que um novo usuário deseja se registrar, ele envia ao sistema uma solicitação com o seu nome. Se esse nome ainda não existe no banco de dados do sistema, ele é inserido no banco de dados e o usuário recebe uma resposta "OK", confirmando o registro bem-sucedido. Se o nome já existe no banco de dados do sistema, o sistema cria um novo nome de usuário, envia-o ao usuário como sugestão e também insere a sugestão no banco de dados. O novo nome é formado pela seguinte regra:

Números, começando com 1, são anexados um após o outro ao nome (name1, name2, ...), entre esses números, o menor i é encontrado de forma que namei ainda não exista no banco de dados.

# Exemplo de entrada:

["alex", "steven", "alex", "steven", "alex", "steven"]

# Exemplo de saída:

["OK", "OK", "alex1", "steven1", "alex2", "steven2"]

# Explicação:

Nesse caso, o primeiro usuário recebe uma reposta "OK", pois ele é o primeiro a se registrar. O segundo usuário também recebe um "OK", pois ele é o segundo a se registrar. O terceiro usuário recebe uma sugestão "alex1", pois o nome "alex" já existe no banco de dados. O quarto usuário recebe uma sugestão "steven1", pois o nome "steven" já existe no banco de dados. O quinto usuário recebe uma sugestão "alex2", pois os nomes "alex" e "alex1" já existe no banco de dados. O sexto usuário recebe uma sugestão "steven2", pois os nomes "steven" e "steven1" já existe no banco de dados. E assim, por diante.

# 6.

Implemente uma função que inverta uma lista encadeada, ou seja, altere a ordem dos elementos da lista de forma reversa.

# Estrutura da lista encadeada:

type Node struct {
    Value int
    Next *Node
}

type LinkedList struct {
    Head *Node
}

# Exemplo de entrada:

lista := LinkedList{
        Head: &Node{
                Value: 1,
                Next: &Node{
                        Value: 2,
                        Next: &Node{
                                Value: 3,
                                Next: nil,
                        },
                },
        },        
}
invertedLista(&lista)

# Exemplo de saída:

Não há saída direta da função, mas após a chamada à função, a ordem dos elementos na lista deve ser invertida.

# Explicação:

A lista inicial [1, 2, 3] deve ser invertida para [3, 2, 1].
