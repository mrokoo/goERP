type Account = {
  id: string;
  name: string;
  type: string;
  holder?: string;
  number?: string;
  note?: string;
  state: string;
  balance: number;
};

type Supplier = {
  id: string;
  name: string;
  contact?: string;
  email?: string;
  address?: string;
  account?: string;
  bank?: string;
  note?: string;
  state: string;
  debt: number;
};

type Client = {
  id: string;
  name: string;
  grade: string;
  contact?: string;
  phone?: string;
  email?: string;
  address?: string;
  note?: string;
  state: string;
  debt: number;
};

type Budget = {
  id: string;
  name: string;
  type: string;
  note: string;
  state: string;
};

export { Account, Supplier, Client, Budget };
