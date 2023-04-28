import axios from "@/utils/request";
import { type Product } from "@/@types/goods";

interface ProductAPI {
  getProductList: () => Promise<any>;
  createProduct: (data: Product) => Promise<any>;
  updateProduct: (data: Product) => Promise<any>;
  deleteProduct: (id: string) => Promise<any>;
}

const productAPI: ProductAPI = {
  getProductList: () => axios.get("/products"),
  createProduct: (data: Product) => {
    return axios.post("/products", data);
  },

  updateProduct: (data: Product) => {
    return axios.put(`/products/${data.id}`, data);
  },
  deleteProduct: (id: string) => {
    return axios.delete(`/products/${id}`);
  },
};

export default productAPI;
