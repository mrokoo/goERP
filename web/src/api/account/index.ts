import axios from "@/utils/request";
import { type Account } from "@/@types/basic";

interface AccountAPI {
  getAccountList: () => Promise<any>;
  createAccount: (data: Account) => Promise<any>;
  updateAccount: (data: Account) => Promise<any>;
  deleteAccount: (id: string) => Promise<any>;
}

const accountAPI: AccountAPI = {
  getAccountList: () => axios.get("/accounts"),
  createAccount: (data: Account) => {
    return axios.post("/accounts", data);
  },
  updateAccount: (data: Account) => {
    return axios.put(`/accounts/${data.id}`, data);
  },
  deleteAccount: (id: string) => {
    return axios.delete(`/accounts/${id}`);
  },
};

export default accountAPI;
