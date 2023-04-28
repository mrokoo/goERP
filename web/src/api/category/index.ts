import axios from "@/utils/request";
import { type Category } from "@/@types/goods";

interface CategoryAPI {
  getCategoryList: () => Promise<any>;
  createCategory: (data: Category) => Promise<any>;
  updateCategory: (data: Category) => Promise<any>;
  deleteCategory: (id: string) => Promise<any>;
}

const categoryAPI: CategoryAPI = {
  getCategoryList: () => axios.get("/categories"),
  createCategory: (data: Category) => {
    return axios.post("/categories", data);
  },

  updateCategory: (data: Category) => {
    return axios.put(`/categories/${data.id}`, data);
  },
  deleteCategory: (id: string) => {
    return axios.delete(`/categories/${id}`);
  },
};

export default categoryAPI;
