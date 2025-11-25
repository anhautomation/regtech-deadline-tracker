# 🏛 Regulatory Deadline Calendar API (Backend – Go)

This is the backend component for the **Regulatory Deadline Calendar** — a demo RegTech feature that helps startups track compliance deadlines (ASIC, ATO, payroll, tax, etc.).  
The service runs fully **in-memory** for free-tier deployment and does not require a database.

---

## 🚀 Features

- List compliance deadlines
- Create new deadline
- Mark deadline as completed
- Automatic seeding with realistic Australian regulatory tasks
- Fast to deploy on free tier (Render, Railway, Fly)
- REST API + Swagger documentation

---

## 📦 Tech Stack

| Component | Technology       |
|----------|------------------|
| Language | Go 1.21+         |
| Framework | Gin             |
| Docs     | Swagger (swaggo) |
| Hosting  | Render.com (recommended free tier) |

---

## 🧩 Project Structure

```text
cmd/
internal/
  adapters/http   → router + handlers
  core/domain     → entities + types
  repositories    → in-memory deadline store
main.go           → server entrypoint

---

## 🔧 Run locally
```sh
go run main.go
