import axios from "@/utils/request";

import { type Client } from "@/@types/basic";

interface ClientAPI {
  getClientList: () => Promise<any>;
  createClient: (data: Client) => Promise<any>;
  updateClient: (data: Client) => Promise<any>;
  deleteClient: (id: string) => Promise<any>;
}

const clientAPI: ClientAPI = {
  getClientList: () => axios.get("/customers"),
  createClient: (data: Client) => {
    return axios.post("/customers", data);
  },
  updateClient: (data: Client) => {
    return axios.put(`/customers/${data.id}`, data);
  },
  deleteClient: (id: string) => {
    return axios.delete(`/customers/${id}`);
  },
};

export default clientAPI;
