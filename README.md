# 📌 Compliance Deadline Tracker — RegTech Demo (AI-Generated)

This project is a small but meaningful RegTech feature built using **100% AI-generated code** (“vibe coding”).  
It helps early-stage Australian startups **track compliance deadlines** such as ASIC Annual Reviews, BAS, PAYG, STP, and Super.

🚀 Live demo: https://regtech-deadline-tracker.vercel.app  
📄 Backend API docs (Swagger): https://regtech-deadline-tracker.onrender.com/swagger/index.html

---

## 🎯 Feature & RegTech Pain Point

Australian founders frequently miss regulatory filings because deadlines are tracked in spreadsheets, emails, or a founder’s memory.

This demo solves one specific pain point:

> “Show me all compliance deadlines that are upcoming or overdue, allow me to tick them off, and make sure I don’t forget anything.”

Key features:
- List of regulatory deadlines
- Clear **overdue / due soon / upcoming** labels
- Colour indicators (🔴 overdue / 🟢 upcoming or completed)
- One-click **Mark done**
- Automatic popup summarising **overdue + due-soon deadlines** when the user opens the app

➡️ Small feature — **big impact** on fines, stress, and accountant time.

---

## 🧱 Architecture Overview

| Layer      | Technology        |
|-----------|-------------------|
| Frontend  | React + Vite      |
| Backend   | Go (in-memory storage) |
| Database  | None — simple and free to run |
| Deployment | Frontend → Vercel, Backend → Render |

### Backend

- REST API endpoints:
  - `GET /api/deadlines`
  - `POST /api/deadlines`
  - `POST /api/deadlines/:id/complete`
  - `GET /api/deadlines/summary`
- Stores deadlines in memory (no database)
- Seeds real AU regulatory obligations as sample data

### Frontend

- Components:
  - `DeadlinesPage`
  - `DeadlineForm`
  - `DeadlineList`
  - `DeadlineItem`
  - `OverdueSummaryModal`
- State management via custom hook `useDeadlines`
- Optimistic UI when marking as done

---

## 💻 Running the App

### Local – Backend

```bash
cd backend
go run ./cmd/server
```

### Local - Frontend

```bash
cd frontend
npm install
npm run dev
```

---

## 🌍 Production URLs

| Service  | URL                                             |
|----------|-------------------------------------------------|
| Frontend | https://regtech-deadline-tracker.vercel.app     |
| Backend  | https://regtech-deadline-tracker.onrender.com   |
| Swagger  | https://regtech-deadline-tracker.onrender.com/swagger/index.html |

---

## 🤖 Vibe-Coding Process (AI-only)

This assignment required **zero manual business-logic typing**.  
My role was **architect + reviewer**, not typist.

Development flow:
1. Designed the data model and REST API through prompting
2. Generated backend code (handlers, router, repository) using ChatGPT
3. Generated frontend components, hooks and UI behaviour using ChatGPT
4. Debugged by prompting — not manual code rewriting

Example prompts:

```
Mark-done does not update the UI. Fix the state logic so it treats items immutably.
```

```
Deadlines due today are marked overdue. Rewrite comparison to operate on date only.
```

**Takeaway:** AI is fast at generating working code, but human direction, product thinking  
and validation are still essential to ensure the final result is useful and reliable.

---

## 🧪 QA / Manual Quality Checks

Even without automated tests, the core user journeys were manually verified:

- Adding a deadline → displayed correctly
- Marking a deadline as done → status and colour update properly
- Overdue / due-today logic behaves correctly
- Summary popup appears only when relevant
- Refresh keeps correct UI state
- Network failure shows a friendly message instead of crashing the UI

---

## 🚀 Future Enhancements

- Proactive email / push notifications before deadlines
- Calendar sync (Google / Outlook)
- Xero / MYOB integration for accounting automation
- Multi-user workspaces with authentication
- Database support (Supabase / PostgreSQL)
- CI/CD pipeline and containerised deployment with Docker
- Separate deployment environments (dev / staging / production)

---

## 📎 Contact

Thank you for reviewing this project. I’m happy to discuss technical decisions or the development approach in more detail if needed.
