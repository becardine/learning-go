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
- `go test -v`: executa testes com verbose
- `go test -cover`: executa testes com cobertura
- `go test -coverprofile=coverage.out`: gera um arquivo de cobertura
- `go tool cover -html=coverage.out`: abre o arquivo de cobertura no navegador
- `go test -bench .`: executa testes de benchmark
- `go test fuzz`: executa testes de fuzzing
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
- `go mod vendor`: cria uma pasta `vendor` com as dependências do projeto
- `go.sum`: arquivo que contém as dependências do projeto e suas versões exatas, através do hash SHA-256 de cada pacote e versão baixada do repositório de origem.
- `go mod edit -replace github.com/username/repo@v1.0.0=github.com/username/repo@v1.1.0`: substitui uma dependência por outra (local pode ser usado para substituir um pacote remoto por um local, mas produção não é recomendado)
- `go work init`: cria um módulo de trabalho (não é recomendado para projetos que serão compartilhados)

### RUNTIME

```
go runtime + seu código = binário
```

### Multithreading

- **Go Routines** (Goroutines): funções que são executadas de forma concorrente
- Rodar simultaneamente em um único processo (thread)
- Utiliza a função `go` para criar uma goroutine
- **WaitGroup**: sincroniza as goroutines
  - Adicionar quantidades de tarefas/operações;
  - Informar que você terminou uma tarefa/operação;
  - Aguardar a finalização de todas as tarefas/operações.
- **Channels**: comunicação entre goroutines
- Segurança para uma thread saber o momento em que ela pode trabalhar com um determinado dado
- Esperar até que as operações sejam concluídas
  - `make(chan tipo)`: cria um canal
  - `canal <- valor`: envia um valor para o canal
  - `valor := <- canal`: recebe um valor do canal
  - `close(canal)`: fecha o canal
  - `range canal`: itera sobre os valores do canal

### Eventos

- Algo no passado. Exemplo: Inseri o registro -> registro inserido.
- Evento (Carregar dados) -> operações que são executadas em sequência (Carregar dados -> Salvar dados -> Enviar email) -> Gerenciar os eventos/operações

### RabbitMQ

- **Message Broker**: intermediário de mensagens
- **Producer**: envia mensagens
- **Consumer**: recebe mensagens
- **Queue**: fila de mensagens
- **Exchange**: roteamento de mensagens. Tipos: fanout, direct, topic, headers. O producer envia mensagens para o exchange e o exchange envia para a fila.
  - direct: roteamento para uma fila. precisa de routing key
  - fanout: roteamento para todas as filas. não precisa de routing key
  - topic: parecido com o direct, mas com wildcards. precisa de routing key. `*` (uma palavra) e `#` (zero ou mais palavras). Exemplo: `*.amarelo.*` ou `#`
- **Binding**: regras de roteamento
- **Connection**: conexão com o broker
- **Channel**: canal de comunicação
- **Delivery**: mensagem entregue
- **Acknowledge**: confirmação de entrega
- **Requeue**: reenviar mensagem
- **Dead Letter Exchange**: troca de mensagens mortas
- **Dead Letter Queue**: fila de mensagens mortas
- **Prefetch**: quantidade de mensagens pré-carregadas
- **QoS**: qualidade de serviço
- **Durable**: mensagens persistentes
- **Transient**: mensagens não persistentes
- **Exclusive**: fila exclusiva
- **Auto Delete**: fila excluída automaticamente
- **Binding Key**: chave de roteamento
- **Routing Key**: chave de roteamento
- **RPC**: chamada de procedimento remoto
- **Pub/Sub**: publicar/assinar
- **Message**: mensagem
- **Payload**: carga útil
- **Broker**: intermediário
- **Cluster**: agrupamento
- **Shovel**: transferência de mensagens
- **Federation**: federação de mensagens
- **Management Plugin**: plugin de gerenciamento
- **Stomp**: protocolo de mensagens
- **Web STOMP**: protocolo de mensagens via web
- **AMQP**: protocolo de mensagens avançado
- **STOMP**: protocolo de mensagens simples
- **MQTT**: protocolo de mensagens para IoT
- **WebSocket**: protocolo de comunicação
- **TLS**: segurança de transporte
- **SSL**: segurança de transporte
- **SASL**: camada de segurança
- **OAuth**: autenticação

