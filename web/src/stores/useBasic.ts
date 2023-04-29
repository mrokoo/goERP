import { defineStore } from "pinia";
import api from "@/api";
import { type Account, type Supplier, type Client } from "@/@types/basic";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const useBasic = defineStore("basic", {
  state: () => ({
    warehouse: [] as any[],
    account: [] as Account[],
    supplier: [] as Supplier[],
    client: [] as Client[],
    budget: [] as any[],
  }),
  getters: {
    warehouseOptions: (state) => {
      const option: any[] = [];
      state.warehouse.map((item) => {
        option.push({
          label: item.name,
          value: item.id,
        });
      });
      return option;
    },
  },

  actions: {
    async getWarehouse() {
      api.warehouse
        .getWarehouseList()
        .then((res: any) => {
          this.warehouse = res.data.data;
        })
        .catch((err: any) => {
          throw err;
        });
    },
    async createWarehouse(data: any) {
      try {
        const res = await api.warehouse.createWarehouse(data);
        this.warehouse.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
    async updateWarehouse(data: any) {
      try {
        const res = await api.warehouse.updateWarehouse(data);
        const index = this.warehouse.findIndex((item) => item.id === data.id);
        this.warehouse[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
    async deleteWarehouse(id: string) {
      try {
        const res = await api.warehouse.deleteWarehouse(id);
        const index = this.warehouse.findIndex((item) => item.id === id);
        this.warehouse.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async getAccount() {
      api.account
        .getAccountList()
        .then((res: any) => {
          this.account.push(...res.data.data);
        })
        .catch((err: any) => {
          throw err;
        });
    },

    async createAccount(data: any) {
      try {
        const res = await api.account.createAccount(data);
        this.account.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
    async updateAccount(data: any) {
      try {
        const res = await api.account.updateAccount(data);
        const index = this.account.findIndex((item) => item.id === data.id);
        this.account[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async deleteAccount(id: string) {
      try {
        const res = await api.account.deleteAccount(id);
        const index = this.account.findIndex((item) => item.id === id);
        this.account.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async getSupplier() {
      api.supplier
        .getSupplierList()
        .then((res: any) => {
          this.supplier = res.data.data;
        })
        .catch((err: any) => {
          throw err;
        });
    },

    async createSupplier(data: any) {
      try {
        const res = await api.supplier.createSupplier(data);
        this.supplier.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async updateSupplier(data: any) {
      try {
        const res = await api.supplier.updateSupplier(data);
        const index = this.supplier.findIndex((item) => item.id === data.id);
        this.supplier[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async deleteSupplier(id: string) {
      try {
        const res = await api.supplier.deleteSupplier(id);
        const index = this.supplier.findIndex((item) => item.id === id);
        this.supplier.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async getClient() {
      api.client
        .getClientList()
        .then((res: any) => {
          this.client = res.data.data;
        })
        .catch((err: any) => {
          throw err;
        });
    },

    async createClient(data: any) {
      try {
        const res = await api.client.createClient(data);
        this.client.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async updateClient(data: any) {
      try {
        const res = await api.client.updateClient(data);
        const index = this.client.findIndex((item) => item.id === data.id);
        this.client[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async deleteClient(id: string) {
      try {
        const res = await api.client.deleteClient(id);
        const index = this.client.findIndex((item) => item.id === id);
        this.client.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async getBudget() {
      api.budget
        .getBudgetList()
        .then((res: any) => {
          this.budget = res.data.data;
        })
        .catch((err: any) => {
          throw err;
        });
    },

    async createBudget(data: any) {
      try {
        const res = await api.budget.createBudget(data);
        this.budget.push(res.data.data);
        message.success("创建成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async updateBudget(data: any) {
      try {
        const res = await api.budget.updateBudget(data);
        const index = this.budget.findIndex((item) => item.id === data.id);
        this.budget[index] = data;
        message.success("更新成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },

    async deleteBudget(id: string) {
      try {
        const res = await api.budget.deleteBudget(id);
        const index = this.budget.findIndex((item) => item.id === id);
        this.budget.splice(index, 1);
        message.success("删除成功");
      } catch (error) {
        message.error("编号验证无效");
      }
    },
  },
});
