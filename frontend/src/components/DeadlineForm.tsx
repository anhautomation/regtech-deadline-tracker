import { useState, type ChangeEvent, type FormEvent } from "react";
import type { DeadlineCreateInput } from "../types/deadline";

export default function DeadlineForm({ onSubmit }: { onSubmit: (v: DeadlineCreateInput) => void }) {
  const [form, setForm] = useState<DeadlineCreateInput>({
    title: "",
    category: "",
    dueDate: "",
    notes: "",
  });

  const change = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) =>
    setForm({ ...form, [e.target.name]: e.target.value });

  const submit = (e: FormEvent) => {
    e.preventDefault();
    onSubmit(form);
    setForm({ title: "", category: "", dueDate: "", notes: "" }); // reset
  };

  return (
    <form className="space-y-3 p-4 bg-white rounded shadow" onSubmit={submit}>
      <input name="title" placeholder="Title" className="input" value={form.title} onChange={change} />
      <input name="category" placeholder="Category" className="input" value={form.category} onChange={change} />
      <input name="dueDate" type="date" className="input" value={form.dueDate} onChange={change} />
      <textarea name="notes" placeholder="Notes" className="input" value={form.notes} onChange={change} />
      <button className="btn-primary w-full">Add Deadline</button>
    </form>
  );
}
