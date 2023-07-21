# BeginGoAPI
Um ponto de partida para explorar a construção de APIs com Go.

Iniciando com APIs em Go (Golang), o **BeginGoAPI** é um marco na minha jornada de autoaprendizagem. Embora a resposta seja um simples "Hello, World!", o projeto serve como uma introdução prática à construção de servidores web usando a linguagem Go.

## Funcionalidades

- Resposta HTTP simples com "Hello, World!".
- Uso do pacote `net/http` para gerenciamento de rotas e servidor.

## Como Rodar

### Pré-requisitos

- [Go (Golang)](https://golang.org/dl/) instalado na sua máquina.

### Instruções

1. Clone este repositório:
   #git clone https://github.com/Luizecafe/BeginGoAPI

2. Entre no diretório do projeto:
  #cd BeginGoAPI

3. Execute a aplicação:
  #go run main.go

4. Acesse http://localhost:8080 em seu navegador. Você deve ver a resposta "Hello, World!".

# Detalhes Técnicos
Pacote net/http: A espinha dorsal deste projeto. Através deste pacote, pude:

Definir rotas com http.HandleFunc.
Iniciar o servidor usando http.ListenAndServe.
Manipular respostas HTTP e ler informações de requisição usando http.ResponseWriter e http.Request.
# Feedback e Contribuições
Feedbacks, sugestões de melhorias e contribuições são muito bem-vindos. Sinta-se à vontade para abrir uma issue ou enviar um pull request.
