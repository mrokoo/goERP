import axios from "@/utils/request";

interface Warehouse {
  getWarehouseList: () => Promise<any>;
  createWarehouse: (data: WarehouseType) => Promise<any>;
  updateWarehouse: (data: WarehouseType) => Promise<any>;
  deleteWarehouse: (id: string) => Promise<any>;
}

type WarehouseType = {
  name: string;
  admin?: string;
  phone?: string;
  address?: string;
  note?: string;
  state: string;
  id: string;
};

const warehouse: Warehouse = {
  getWarehouseList: () => axios.get("/warehouses"),
  createWarehouse: (warehouse: WarehouseType) => {
    return axios.post("/warehouses", warehouse);
  },
  updateWarehouse: (warehouse: WarehouseType) => {
    return axios.put(`/warehouses/${warehouse.id}`, warehouse);
  },
  deleteWarehouse: (id: string) => {
    return axios.delete(`/warehouses/${id}`);
  },
};
export default warehouse;