### GraphQL for Go

- [gqlgen](https://gqlgen.com/): gerador de código GraphQL
- `printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go`: cria um arquivo `tools.go` para instalar as ferramentas do gqlgen
- `go run github.com/99designs/gqlgen init`: inicializa o gqlgen
- `go run github.com/99designs/gqlgen generate`: gera o código GraphQL
- Resolver: função que executa a lógica de uma query ou mutation

### gRPC

- [gRPC](https://grpc.io/): framework RPC de alto desempenho. Desenvolvido pela Google, para facilitar a comunicação entre serviços independente da linguagem de programação.
- Faz parte do CNCF (Cloud Native Computing Foundation)
- Ideal para comunicação entre microserviços
- Mobile, Backend
- Geração de forma automática de código
- Streaming bidirecional usando HTTP/2
- Linguagens com suporte oficial: gRPC-GO, gRPC-JAVA, gRPC-C -> C++, Phyton, Ruby, Objective C, PHP, C#, Node.js, Dart, Kotlin/JVM
- Remote Procedure Call (RPC): chamada de procedimento remoto
- Protocol buffers: serialização de dados (pense em um XML, porém mais rápido e eficiente)
- Protocol Buffers vs JSON
  - Arquivo binário < JSON
  - Serialização e desserialização mais rápidas e mais leves
  - Menos tráfego de rede
  - Processo mais rápido
- Sintaxe: segue um contrato. os números são referentes a ordem dos campos para controle interno do protocol buffer
- `protoc --go_out=. --go-grpc_out=. proto/entities.proto`: gera o código Go e gRPC

```proto3
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

```

- HTTP/2: multiplexação, cabeçalhos comprimidos, priorização de requisições, push de servidor para cliente
- Nasceu como SPDY (Google) e foi padronizado como HTTP/2
- Lançado em 2015
- Dados trafegados são binários e não textuais como no HTTP/1.1
- Utiliza a mesma conexão TCP para várias requisições (Multiplex)
- Server Push: servidor envia recursos para o cliente antes mesmo dele solicitar
- Cabeçalhos comprimidos: HPACK
- Gasta menos recursos de rede e processamento

- **gRPC - API "unary"**
  - Chamada de procedimento remoto
  - Cliente envia uma requisição e o servidor responde
  - Exemplo: `Unary RPC` (Cliente -> Servidor)
- **gRPC - API "server streaming"**
  - Cliente envia uma requisição e o servidor responde com uma sequência de mensagens
  - Exemplo: `Server Streaming RPC` (Cliente -> Servidor)
- **gRPC - API "client streaming"**
  - Cliente envia uma sequência de mensagens e o servidor responde
  - Exemplo: `Client Streaming RPC` (Cliente <- Servidor)
- **gRPC - API "bidirectional streaming"**
  - Cliente e servidor enviam uma sequência de mensagens
  - Exemplo: `Bidirectional Streaming RPC` (Cliente <-> Servidor)

### REST vs gRPC

| Característica      | REST                   | gRPC                       |
| ------------------- | ---------------------- | -------------------------- |
| Protocolo           | HTTP/1.1               | HTTP/2                     |
| Formato de mensagem | JSON, XML              | Protocol Buffers           |
| Verbos              | GET, POST, PUT, DELETE | RPC                        |
| Stateless           | Sim                    | Sim                        |
| Cache               | Sim                    | Sim                        |
| Segurança           | HTTPS, OAuth           | TLS                        |
| Ferramentas         | Postman, Swagger       | gRPC CLI                   |
| Comunicação         | Síncrona               | Síncrona e Assíncrona      |
| Comunicação         | Unidirecional          | Bidirecional               |
| Performance         | Mais lenta             | Mais rápida                |
| Comunicação         | HTTP/1.1               | HTTP/2                     |
| Comunicação         | Texto                  | Binário                    |
| Comunicação         | JSON                   | Protocol Buffers           |
| Comunicação         | -                      | Streaming                  |
| Comunicação         | -                      | Multiplexação              |
| Comunicação         | -                      | Cabeçalhos comprimidos     |
| Comunicação         | -                      | Priorização de requisições |

### cobra-cli

- [cobra-cli](github.com/spf13/cobra-cli) é uma ferramenta CLI para criar projetos com Cobra
- `cobra-cli init nome-do-projeto`: inicializa o projeto
- `cobra-cli add nome-do-comando`: adiciona um comando
- `cobra-cli add nome-do-comando --type=local`: adiciona um comando local
- `cobra-cli add nome-do-comando --type=alias`: adiciona um comando de alias
- `cobra-cli add nome-do-comando --type=parent`: adiciona um comando pai
- persistent flags: flags que são passadas para todos os comandos
- local flags: flags que são passadas apenas para o comando

### Migrations

- [golang-migrate/migrate](github.com/golang-migrate/migrate) é uma ferramenta para gerenciar migrações de banco de dados
- `migrate create -ext sql -dir migrations -seq nome-da-migracao`: cria uma migração
- `migrate -path migrations -database "mysql://root:root@tcp(localhost:3306)/sqlc?query" up`: executa as migrações
- `migrate -path migrations -database "mysql://root:root@tcp(localhost:3306)/sqlc?query" down`: reverte as migrações
- sqlx: extensão do pacote `database/sql` que facilita a execução de queries
- sqlc: gerador de código SQL para Go
- instalação: `go get github.com/kyleconroy/sqlc/cmd/sqlc`
- sqlc sempre que tiver uma coluna do tipo decimal, ele vai salvar como string, para corrigir isso, é necessário adicionar um override no arquivo `sqlc.yaml`, exemplo:
  ```yaml
  overrides:
    - db_type: 'decimal'
      go_type: 'float64'
  ```
- `sqlc generate`: gera o código SQL

### UOW - Unit of Work

- **Unit of Work**: unidade de trabalho que agrupa todas as operações de um banco de dados em uma única transação
- **Repository**: camada de abstração que isola a lógica de negócio da lógica de acesso a dados
- **Service**: camada de abstração que isola a lógica de negócio da lógica de acesso a dados
- **Model**: estrutura de dados que representa uma entidade do banco de dados
- **Controller**: camada de abstração que recebe as requisições HTTP e chama os serviços
- **Handler**: camada de abstração que recebe as requisições HTTP e chama os serviços
- **Middleware**: intercepta as requisições HTTP e executa uma lógica antes de passar para o próximo handler
- **Context**: passa valores entre os middlewares e handlers
- **Error Handling**: tratamento de erros
- **Logger**: registra as requisições HTTP
- **Config**: configurações do projeto
- **Docker**: containerização
- **Docker Compose**: orquestração de containers
- **Makefile**: automatiza tarefas
- **Testes**: testes unitários e de integração
- **Mock**: simula o comportamento de uma dependência
- **Dockerfile**: arquivo de configuração do Docker

### DI - Dependency Injection

- **Dependency Injection**: técnica de injeção de dependências
- **Container**: armazena as dependências e resolve as dependências
- [FX](github.com/uber-go/fx): framework de injeção de dependências
  - `fx.New()`: cria um novo container
  - Utiliza reflection para injetar as dependências
- [Wire](github.com/google/wire): gerador de código para injeção de dependências
  - `wire`: gera o código de injeção de dependências
  - Não utiliza reflection
  - `go install github.com/google/wire/cmd/wire@latest`: instala o wire
  - `go run main.go wire_gen.go`: rodar a aplicação junto com o wire

### Clean Architecture

- Arquitetura de software que separa as responsabilidades em camadas
- Termo criado por Robert C. Martin (Uncle Bob) em 2012
- Livro "Clean Architecture: A Craftsman's Guide to Software Structure and Design" (2017)
  - Curiosidades do livro:
    - ele fala especificamente sobre "Clean Architecture" somente em 7 páginas do livro
    - tudo que ele fala sobre "Clean Architecture" está literalmente em um artigo do blog dele
- Buzz word no mundo do desenvolvimento de software
- Proteção do domínio da aplicação
- Baixo acoplamento entre as camadas (limites arquiteturais)
- Orientada a casos de uso (use cases)

  - Use Cases: regras de negócio da aplicação

- **Pontos sobre arquitetura:**

  - formato que o software terá
  - divisão de componentes
  - comunicação entre componentes
  - uma boa arquitetura vai facilitar o processo de desenvolvimento, deploy e manutenção
  - "The strategy behind that facilitation is to leave as many options opem as possible for as long as possible" - Uncle Bob

- **Objetivos de uma boa arquitetura:**

  - "o objetivo principal da arquitetura é dar suporte ao ciclo de vida do sistema. Uma boa arquitetura torna o sistema fácil de entender, fácil de desenvolver, fácil de manter e fácil de implantar. O objetivo final é minimizar o custo de vida útil do sistema e maximizar a produtividade do programador." - Uncle Bob

```
 "Keep options open" - Uncle Bob
```

### Regras vs Detalhes

- Regras de negócio trazem o real valor para o software
- Detalhes ajudam a suportar as regras de negócio
- Detalhes não devem impactar as regras de negócio
- Frameworks, banco de dados, apis, não devem impactar as regras de negócio

```
Atacar a complexidade no coração do software (DDD)
```

### Use Cases

- Intenção de uma ação do usuário
- Clareza de cada comportamento do sistema
- Detalhes não devem impactar as regra de negócio
- Use cases contam uma história

#### Use Cases vs SRP (Single Responsibility Principle)

- Temos a tendência de "reaproveitar" use cases por serem muitos parecidos
- Ex: Alterar vs Inserir. Ambos consultam se o registro existe, persistem dados. Mas são use cases diferentes. Por quê?

  - **Alterar**: verifica se o registro existe, persiste os dados
  - **Inserir**: não verifica se o registro existe, persiste os dados
  - **SRP**: mudam por razões diferentes, logo são responsabilidades diferentes

### Limites Arquiteturais

- Tudo que não impacta diretamente as regras de negócio deve estar em um limite arquitetural diferente. Ex: Não será o frontend, bando de dados que mudarão as regras de negócio da aplicação.

#### Input vs Output

- Tudo se resume a um Input que retorna um Output
- Ex: Criar um pedido (dados do pedido = input) -> Pedido criado (output)
- Simplifique seu raciocínio ao criar um software sempre pensando em Input e Output

### DTO - Data Transfer Object

- Trafegar dados entre os limites arquiteturais
- Não possuem regras de negócio
- Objeto anêmico, sem comportamento
- Não faz nada, apenas transporta dados
- API -> Controller -> Use Case -> Entity
  - Controller cria um DTO e passa para o Use Case
  - Use case executa o fluxo, pega o resultado, cria um DTO para output e passa para o Controller

### Presenters

- Objetos de transformação
- Adequa o DTO de output para o formato que o cliente deseja
- Um sistema pode ter diversos formatos de entrega: ex: JSON, XML, HTML, etc

```go
  input = new CategoryInputDto("name")
  output = CreateCategoryUseCase(input)
  jsonResult = CategoryPresenter(output).ToJson()
  xmlResult = CategoryPresenter(output).ToXml()
```

### Entities

- Entities da Clean Architecture são diferentes das Entities do DDD
- Entities da Clean Architecture são camadas de regras de negócio
- Elas se aplicam em qualquer situação
- Não há definição explicita de como criar uma Entity
- Normalmente utilizamos táticas do DDD para criar Entities
- Entities = Agregados + Domain Services
-
