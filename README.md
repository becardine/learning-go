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

### Comandos

- `go run main.go`: compila e executa o código
- `go build main.go`: compila o código
- `./main`: executa o binário
- `go install`: compila e instala o binário no `$GOPATH/bin`
- `go get`: baixa e instala pacotes e dependências
- `go test`: executa testes
- `go fmt`: formata o código
- `go vet`: verifica erros no código
- `go doc`: documentação
- `go mod`: gerenciamento de dependências
- `go clean`: remove arquivos gerados pelo build
- `go env`: informações sobre o ambiente Go
- `go version`: versão do Go
- `go list`: lista pacotes
- `go tool`: ferramentas auxiliares
- `GOOS=linux GOARCH=amd64 go build main.go`: compila para Linux
- `GOOS=windows GOARCH=amd64 go build main.go`: compila para Windows
- `GOOS=darwin GOARCH=amd64 go build main.go`: compila para macOS
- `go mod init github.com/username/repo`: cria um arquivo go.mod
- `go mod tidy`: remove dependências não utilizadas ou adiciona dependências necessárias, alternativa ao `go get`
- `go.sum`: arquivo que contém as dependências do projeto e suas versões exatas, através do hash SHA-256 de cada pacote e versão baixada do repositório de origem.
- `go mod edit -replace github.com/username/repo@v1.0.0=github.com/username/repo@v1.1.0`: substitui uma dependência por outra (local pode ser usado para substituir um pacote remoto por um local, mas produção não é recomendado)
- `go work init`: cria um módulo de trabalho

### RUNTIME

```
go runtime + seu código = binário
```
