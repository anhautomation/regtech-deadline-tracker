import type { Deadline, DeadlineCreateInput } from "../types/deadline";

const API_BASE = import.meta.env.VITE_API_URL || "https://regtech-deadline-tracker.onrender.com/api/v1";

export async function listDeadlines() {
  const res = await fetch(`${API_BASE}/deadlines`);
  const data = await res.json();
  return data.data;
}

export async function createDeadline(payload: DeadlineCreateInput): Promise<Deadline> {
  const res = await fetch(`${API_BASE}/deadlines`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  return res.json();
}

export async function markCompleted(id: string): Promise<Deadline> {
  const res = await fetch(`${API_BASE}/deadlines/${id}/complete`, {
    method: "PUT",
  });
  return res.json();
}
