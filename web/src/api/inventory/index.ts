import axios from "@/utils/request";

interface InventoryAPI {
  getTaskList: () => Promise<any>;
  addRecord: (id: string, data: any) => Promise<any>;
  invalidateRecord: (id: string, rid: string) => Promise<any>;
}

const inventoryAPI: InventoryAPI = {
  getTaskList: () => axios.get("/tasks"),
  addRecord: (id: string, data: TaskRecordItem) => {
    return axios.post(`/tasks/${id}/records`, data);
  },
  invalidateRecord: (id: string, rid: string) => {
    return axios.patch(`/tasks/${id}/records/${rid}`);
  },
};

export default inventoryAPI;
