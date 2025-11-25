import { useEffect, useState } from "react";
import type { Deadline, DeadlineCreateInput } from "../types/deadline";
import { listDeadlines, createDeadline, markCompleted } from "../api/deadline";

export function useDeadlines() {
  const [items, setItems] = useState<Deadline[]>([]);
  const [loading, setLoading] = useState(true);

  const refresh = async () => {
    try {
      const data = await listDeadlines();
      setItems(data);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    void refresh();
  }, []);

  const add = async (input: DeadlineCreateInput) => {
    await createDeadline(input);
    await refresh();
  };

  const complete = async (id: string) => {
    await markCompleted(id);
    await refresh();
  };

  return { items, add, complete, loading };
}