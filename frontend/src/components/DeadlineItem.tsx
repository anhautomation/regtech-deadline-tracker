import { isBefore, differenceInCalendarDays, parseISO } from "date-fns";
import { useMemo } from "react";
import type { Deadline } from "../types/deadline";

type Props = {
  d: Deadline;
  onComplete: (id: string) => void | Promise<void>;
};

export default function DeadlineItem({ d, onComplete }: Props) {
  const { text, color } = useMemo(() => {
    if (d.status === "COMPLETED") {
      return { text: "Completed", color: "gray" };
    }

    const today = new Date();
    const due = parseISO(d.dueDate);
    const diff = differenceInCalendarDays(due, today);

    if (isBefore(due, today)) {
      return { text: `Overdue by ${Math.abs(diff)} day(s)`, color: "red" };
    }
    if (diff <= 3) {
      return { text: `Due in ${diff} day(s)`, color: "orange" };
    }
    return { text: "Upcoming", color: "green" };
  }, [d.dueDate, d.status]);

  const statusDotClass =
    color === "green"
      ? "bg-green-500"
      : color === "orange"
      ? "bg-orange-400"
      : color === "red"
      ? "bg-red-500"
      : "bg-gray-400"; // completed

  return (
    <div className="border p-4 rounded flex items-center justify-between">
      <div>
        <h3 className="text-lg font-semibold">{d.title}</h3>
        <p className="text-sm text-gray-600">{text}</p>
        <p className="text-xs text-gray-500">Due: {d.dueDate}</p>
      </div>

      <div className="flex items-center gap-3">
        <span className={`w-3 h-3 rounded-full ${statusDotClass}`} />

        <button
          className="px-3 py-1 bg-emerald-600 text-white rounded text-sm disabled:opacity-40"
          onClick={() => onComplete(d.id)}
          disabled={d.status === "COMPLETED"}
        >
          {d.status === "COMPLETED" ? "Done" : "Mark done"}
        </button>
      </div>
    </div>
  );
}
