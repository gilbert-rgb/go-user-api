# Go User API

A simple and clean REST API built with **Go**, **Gin**, **GORM**, **SQLite**, and **JWT authentication**.

This project is designed to be easy to run locally and demonstrates real-world API structure and best practices.
This project uses SQLite for simplicity.

No Docker Compose or PostgreSQL is included.

SQLite database file (app.db) is created automatically when the server runs.

---

## 🚀 Features

- User registration and login
- JWT-based authentication
- CRUD operations for users
- Clean project structure
- SQLite database for easy setup
- Optional Docker support

---

##  Project Structure
go-user-api/
├── cmd/
│ └── main.go
├── internal/
│ ├── config/
│ ├── handlers/
│ ├── models/
│ ├── routes/
│ └── middleware/
├── Dockerfile
├── go.mod
├── go.sum
├── app.db
└── README.md


---
##  Repository Link

 **Go User API Repository**  
[https://github.com/gilbert-rgb/go-user-api](https://github.com/gilbert-rgb/go-user-api)


##  Requirements

- Go 1.21+
- SQLite (automatically created)
- Docker (optional)

---

##  Setup (Local)

### 1. Install dependencies
```bash
go mod tidy
```

2. Run the server
```bash
go run cmd/main.go
```

3. Check server
```bash
curl http://localhost:8080/health

```

### Docker (Optional)
1.Build Docker image
```bash
docker build -t go-user-api .
```
2.Run Docker container
```bash
docker run -p 8080:8080 --env-file .env -d go-user-api
```

## Author &contacts
- Gilbert cheboi
- Tell No: 254743143013
- Email: icheboigilbert@gmail.com
 
 ## Lincense

MIT License
