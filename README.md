# Calculadora em Go

Uma calculadora de linha de comando escrita em Golang, capaz de interpretar e calcular expressões matemáticas básicas, incluindo operações com parênteses. O projeto implementa o algoritmo **Shunting Yard** para converter expressões infixas em pós-fixas (notação polonesa reversa), constrói uma **árvore de expressão** e avalia o resultado de forma recursiva.

## Funcionalidades

- **Operações suportadas:** adição (`+`), subtração (`-`), multiplicação (`*`), divisão (`/`)
- **Parênteses:** Suporte completo para precedência de operações usando parênteses
- **Algoritmo Shunting Yard:** Conversão eficiente de expressões infixas para pós-fixas
- **Árvore de Expressão:** Montagem de uma árvore binária para avaliação das expressões
- **Cálculo Recursivo:** Avaliação da árvore de expressão de forma recursiva
- **Interface CLI:** Interação simples via terminal

## Como funciona

1. **Entrada do Usuário:** O usuário digita uma expressão matemática no terminal.
2. **Conversão para Pós-fixa:** A expressão é convertida usando o algoritmo Shunting Yard.
3. **Construção da Árvore:** A expressão pós-fixa é transformada em uma árvore binária.
4. **Avaliação:** A árvore é avaliada recursivamente para obter o resultado final.

## Exemplo de uso

```bash
$ go run index.go
------ CALCULATOR -------
q - to quit
Expression: (2+3)*4
Result: 20

Expression: 7/(1+2)
Result: 2.3333333333333335

Expression: q
quitting...
```

## Estrutura do Projeto

```
.
├── index.go           # Arquivo principal, inicia a aplicação
├── calc/
│   ├── start.go       # Função de inicialização e loop de entrada
│   ├── shuntingYard.go# Implementação do algoritmo Shunting Yard
│   ├── tree.go        # Montagem e impressão da árvore de expressão
│   ├── eval.go        # Avaliação recursiva da árvore
│   └── types.go       # Funções auxiliares de verificação de tipos
├── go.mod             # Módulo Go
└── .gitignore
```

## Como rodar

1. Certifique-se de ter o Go instalado ([download aqui](https://golang.org/dl/)).
2. Clone o repositório:
   ```bash
   git clone <url-do-repositorio>
   cd go-calculator
   ```
3. Execute a calculadora:
   ```bash
   go run index.go
   ```

## Sobre a implementação

- **Algoritmo Shunting Yard:** Responsável por respeitar a precedência dos operadores e transformar a expressão para facilitar a avaliação.
- **Árvore de Expressão:** Cada nó representa um número ou operador, permitindo a avaliação correta mesmo em expressões complexas.
- **Tratamento de erros:** O programa lida com entradas inválidas, divisões por zero e outros casos comuns de erro.

## Contribuição

Sinta-se à vontade para abrir issues ou pull requests com melhorias, correções ou sugestões! 