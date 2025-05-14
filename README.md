# ⚡ Go Rate Limiter (Token Bucket Algorithm)

A lightweight, production-minded **Token Bucket Rate Limiter** built in **Golang** using the Gin framework.

Use it to protect your API from abuse, apply fair-usage policies, or build a gateway-like access layer — one token at a time.

---

## 🚀 Features

- ⚙️ Token Bucket algorithm (burst + refill)
- 👤 Per-user in-memory tracking
- 💡 Customizable rate & burst limits
- 🔐 Thread-safe using `sync.Mutex`
- 🧪 Simple HTTP API to test behavior

---

## 🧠 How It Works

- Each user has a **bucket** of tokens.
- Each request **uses one token**.
- If the bucket is empty → ❌ `429 Too Many Requests`
- Tokens **refill over time**, based on a fixed rate.

📦 Example:
- Rate: `5 tokens/second`
- Burst: `10 tokens max`

→ Users can make up to **10 requests instantly**, then refill at 5/second.

---

## 🗂 Project Structure

