type PurchaseOrder = {
  id: string;
  warehouse_id: string;
  supplier_id: string;
  user_id: string;
  account_id: string;
  other_cost: number;
  total_cost: number;
  actal_payment: number;
  debt: number;
  created_at: string;
  is_validated: boolean;
  kind: string;
  basic?: string;
  items: Item[];
};

type Item = {
  product_id: string;
  quantity: number;
  price: number;
};
