# ⚡ ratelimiter – Token Bucket Rate Limiter for Go

&#x20;

A production-ready, pluggable token bucket rate limiter written in Go. Built with clean abstractions for extensibility and includes drop-in support for the Gin web framework.

---

## 🚀 Features

* Token Bucket algorithm with refill logic
* Per-user rate limiting support
* In-memory backend (Redis coming soon)
* Gin middleware support (`r.Use(...)`)
* Lightweight, no external dependencies (for memory mode)

---

## 📦 Installation

```bash
go get github.com/hsraktu17/ratelimiter@v0.1.0
```

---

## 🧱 Usage

### 🧠 Basic Go

```go
import "github.com/hsraktu17/ratelimiter/limiter"

lim := limiter.NewMemoryLimiter(5, 10)
if lim.Allow("user123") {
    // process request
}
```

### 🌐 Gin Middleware

```go
import (
    "github.com/hsraktu17/ratelimiter/limiter"
    "github.com/hsraktu17/ratelimiter/ginratelimiter"
)

r := gin.Default()
r.Use(ginratelimiter.Middleware(limiter.NewMemoryLimiter(5, 10)))
```

---

## 📂 Project Structure

```
ratelimiter/
├── limiter/           # Core Limiter interface + memory implementation
│   ├── limiter.go     # Limiter interface
│   └── memory.go      # Memory-backed rate limiter
├── gin/               # Gin-specific middleware
│   └── middleware.go
├── examples/          # Example Gin server
│   └── main.go
├── go.mod
└── README.md
```

---

## 📈 Roadmap

*

---

## 📄 License

This project is licensed under the **MIT License** – see the [LICENSE](./LICENSE) file for details.

---

## 👤 Author

**Utkarsh Raj Srivastava**
Passionate about backend systems, performance engineering, and product-scale infrastructure.

> Drop a ⭐ if this helped. Contributions are welcome!
