export type DeadlineStatus = "PENDING" | "COMPLETED" | "OVERDUE" | "UPCOMING" | string;

export type Deadline = {
  id: string;
  title: string;
  category: string;
  dueDate: string;
  notes: string;
  status: DeadlineStatus;
};

export type DeadlineCreateInput = {
  title: string;
  category: string;
  dueDate: string;
  notes: string;
};
