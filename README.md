# 🚀 Exercise API Golang

This is a simple API built using the Go programming language with 4 main endpoints:

- `GET /health` - Check server status (always returns 200 OK)
- `POST /submit` - Accepts JSON data (name, job, citizen, hobbies)
- `GET /all` - Returns more than 10 random JSON data entries
- `GET /random` - Returns random JSON response at a fixed `/random` endpoint

---

## 📦 Tools & Requirements

Before running this project, make sure you have:

- [Go](https://go.dev/dl/) version `1.18` or later (latest recommended)
- A terminal (bash, zsh, PowerShell, etc.)
- (Optional) [Postman](https://www.postman.com/) or `curl` for API testing

---

## ⚙️ How to Set Up Go

### 🔽 1. Download and Install Go

Go to: https://go.dev/dl/  
Download and install the version suitable for your OS.

### 🛠 2. Verify Installation

After installation, check if Go is installed:
```bash
go version
