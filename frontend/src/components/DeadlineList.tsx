import type { Deadline } from "../types/deadline";
import DeadlineItem from "./DeadlineItem";

type Props = {
  items?: Deadline[];
  onComplete: (id: string) => Promise<void> | void;
};

export default function DeadlineList({ items = [], onComplete }: Props) {
  if (items.length === 0) {
    return (
      <p className="text-sm text-gray-500">
        No deadlines yet. Add your first compliance task above.
      </p>
    );
  }

  return (
    <div className="space-y-3">
      {items.map((d) => (
        <DeadlineItem key={d.id} d={d} onComplete={onComplete} />
      ))}
    </div>
  );
}
