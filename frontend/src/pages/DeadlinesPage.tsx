import { useEffect, useState } from "react";
import DeadlineForm from "../components/DeadlineForm";
import DeadlineList from "../components/DeadlineList";
import { useDeadlines } from "../hooks/useDeadline";
import ReminderPopup, { type Reminder } from "../components/ReminderPopup";
import { differenceInCalendarDays, isBefore, parseISO } from "date-fns";

export default function DeadlinesPage() {
  const { items, add, complete, loading } = useDeadlines();
  const [reminders, setReminders] = useState<Reminder[]>([]);
  const [showReminders, setShowReminders] = useState(false);

  useEffect(() => {
    if (!items) return;

    const today = new Date();
    const next: Reminder[] = [];

    items.forEach((d) => {
      if (d.status === "COMPLETED") return;

      const due = parseISO(d.dueDate);
      const diff = differenceInCalendarDays(due, today);

      if (isBefore(due, today)) {
        next.push({
          id: d.id,
          title: d.title,
          message: `Overdue by ${Math.abs(diff)} day(s)`,
        });
      } else if (diff === 0) {
        next.push({
          id: d.id,
          title: d.title,
          message: "Due today",
        });
      } else if (diff <= 2) {
        next.push({
          id: d.id,
          title: d.title,
          message: `Due in ${diff} day(s)`,
        });
      }
    });

    queueMicrotask(() => {
      setReminders(next);
      setShowReminders(next.length > 0);
    });
  }, [items]);

  return (
    <div className="max-w-3xl mx-auto p-6 space-y-6">
      <h1 className="text-2xl font-bold">Compliance Deadlines</h1>

      <DeadlineForm onSubmit={add} />

      {loading ? (
        <p>Loading...</p>
      ) : (
        <DeadlineList items={items} onComplete={complete} />
      )}

      {showReminders && (
        <ReminderPopup
          reminders={reminders}
          onClose={() => setShowReminders(false)}
        />
      )}
    </div>
  );
}
