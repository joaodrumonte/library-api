
# 📚 Biblioteca API - CRUD com Go + MongoDB + Docker

Este projeto é uma API para gerenciar livros de uma biblioteca criada para um desafio técnico, usando:

- Golang 1.21
- MongoDB (via Docker)
- Docker Compose

---

## 🚀 Como executar o projeto

### 📦 Pré-requisitos

- [Go 1.21](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### ▶️ Rodando com Docker Compose

```bash
docker-compose up --build
```

> A aplicação estará disponível em `http://localhost:8080`

---

## 🔗 Endpoints da API

### 📘 1. Criar Livro

- **POST** `/livros`
- **Content-Type:** `multipart/form-data`

**Body:**
| Campo | Obrigatório | Tipo |
|-------|-------------|------|
| titulo | ✅ | string |
| autor  | ✅ | string |
| desc   | ✅ | string |

**Exemplo Postman (form-data):**
```bash
POST http://localhost:8080/livros
```
```form-data
titulo: O senhor dos Anéis
autor: J. R. R. Tolkien
desc: Numa cidadezinha indolente do Condado, um jovem hobbit é encarregado de uma imensa tarefa. Deve empreender uma perigosa viagem através da Terra-média até as Fendas da Perdição, e lá destruir o Anel do Poder - a única coisa que impede o domínio maléfico do Senhor do Escuro.
```

---

### 📚 2. Listar Livros

- **GET** `/livros`

**Exemplo:**
```bash
GET http://localhost:8080/livros
```

Se nenhum livro for encontrado, retorna:
```json
{ "error": "Nenhum livro encontrado. Cadastre um novo." }
```

---

### 🔍 3. Obter Livro por ID

- **GET** `/livro?id=<id>`

**Exemplo:**
```bash
GET http://localhost:8080/livro?id=665f1d2a4a99a824e1f4bb01
```

---

### ✏️ 4. Atualizar Livro

- **PUT** `/livro?id=<id>`
- **Content-Type:** `multipart/form-data`
- Pode atualizar 1 ou mais campos.

**Exemplo:**
```bash
PUT http://localhost:8080/livro?id=665f1d2a4a99a824e1f4bb01
```
```form-data
autor: Novo Autor
```

> Retorna erro se nenhum campo for enviado:
```json
{ "error": "Pelo menos um campo (titulo, autor ou desc) deve ser enviado para atualização" }
```

---

### 🗑️ 5. Deletar Livro

- **DELETE** `/livro?id=<id>`

**Exemplo:**
```bash
DELETE http://localhost:8080/livro?id=665f1d2a4a99a824e1f4bb01
```

---

## ⚙️ Tecnologias e Dependências

- Go 1.21
- Gin
- MongoDB Go Driver
- MongoDB (via Docker)

---


