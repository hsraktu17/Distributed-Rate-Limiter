# 💬 Bot – Uplifting CLI Chatbot with Groq API

A simple command-line chatbot built in Go that gives you encouraging replies using the [Groq API](https://console.groq.com). It can also log your conversations to Notion (via Zapier – *Pro plan required*).

---

## ✨ Features

* Uplifting and kind replies using Groq's LLaMA or Mixtral models
* CLI-based interaction
* Logs each question + response with a timestamp
* Modular structure (Groq API logic and Webhook logic in separate files)
* `.env` powered config

---

## 📁 Project Structure

```
go-positivity-bot/
├── main.go                         # CLI logic
├── .env                            # Your secrets
├── internal/
│   ├── groq/
│   │   └── chat.go                 # Groq API call
│   └── zapier/
│       └── webhook.go             # Webhook logic (Zapier)
├── prompt.txt                      # System prompt used for the bot
```

---

## 🔧 Requirements

* Golang ≥ 1.18
* [Groq API Key](https://console.groq.com)
* [Zapier Webhooks (Pro plan only)](https://zapier.com)
* Notion account + shared integration

---

## 🚀 Getting Started

1. **Clone the repo**

```bash
git clone https://github.com/your-username/go-positivity-bot.git
cd go-positivity-bot
```

2. **Set up environment**

Create a `.env` file:

```env
GROQ_API_KEY=sk-...
ZAPIER_WEBHOOK_URL=https://hooks.zapier.com/hooks/catch/xxxx/yyyy
```

3. **Write your prompt**

Edit `prompt.txt` with the system message (optional).

4. **Run the bot**

```bash
go run main.go
```

---

## 🤖 Sample Interaction

```bash
🌟 Welcome to PositivityBot (Zapier)

You: I'm feeling stuck today
Bot: You're not alone — even small steps move you forward. Keep going 💪
```

---

## 📝 Notes

* Zapier requires a **Pro plan** to use webhooks.
* You can replace `zapier/webhook.go` with an `n8n` or `Pipedream` webhook for a free alternative.
* This project is great as a personal mood journal + daily motivation tool.

