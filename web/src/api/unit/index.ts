import axios from "@/utils/request";
import { type Unit } from "@/@types/goods";

interface UnitAPI {
  getUnitList: () => Promise<any>;
  createUnit: (data: Unit) => Promise<any>;
  updateUnit: (data: Unit) => Promise<any>;
  deleteUnit: (id: string) => Promise<any>;
}

const unitAPI: UnitAPI = {
  getUnitList: () => axios.get("/units"),
  createUnit: (data: Unit) => {
    return axios.post("/units", data);
  },

  updateUnit: (data: Unit) => {
    return axios.put(`/units/${data.id}`, data);
  },
  deleteUnit: (id: string) => {
    return axios.delete(`/units/${id}`);
  },
};

export default unitAPI;
