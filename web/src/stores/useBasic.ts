import { defineStore } from "pinia";
import api from "@/api";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const useBasic = defineStore("basic", {
  state: () => ({
    warehouse: [] as any[],
  }),
  getters: {},

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
    }
  },
});
