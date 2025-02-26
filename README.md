# Go-Go-Gadget-Code

## 📌 Visão Geral
Go-Go-Gadget-Code é um projeto desenvolvido em Go com o objetivo de integrar com a API da OpenAI para permitir conversas interativas e armazenar tópicos de discussão em um banco de dados.

## 🚀 Funcionalidades Futuras
- **Integração com OpenAI**: Permitir que os usuários interajam com modelos de IA e salvem tópicos relevantes no banco de dados.
- **Autenticação de Usuário**: Implementação de autenticação JWT para segurança dos dados.
- **Gerenciamento de Usuários**: CRUD completo para gerenciar perfis e históricos de conversa.
- **Painel de Administração**: Interface para visualizar e gerenciar conversas armazenadas.

## 🛠 Tecnologias Utilizadas
- **Linguagem**: Go
- **Framework Web**: `gin-gonic/gin` para criação de APIs RESTful
- **Banco de Dados**: Microsoft SQL Server (MSSQL)
- **ORM**: GORM para interação com o banco de dados
- **Autenticação**: JWT para segurança e controle de acesso
- **API Externa**: OpenAI para processamento de linguagem natural
- **Gerenciamento de Configuração**: Arquivo `config.go` para armazenar variáveis de ambiente e configurações do sistema

## 📂 Estrutura do Projeto
```
Go-Go-Gadget-Code/
│── main.go
│── go.mod
│── config/
│   ├── config.go
│── handlers/
│   ├── user_handler.go
│── models/
│   ├── user.go
│── repository/
│   ├── user_repository.go
│── services/
│   ├── user_service.go
```

## 📌 Rotas Disponíveis
- **Obter Usuários**
  - **Rota**: `GET /users`
  - **Exemplo de requisição**:
    ```sh
    curl http://localhost:8080/users
    ```

## 📖 Como Rodar o Projeto
1. Clone o repositório
   ```sh
   git clone https://github.com/seu-usuario/go-go-gadget-code.git
   cd go-go-gadget-code
   ```
2. Instale as dependências
   ```sh
   go mod tidy
   ```
3. Configure as variáveis de ambiente no `config.go`
4. Execute o projeto
   ```sh
   go run main.go
   ```

## 📌 Contribuição
Sinta-se à vontade para abrir issues e enviar pull requests para melhorias no projeto!

---


