# ⚡ Distributed Rate Limiter (Token Bucket) – Golang

A high-performance, customizable **Token Bucket Rate Limiter** built in Go using the Gin framework. Perfect for protecting APIs from abuse and enforcing fair usage policies.

This project is built with clean code principles, structured modules, and concurrency-safe logic. Future plans include a Redis backend and cluster-aware support.

---

## 🚀 Features

* 🔄 Token Bucket algorithm: allows bursts + controlled refill
* 👤 Per-user rate limiting (via `user_id`)
* ⚙️ Configurable rate (tokens/sec) and burst size
* 🧵 Thread-safe (uses `sync.Mutex`)
* 🌐 HTTP API built with Gin
* 🧼 Clean folder structure, easy to extend

---

## 🧠 How It Works

* Every user has a **bucket**.
* Each request consumes **1 token**.
* If tokens are available → ✅ request is allowed.
* If bucket is empty → ❌ `429 Too Many Requests`.
* Tokens are **refilled automatically** based on elapsed time.

### 🔢 Example

| Config         | Value        |
| -------------- | ------------ |
| Rate           | 5 tokens/sec |
| Burst Capacity | 10 tokens    |

* User can make 10 quick requests → all allowed
* 11th request → blocked
* After 2 seconds → 10 tokens restored

---

## 🗂️ Folder Structure

```
Distributed-Rate-Limiter/
├── cmd/
│   └── server/           # App entrypoint
│       └── main.go
├── internal/
│   ├── api/              # Gin handlers
│   │   └── handler.go
│   └── limiter/          # Token bucket logic
│       └── limiter.go
├── go.mod
├── go.sum
└── README.md
```

---

## 📦 Installation

```bash
git clone https://github.com/hsraktu17/Distributed-Rate-Limiter.git
cd Distributed-Rate-Limiter
go mod tidy
go run cmd/server/main.go
```

---

## 📬 API Usage

### Endpoint:

```
GET /api/check?user_id=some-user
```

### Responses:

```json
// ✅ Request Allowed
{ "message": "Request Allowed" }

// ❌ Rate Limit Exceeded
{ "error": "Too Many Requests" }
```

---

## ⚙️ Customize Limiting Rules

In `handler.go`, change this line:

```go
l = limiter.NewLimiter(5, 10) // 5 req/sec, burst of 10
```

To whatever rate/burst you want.

---

## 🧪 Test Locally

Use `curl` or Postman:

```bash
for i in {1..15}; do
  curl "http://localhost:8080/api/check?user_id=testuser"
done
```

Try hitting it rapidly to trigger `429` responses.

---

## ⏭️ Roadmap

* [ ] Redis-based limiter (shared memory across nodes)
* [ ] Consistent hashing for sharding
* [x] Middleware support
* [ ] Docker support & deployment
* [ ] Prometheus metrics

---

## 👨‍💻 Author

**Utkarsh Raj Srivastava**
Building backend infrastructure with Go & system design.

> Drop a ⭐ if this helped!
