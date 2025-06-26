
# 📚 Biblioteca API - CRUD com Go + MongoDB + Docker

Este projeto é uma API REST simples para gerenciar livros de uma biblioteca, usando:

- Golang 1.21
- MongoDB (via Docker)
- Gin Web Framework
- Docker Compose

---

## 🚀 Como executar o projeto

### 📦 Pré-requisitos

- [Go 1.21](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### ▶️ Rodando com Docker Compose

1. Clone este repositório:

```bash
git clone https://github.com/seu-usuario/biblioteca-api.git
cd biblioteca-api
```

2. Rode o projeto:

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
titulo: Dom Casmurro
autor: Machado de Assis
desc: Romance clássico brasileiro
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
{ "message": "Nenhum livro encontrado. Cadastre um novo." }
```

---

### 🔍 3. Obter Livro por ID

- **GET** `/livros/obter?id=<id>`

**Exemplo:**
```bash
GET http://localhost:8080/livros/obter?id=665f1d2a4a99a824e1f4bb01
```

---

### ✏️ 4. Atualizar Livro

- **PUT** `/livros?id=<id>`
- **Content-Type:** `multipart/form-data`
- Pode atualizar 1 ou mais campos.

**Exemplo:**
```bash
PUT http://localhost:8080/livros?id=665f1d2a4a99a824e1f4bb01
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

- **DELETE** `/livros?id=<id>`

**Exemplo:**
```bash
DELETE http://localhost:8080/livros?id=665f1d2a4a99a824e1f4bb01
```

---

## ⚙️ Tecnologias e Dependências

- [Go 1.21](https://golang.org/doc/go1.21)
- [Gin](https://github.com/gin-gonic/gin)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [MongoDB (via Docker)](https://hub.docker.com/_/mongo)

---

## 📬 Contato

João Carlos  
🔗 [linkedin.com/in/seu-usuario](https://linkedin.com/in/seu-usuario)  
📧 seu.email@exemplo.com

---
