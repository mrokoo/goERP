type Product = {
  id: string;
  name: string;
  category_id: string;
  unit_id: string;
  openStock: {
    warehouse_id: string;
    amount: number;
  }[];
  state: string;
  note: string;
  purchase: number;
  retail: number;
  grade1: number;
  grade2: number;
  grade3: number;
  img: string;
  intro: string;
};

type Unit = {
  id: string;
  name: string;
  note?: string;
};

type Category = {
  id: string;
  name: string;
  note?: string;
};

export { Product, Unit, Category };
