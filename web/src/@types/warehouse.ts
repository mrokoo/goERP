type Task = {
  id: string;
  warehouse_id: string;
  kind: string;
  state: string;
  items: TaskItem[];
  records: TaskRecord[];
  io: boolean;
  purchase_order_id?: string;
  purchase_return_order_id?: string;
  sale_order_id?: string;
  sale_return_order_id?: string;
  allot_id?: string;
  created_at: string;
};

type TaskItem = {
  product_id: string;
  quantity: number;
  total: number;
};

type TaskRecord = {
  id: string;
  warehouse_id: string;
  user_id: string;
  created_at: string;
  state: string;
  items: TaskRecordItem[];
};

type TaskRecordItem = {
  product_id: string;
  quantity: number;
};

type InventoryFlow = {
  id: string;
  task_id?: string;
  take_id?: string;
  product_id: string;
  warehouse_id: string;
  flow: string;
  previous: number;
  change: number;
  present: number;
  created_at: string;
};
