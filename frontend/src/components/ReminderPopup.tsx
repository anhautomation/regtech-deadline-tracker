export type Reminder = {
  id: string;
  title: string;
  message: string;
};

type Props = {
  reminders: Reminder[];
  onClose: () => void;
};

export default function ReminderPopup({ reminders, onClose }: Props) {
  if (!reminders.length) return null;

  const visible = reminders.slice(0, 3);
  const extraCount = reminders.length - visible.length;

  return (
    <div className="fixed inset-0 flex items-start justify-center pt-24 bg-black/20 z-50">
      <div className="bg-white rounded-lg shadow-lg w-full max-w-md p-4">
        <h2 className="text-lg font-semibold mb-2">
          Upcoming & overdue deadlines
        </h2>

        <ul className="space-y-2 text-sm mb-3">
          {visible.map((r) => (
            <li key={r.id}>
              <p className="font-medium">{r.title}</p>
              <p className="text-gray-600">{r.message}</p>
            </li>
          ))}
          {extraCount > 0 && (
            <li className="text-gray-500 text-xs">
              …and {extraCount} more deadline(s)
            </li>
          )}
        </ul>

        <div className="flex justify-end">
          <button
            type="button"
            onClick={onClose}
            className="px-3 py-1 text-sm rounded border border-gray-300 hover:bg-gray-100"
          >
            Got it
          </button>
        </div>
      </div>
    </div>
  );
}