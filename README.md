# ⚡ ratelimiter – Token Bucket Rate Limiter for Go

A production-ready, pluggable token bucket rate limiter written in Go. Built with clean abstractions for extensibility and includes drop-in support for the Gin web framework.

---

## 🚀 Features

* Token Bucket algorithm with refill logic
* Per-user rate limiting support
* In-memory backend (Redis coming soon)
* Gin middleware support (`r.Use(...)`)
* Lightweight, no external dependencies (for memory mode)

---

## 🔧 Prerequisites

* **Go** 1.24 or newer
* **Docker** and **docker-compose** for containerized runs

---

## 📦 Installation

```bash
go get github.com/hsraktu17/Distributed-Rate-Limiter@v0.1.0
```

---

## 🧱 Usage

### 🧠 Basic Go

```go
import "github.com/hsraktu17/Distributed-Rate-Limiter/limiter"

lim := limiter.NewMemoryLimiter(5, 10)
if lim.Allow("user123") {
    // process request
}
```

### 🌐 Gin Middleware

```go
import (
    "github.com/hsraktu17/Distributed-Rate-Limiter/limiter"
    ginratelimiter "github.com/hsraktu17/Distributed-Rate-Limiter/gin"
)

r := gin.Default()
r.Use(ginratelimiter.Middleware(limiter.NewMemoryLimiter(5, 10)))
```

---

## 📂 Project Structure

```
Distributed-Rate-Limiter/
├── limiter/           # Core Limiter interface + memory implementation
│   ├── limiter.go     # Limiter interface
│   └── memory.go      # Memory-backed rate limiter
├── gin/               # Gin-specific middleware
│   └── middleware.go
├── cmd/
│   └── server/        # Example Gin server
│       └── main.go
├── examples/          # Legacy example server
│   └── main.go
├── go.mod
└── README.md
```

---

## ▶️ Running the Example Server

### Local execution

```bash
go run examples/main.go
```

Open `http://localhost:8080/ping` in your browser or via curl to test.

### Docker Compose

Build and run the containerized server:

```bash
docker-compose up --build
```

The service exposes port **8080** on the host.

---

## 📈 Roadmap

* Add Redis backend support

---

## 📄 License

This project is licensed under the **MIT License** – see the [LICENSE](./LICENSE) file for details.

---

## 👤 Author

**Utkarsh Raj Srivastava**
Passionate about backend systems, performance engineering, and product-scale infrastructure.

> Drop a ⭐ if this helped. Contributions are welcome!
