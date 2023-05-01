import axios from "@/utils/request";
import { type Supplier } from "@/@types/basic";

interface SupplierAPI {
  getSupplierList: () => Promise<any>;
  createSupplier: (data: Supplier) => Promise<any>;
  updateSupplier: (data: Supplier) => Promise<any>;
  deleteSupplier: (id: string) => Promise<any>;
}

const supplierAPI: SupplierAPI = {
  getSupplierList: () => axios.get("/suppliers"),
  createSupplier: (data: Supplier) => {
    return axios.post("/suppliers", data);
  },

  updateSupplier: (data: Supplier) => {
    return axios.put(`/suppliers/${data.id}`, data);
  },
  deleteSupplier: (id: string) => {
    return axios.delete(`/suppliers/${id}`);
  },
};

export default supplierAPI;
