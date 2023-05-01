import axios from "@/utils/request";

interface BudgetAPI {
  getBudgetList: () => Promise<any>;
  createBudget: (data: any) => Promise<any>;
  updateBudget: (data: any) => Promise<any>;
  deleteBudget: (id: string) => Promise<any>;
}

const budgetAPI: BudgetAPI = {
  getBudgetList: () => axios.get("/budgets"),
  createBudget: (data: any) => {
    return axios.post("/budgets", data);
  },
  updateBudget: (data: any) => {
    return axios.put(`/budgets/${data.id}`, data);
  },

  deleteBudget: (id: string) => {
    return axios.delete(`/budgets/${id}`);
  },
};

export default budgetAPI;
