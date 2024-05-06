### O que é a linguagem Go? Quem criou? O que não é Go?

```
 ⚠️ Uma linguagem de programação não vai resolver todos os seus problemas.
```

- Linguagem de programação open source criada pela Google em 2007.
- Rob Pike (Unix e UTF-8), Ken Thompson (Unix e UTF-8) e Robert Griesemer (V8) são os criadores da linguagem.
- Versão 1.0 foi lançada em 2012.

```
 A partir da versão 1.5, o compilador Go é escrito em Go.
```

- É uma linguagem compilada e fortemente tipada.
- Possui uma sintaxe simples e fácil de aprender.
- É uma linguagem de alto desempenho, com uma execução rápida e ao mesmo tempo trabalha com garbage collection.
- Foi criada para aproveitar ao máximo os processadores multi-core e de rede
- Compilada em apenas um arquivo binário, sem necessidade de instalar bibliotecas ou dependências.
- Linguagem retrocompatível na versão 1, ou seja, programas escritos em versões anteriores da linguagem continuarão funcionando nas versões mais recentes.
- Framework de testes e profiling nativos
- Detecção de Race Conditions nativa
- Deploy simples
- Baixa curva de aprendizado

### Instalação

- [Download Go](https://golang.org/dl/)
  - chocolatey: `choco install golang` (Windows)

### Principais pastas `/user/go`

- `src`: onde podem ser criados os projetos
- `pkg`: onde ficam os pacotes compilados
- `bin`: onde ficam os binários

### Extensão VSCode

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- `Ctrl + Shift + P` > `Go: Install/Update Tools` > `all`
