# Estrutura do projeto

```
NEADTech/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   │
│   ├── platform/
│   │   ├── config/
│   │   │   └── config.go
│   │   └── database/
│   │       └── mysql.go
│   │
│   ├── user/
│   │   ├── domain/
│   │   │   └── user.go
│   │   ├── repository/
│   │   │   └── user_repository.go
│   │   ├── service/
│   │   │   └── user_service.go
│   │   └── handler/
│   │       ├── user_handler.go
│   │       └── dto.go
│   │
│   └── http/
│       └── router.go
│
│
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```
